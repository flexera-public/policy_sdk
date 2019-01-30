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
upload <file>...
  Upload Policy Template.

check <file>...
  Check syntax for a Policy Template.

run [<flags>] <file> [<options>...]
  Uploads and applys the PolicyTemplate.

  Execution of the policy will then be followed. Execution log will be tailed and followed and incident printed out.

  -n, --no-log           Do not print policy execution log.
  -k, --keep             Keep applied policy running at end, for inspection in UI. Normally policy is terminated at the end.
  -r, --run-escalations  If set, escalations will be run. Normally dry_run is set to avoid running any escalations.
```

## Contributors

This tool is maintained by [Douglas Thrift (douglaswth)](https://github.com/douglaswth),
[Peter Schroeter (psschroeter)](https://github.com/psschroeter)

## License

The `right_pt` source code is subject to the MIT license, see the
[LICENSE](https://github.com/douglaswth/right_pt/blob/master/LICENSE) file.
