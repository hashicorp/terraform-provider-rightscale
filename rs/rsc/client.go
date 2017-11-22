package rsc

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/rightscale/rsc.v6/rsapi"
)

var (
	// ErrNotFound is returned by Get and Update if no resource with the
	// given locator could be found.
	ErrNotFound = fmt.Errorf("resource not found")

	// rshosts list the RightScale API hostnames.
	rshosts = []string{"us-3.rightscale.com", "us-4.rightscale.com"}
)

type (
	// Client is the RightScale API client.
	Client interface {
		// Create creates a new resource.
		Create(namespace, typ string, fields Fields) (*Resource, error)
		// Get retrieves a resource given its locator. Get returns
		// ErrNotFound if no resource could be found for the given
		// locator.
		Get(*Locator) (*Resource, error)
		// Update updates the fields of a resource. Update returns
		// ErrNotFound if no resource could be found for the given
		// locator.
		Update(*Locator, Fields) error
		// Delete deletes an existing resource. Delete does nothing if
		// the resource cannot be found.
		Delete(*Locator) error
		// Run runs custom RCL code. The code may make use of the
		// @res reference to run actions on the resource.
		Run(*Locator, string) error
	}

	// client is the Client interface implementation.
	client struct {
		// APIToken is the token used to authenticate RightScale API
		// requests.
		APIToken string
		// ProjectID is the id of the RightScale project (a.k.a.
		// account)
		ProjectID int

		rs *rsapi.API
	}

	// Resource represents a resource managed by the RightScale platform.
	Resource struct {
		// Locator is the resource locator.
		Locator *Locator
		// Fields lists the resource fields.
		Fields Fields
	}

	// Locator contains the information needed to manage a cloud resource
	// via the RightScale APIs.
	Locator struct {
		// Href is the resource path as defined by the underlying
		// service provider.
		Href string
		// Namespace identifies the service that exposes the resource.
		// The value can be one of the RightScale built-in namespaces:
		// "rs_cm", "rs_ss" or "rs_ca" or the name of a RightScale
		// plugin.
		Namespace string
		// Type is the name of the resource type scoped by the
		// namespace, e.g. "servers".
		Type string
	}

	// Fields represent arbitrary resource fields as consumed by the
	// underlying API.
	Fields map[string]interface{}
)

// New attempts to auth against all the RightScale hosts and initializes the
//  RightScale client on success.
func New(token string, projectID int) (Client, error) {
	auth := rsapi.NewOAuthAuthenticator(token, projectID)
	for _, host := range rshosts {
		rs := rsapi.New(host, auth)
		if err := rs.CanAuthenticate(); err == nil {
			req, err := rs.BuildHTTPRequest("GET", "/api/sessions/accounts", "1.5", nil, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to make session accounts request: %s", err)
			}
			resp, err := rs.PerformRequest(req)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve accounts: %s", err)
			}
			as, err := rs.LoadResponse(resp)
			if err != nil {
				return nil, fmt.Errorf("failed to load session accounts: %s", err)
			}
			if err := checkProject(as.([]interface{}), projectID); err != nil {
				return nil, err
			}
			return &client{
				APIToken:  token,
				ProjectID: projectID,
				rs:        rs,
			}, nil
		}
	}
	return nil, fmt.Errorf("failed to authenticate")
}

// Get retrieves the details of the resource pointed to by l.
// The field 'Type' of the given Locator may be ommitted.
//
// Get returns nil if there is no resource for the given locator.
func (rsc *client) Get(l *Locator) (*Resource, error) {
	rcl := fmt.Sprintf(`@resource = %s.get(href: "%s")
		$resource = to_object(@resource)
		$fields = to_json($resource["details"][0])
		$type = $resource["type"]`, l.Namespace, l.Href)

	outputs, err := rsc.runRCL(rcl, "$fields", "$type")
	if err != nil {
		return nil, err
	}

	var fields Fields
	fs := outputs[0].(string)
	err = json.Unmarshal([]byte(fs), &fields)
	if err != nil {
		return nil, err
	}
	typ := outputs[1].(string)
	loc := Locator{Namespace: l.Namespace, Type: typ, Href: l.Href}

	return &Resource{Locator: &loc, Fields: fields}, nil
}

// Update overwrite the fields of the resource.
// The field 'Type' of the resource Locator may be ommitted.
func (rsc *client) Update(l *Locator, fields Fields) error {
	// Make it more convenient to update CM resources
	if l.Namespace == "rs_cm" {
		scoped := len(fields) == 1
		if scoped {
			for k := range fields {
				scoped = k == l.Type
			}
		}
		if !scoped {
			fields = Fields{l.Type: fields}
		}
	}

	js, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	rcl := fmt.Sprintf(`@resource = %s.get(href: "%s")
		@resource.update(%s)`, l.Namespace, l.Href, js)

	_, err = rsc.runRCL(rcl)
	return err
}

// Create creates the given resource. The "Href" field of the resource locator
// should not be set on input, it is set in the result.
func (rsc *client) Create(namespace, typ string, fields Fields) (*Resource, error) {
	m := map[string]interface{}{
		"namespace": namespace,
		"type":      typ,
		"fields":    fields,
	}
	js, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	rcl := fmt.Sprintf(`@res = %s
		provision(@res)
		$href = @res.href
		$res = to_object(@res)
		$fields = to_json($res["details"][0])`, js)

	outputs, err := rsc.runRCL(rcl, "$href", "$fields")
	if err != nil {
		return nil, err
	}

	var ofields Fields
	err = json.Unmarshal([]byte(outputs[1].(string)), &ofields)
	if err != nil {
		return nil, err
	}

	loc := Locator{
		Namespace: namespace,
		Type:      typ,
		Href:      outputs[0].(string),
	}
	return &Resource{Locator: &loc, Fields: ofields}, nil
}

