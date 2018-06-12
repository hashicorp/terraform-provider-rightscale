## 1.2.0 (Unreleased)

ENHANCEMENTS:

  - Resource: Server
    - New: [Resource Tags](http://docs.rightscale.com/cm/rs101/tagging.html)
  - Resource: Instance
    - New: Private IP Address field exposed
  - Tests:
    - Enhanced: More test coverage for rsc package
  - Datasources:
    - New: All Datasources that expose a "resource_uid" field, indicating a known value that should appear after a duration of cloud polling elapses, will enter a retry loop if defined in filter block before returning no results.
  - Resource: CWF Process
    - Modified 'array' parameter type to cast to proper formatted cwf element type vs. requiring this work to be done at the tf variable level.

BUG FIXES:

  - Resource: Server
    - Handle race condition where parent object 'server' attempts to be destroyed (and fails w/422) before child 'instance' is actually terminated.
  - Docs cleanup
  - Code tightening

## 1.1.0 (Unreleased)

NEW FEATURES:

  - RightScale [Inputs](https://docs.rightscale.com/cm/rs101/understanding_inputs.html) for Server and ServerArrays.

BUG FIXES:

  - Acceptance test race conditions
  - Acceptance test concurrency
  - Docs cleanup

## 1.0.0 (Unreleased)

NEW FEATURES:

  - Datasource: Cloud
  - Datasource: Datacenter
  - Datasource: Instance Type
  - Datasource: Multicloud ("MCI") Image
  - Datasource: Image
  - Datasource: ServerTemplate
  - Datasource: Volume
  - Datasource: Volume Type
  - Datasource: Volume Snapshot
  - Resource: CWF Process
  - Resource: Server Array
  - Resource: Security Group Rule
  - Resource: Route
  - Resource & Datasource: Instance
  - Resource & Datasource: Server
  - Resource & Datasource: SSH Key
  - Resource & Datasource: Credential
  - Resource & Datasource: Deployment
  - Resource & Datasource: Network
  - Resource & Datasource: Network Gateway
  - Resource & Datasource: Route Table
  - Resource & Datasource: Subnet
  - Resource & Datasource: Security Group Table
