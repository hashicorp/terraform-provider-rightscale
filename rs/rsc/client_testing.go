package rsc

type (
	// TestClient is a mock of the Client.
	TestClient struct {
		*TestClientExpectation
	}
)

// Make sure TestClient implements Client interface
var _ Client = (*TestClient)(nil)

// NewTestClient creates a Client which uses expectations set by the tests to
// implement the behavior.
func NewTestClient() *TestClient {
	return &TestClient{NewTestClientExpectation()}
}

// Create runs any preset expectation.
func (c *TestClient) Create(namespace, typ string, fields Fields) (*Resource, error) {
	if e := c.Expectation("Create"); e != nil {
		return e.(func(string, string, Fields) (*Resource, error))(namespace, typ, fields)
	}
	return nil, nil
}

// Get runs any preset expectation.
func (c *TestClient) Get(l *Locator) (*Resource, error) {
	if e := c.Expectation("Get"); e != nil {
		return e.(func(*Locator) (*Resource, error))(l)
	}
	return nil, nil
}

// Update runs any preset expectation.
func (c *TestClient) Update(l *Locator, fields Fields) error {
	if e := c.Expectation("Update"); e != nil {
		return e.(func(*Locator, Fields) error)(l, fields)
	}
	return nil
}

// Delete runs any preset expectation.
func (c *TestClient) Delete(l *Locator) error {
	if e := c.Expectation("Delete"); e != nil {
		return e.(func(*Locator) error)(l)
	}
	return nil
}

// Run runs any preset expectation.
func (c *TestClient) Run(l *Locator, rcl string) error {
	if e := c.Expectation("Custom"); e != nil {
		return e.(func(*Locator, string) error)(l, rcl)
	}
	return nil
}