// Delete deletes the given resource.
// Only the Href field or res needs to be initialized.
func (rsc *client) Delete(l *Locator) error {
	rcl := fmt.Sprintf("@res = %s.get(href: %q)\ndelete(@res)",
		l.Namespace, l.Href)
	_, err := rsc.runRCL(rcl)
	return err
}

// Run runs custom RCL code that may make use of @res to run actions on the resource.
func (rsc *client) Run(l *Locator, rcl string) error {
	rcl = fmt.Sprintf("@res = %s.get(href: %q)\n%s",
		l.Namespace, l.Href, rcl)
	_, err := rsc.runRCL(rcl)
	return err
}

// runRCL runs the given RCL code synchronously and returns the process outputs.
func (rsc *client) runRCL(rcl string, outputs ...string) ([]interface{}, error) {
	var (
		projectID = strconv.Itoa(rsc.ProjectID)
		processID string
	)
	{
		source := "define main() "
		if len(outputs) > 0 {
			source += "return " + strings.Join(outputs, ", ") + " "
		}
		source += "do\nsub timeout: 1h do\n" + rcl + "\nend\nend"
		payload := rsapi.APIParams{
			"source":      source,
			"main":        "main",
			"rcl_version": "2",
			"parameters":  nil,
			"application": "cwfconsole",
			"created_by": map[string]interface{}{
				"id":    0,
				"name":  "Terraform",
				"email": "support@rightscale.com",
			},
			"refresh_token": rsc.APIToken,
		}
		res, err := rsc.requestCWF("post", "/cwf/v1/accounts/"+projectID+"/processes", nil, payload)
		if err != nil {
			return nil, err
		}
		pref := res.(map[string]interface{})["Location"]
		parts := strings.Split(pref.(string), "/")
		processID = parts[len(parts)-1]

		// print link to CWF console if DEBUG is set, mainly useful for tests
		if os.Getenv("DEBUG") != "" {
			host := strings.Replace(rsc.rs.Host, "us-", "selfservice-", 1)
			fmt.Printf("CWF process created: https://%s/designer/processes/%s\n", host, processID)
		}
	}

	timeout := time.After(1 * time.Hour)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			var (
				res map[string]interface{}
				err error
			)
			{
				path := "/cwf/v1/accounts/" + projectID + "/processes/" + processID
				r, err := rsc.requestCWF("get", path, rsapi.APIParams{"view": "expanded"}, nil)
				if err == nil {
					res = r.(map[string]interface{})
				}
			}
			if err != nil {
				return nil, err
			}

			switch res["status"].(string) {
			case "not_started", "running":
				continue

			case "completed":
				outs := res["outputs"].([]interface{})
				outputs := make([]interface{}, len(outs))
				for i, out := range outs {
					v := out.(map[string]interface{})["value"]
					outputs[i] = v.(map[string]interface{})["value"]
				}
				return outputs, nil

			default:
				var msg string
				{
					task := res["tasks"].([]interface{})[0]
					err := task.(map[string]interface{})["error"]
					if err == nil {
						msg = "[no error]"
					} else {
						msg = err.(map[string]interface{})["message"].(string)
					}
				}
				return nil, rclError(msg)
			}
		case <-timeout:
			return nil, fmt.Errorf("timed out after one hour")
		}
	}
}

// requestCWF makes a request to the RightScale Cloud Workflow API.
func (rsc *client) requestCWF(method, url string, params, payload rsapi.APIParams) (interface{}, error) {
	req, err := rsc.rs.BuildHTTPRequest(strings.ToUpper(method), url, "1.0", params, payload)
	if err != nil {
		return nil, err
	}
	req.Host = strings.Replace(rsc.rs.Host, "us-", "cloud-workflow", 1)

	res, err := rsc.rs.PerformRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := rsc.rs.LoadResponse(res)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// checkProject verifies that the giveh project ID is one of the projects listed in the
// session/accounts response 'as'. It returns nil if so, an error otherwise.
func checkProject(as []interface{}, projectID int) error {
	for _, a := range as {
		links := a.(map[string]interface{})["links"].([]interface{})
		for _, l := range links {
			href := l.(map[string]interface{})["href"].(string)
			idx := strings.LastIndex(href, "/")
			id, err := strconv.Atoi(href[idx+1:])
			if err != nil {
				return fmt.Errorf("invalid project ID %q", href[idx+1:])
			}
			if id == projectID {
				return nil
			}
		}
	}
	return fmt.Errorf("session does not give access to project %d", projectID)
}

// rclError analyzes the error message returned by runCWF and maps it to one of
// the error variables defined in this package. Right now this only looks for
// not found errors. It uses a heuristic that looks for the text
// "ResourceNotFound" as returned by the RightScale API 1.5 or the status code
// 404.
func rclError(err string) error {
	if err == "" {
		return fmt.Errorf("[unknown error]")
	}
	if strings.Contains(err, "ResourceNotFound") || strings.Contains(err, "status code '404'") {
		return ErrNotFound
	}
	return errors.New(err)
}
