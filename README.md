# right_pt

`right_pt` is a command line tool to aid in the development and testing of RightScale Policies. The tool is able to syntax check, upload, and run Policies. 

[![Travis CI Build Status](https://travis-ci.org/rightscale/right_pt.svg?branch=master)](https://travis-ci.org/rightscale/right_pt?branch=master)
[![AppVeyor Build Status](https://ci.appveyor.com/api/projects/status/github/rightscale/right_pt?branch=master&svg=true)](https://ci.appveyor.com/project/RightScale/right-pt?branch=master)

* [Installation](#installation)
  * [Configuration](#configuration)
* [Managing RightScripts](#managing-rightscripts)
  * [RightScript Usage](#rightscript-usage)
* [Managing ServerTemplates](#managing-servertemplates)
  * [ServerTemplate Usage](#servertemplate-usage)
* [Contributors](#contributors)
* [License](#license)

## Installation

Since `right_pt` is written in Go it is compiled to a single static binary. Extract and run the executable below:

* Linux: [v1/right_pt-linux-amd64.tgz](https://binaries.rightscale.com/rsbin/right_pt/v1/right_pt-linux-amd64.tgz)
* Mac OS X: [v1/right_pt-darwin-amd64.tgz](https://binaries.rightscale.com/rsbin/right_pt/v1/right_pt-darwin-amd64.tgz)
* Windows: [v1/right_pt-windows-amd64.zip](https://binaries.rightscale.com/rsbin/right_pt/v1/right_pt-windows-amd64.zip)

### Configuration

Right PT interfaces with the [RightScale Policy API](https://reference.rightscale.com/governance-policies/). Credentials for the API can be provided in two ways:

1. YAML-based configuration file -  Run `right_pt config account <name>`, where name is a nickname for the account, to interactively write the configuration file into `$HOME/.right_pt.yml` for the first time. You will be prompted for the following fields:
    * Account ID - Numeric account number, such as `60073`
    * API endpoint host - Hostname, typically `governance-3.rightscale.com`
    * Refresh Token - Your personal OAuth token available from **Settings > Account Settings > API Credentials** in the RightScale Cloud Management dashboard
2. Environment variables - These are meant to be used by build systems such as Travis CI. The following vars must be set: `RIGHT_PT_LOGIN_ACCOUNT_ID`, `RIGHT_PT_LOGIN_ACCOUNT_HOST`, `RIGHT_PT_LOGIN_ACCOUNT_REFRESH_TOKEN`. These variables are equivalent to the ones described in the YAML section above.


### Usage
The following Policy related commands are supported:

```
right_pt check <path> ...
  Checks syntax on passed in files. If errors are found, will print out errors and exit with failure status.

right_pt upload <path> ...
  Upload a PolicyTemplate. If the PolicyTemplate is already uploaded, it will be updated with the new content. 

right_pt run [<flags>] <path> [<options>]
  Uploads and applys the PolicyTemplate. If the the policy is already running it will be stopped and reapplied. Will print out debug logs as they come in.
  Flags:
    -n, --no-tail: Runs the policy but does not print out the logs, with the assumption this will be done in the UI.
  Options: Options are user-supplied values for "parameters" defined in the PolicyTemplate language. Options must be in the form of "<name>=<value>". For arrays, values must be a comma separated list
  Examples:
    right_pt run mypolicy.pt "regions=us-east-1,us-east-2,us-west-2" "max_count=2" "param with space=value with space"

```

## Contributors

This tool is maintained by [Douglas Thrift (douglaswth)](https://github.com/douglaswth),
[Peter Schroeter (psschroeter)](https://github.com/psschroeter)

## License

The `right_pt` source code is subject to the MIT license, see the
[LICENSE](https://github.com/douglaswth/right_pt/blob/master/LICENSE) file.
