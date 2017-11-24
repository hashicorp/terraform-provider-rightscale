# Notes On Adding New RightScale Provider Resources

## Use Helper Functions

The file `resource.go` contains default implementations for the resource
callbacks. Take advantage when applicable, see the note in the section below on
fields that cannot be provided as part of creation for an exception.

## Required, Optional and Computed Fields

The resource fields are listed under "Attributes" in the online docs, for
example:
[http://reference.rightscale.com/api1.5/media_types/MediaTypeSubnet.html#attributes].
That list must be extended with the list of parameters that both the "create"
and "update" actions if any accept.

Fields that are required by the RightScale API to create a resource must have
`Required(true)` (it's OK for them to be required even for udpates because they
are always present in the Terraform resource data).

Fields that are optional to create a resource or fields that are only available
to update must have `Optional(true)`. If a resource has fields that can only be
updated and cannot be provided upon creation then create a custom `Create`
function that creates the resource then updates it. See the `Network` resource
for an example.

Fields that are read-only (e.g. `resource_uid`, `links`) must have Computed(true).

## Field Mappings

RightScale resource fields that are strings that can only have the values
`"true"` or `"false"` must be mapped to boolean fields in the corresponding
Terraform resource.

RightScale fields that are objects (e.g. `cloud_specific_attributes`) are
described using a single element list that contains a resource schema listing
the object fields.

Fields that are nested in a top level field with the name of the resource in
RightScale create and update APIs are mapped to top-level fields in Terraform.

## Data Sources

Data sources use a `filter` field in terraform that gets mapped to the
`filter[]` query string paramater when making API requests to RightScale.

## Descriptions

Always provide a description for create and update fields (OK not to provide
description for Compute(true) fields).

## Partial Updates

If a resource state can be modified via actions that do not change fields
directly then see if the action can be represented as a field and if so use
partial state updates to run the action on the resource when the Terraform field
changes. See the  deployment "locked" field as an example.

## TBD

Add support for views: make "view" a optional field of data sources and add
fields from all views as computed fields.