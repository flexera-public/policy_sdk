# fpt

`fpt` is a command line tool to aid in the development and testing of [RightScale Policies](https://docs.rightscale.com/policies/). The tool is able to syntax check, upload, and run Policies.

[![Travis CI Build Status](https://travis-ci.com/rightscale/policy_sdk.svg?branch=master)](https://travis-ci.com/rightscale/policy_sdk)

See [ChangeLog.md](https://github.com/rightscale/policy_sdk/tree/master/cmd/fpt/ChangeLog.md) for changes.

* [Installation](#installation)
  * [Configuration](#configuration)
* [Usage](#usage)
* [Change log](#changelog)
* [Contributors](#contributors)
* [License](#license)

## Installation

Since `fpt` is written in Go it is compiled to a single static binary. Extract and run the executable below:

* Linux: [v1/fpt-linux-amd64.tgz](https://binaries.rightscale.com/rsbin/fpt/v1/fpt-linux-amd64.tgz)
* Mac OS X: [v1/fpt-darwin-amd64.tgz](https://binaries.rightscale.com/rsbin/fpt/v1/fpt-darwin-amd64.tgz)
* Windows: [v1/fpt-windows-amd64.zip](https://binaries.rightscale.com/rsbin/fpt/v1/fpt-windows-amd64.zip)

### Configuration

`fpt` interfaces with the [Policy API](https://reference.rightscale.com/governance-policies/). Credentials for the API can be provided in two ways:

1. YAML-based configuration file -  Run `fpt config account <name>`, where name is a nickname for the account, to interactively write the configuration file into `$HOME/.fpt.yml` for the first time. You will be prompted for the following fields:
    * Account ID - Numeric account number, such as `60073`
    * API endpoint host - Hostname, typically `governance-3.rightscale.com`
    * Refresh Token - Your personal OAuth token available from **Settings > Account Settings > API Credentials** in the RightScale Cloud Management dashboard
2. Environment variables - These are meant to be used by build systems such as Travis CI. The following vars must be set: `FPT_LOGIN_ACCOUNT_ID`, `FPT_LOGIN_ACCOUNT_HOST`, `FPT_LOGIN_ACCOUNT_REFRESH_TOKEN`. These variables are equivalent to the ones described in the YAML section above.


### Usage

Be sure to read [Writing a Custom Policy](https://docs.rightscale.com/policies/getting_started/custom_policy.html) and check out the brief [Tutorial](https://docs.rightscale.com/policies/getting_started/policy_tutorial.html) using this tool to develop a simple policy.

The following Policy related commands are supported:

```
fpt check <file>...
  Check syntax for a Policy Template.

fpt upload <file>...
  Upload Policy Template.

fpt run [<flags>] <file> [<options>...]
  Uploads and applies the PolicyTemplate.

  Execution of the policy will then be followed. Execution log will be tailed and followed and incident printed out.

  Options are user-supplied values for "parameters" defined in the PolicyTemplate language. Options must be in the form of "<name>=<value>". For arrays,
  values must be a comma separated list.

  Example: fpt run max_snapshots.pt regions=us-east-1,us-west-2 max_snapshot_count=100

  -C, --credentials=CREDENTIALS ...
                         Credentials is the map of name and credential used to launch the policy. Credentials must be of the form "--credentials <name1>=<value1> --credentials <name2>=<value2>".
  -n, --no-log           Do not print policy execution log.
  -k, --keep             Keep applied policy running at end, for inspection in UI. Normally policy is terminated at the end.
  -r, --run-escalations  If set, escalations will be run. Normally dry_run is set to avoid running any escalations.

fpt retrieve_data [<flags>] <file> [<options>...]
  Retrieve data from a Policy Template.

  Executes a policy once and retrieve generated datasources, saving them to disk.

  Example: fpt retrieve_data my_policy.pt --names instances

  -C, --credentials=CREDENTIALS ...
                               Credentials is the map of name and credential used to launch the policy. Credentials must be of the form "--credentials <name1>=<value1> --credentials <name2>=<value2>".
  -n, --names=NAMES ...        Names of resources/datasources to retrieve. By default, all datasources will be retrieved.
  -o, --output-dir=OUTPUT-DIR  Directory to store retrieved datasources.

fpt script [<flags>] <file> [<parameters>...]
  Run the body of a script locally.

  Example: fpt script max_snapshots.pt --result snapshots volumes=@ec2_volumes.json instances=@ec2_instances.json max_count=50

  -o, --out="out.json"  Script output file. Defaults to out.json
  -r, --result=RESULT   Name of variable holding final result to extract. Required if supplying a raw JavaScript.
  -n, --name=NAME       Name of script to run, if multiple exist.
```

## Contributors

This tool is maintained by [Douglas Thrift (douglaswth)](https://github.com/douglaswth),
[Peter Schroeter (psschroeter)](https://github.com/psschroeter),
[Avinash Bhashyam (avinashbhashyam-rs)](https://github.com/avinashbhashyam-rs)


## License

The `fpt` source code is subject to the MIT license, see the
[LICENSE](https://github.com/rightscale/fpt/blob/master/LICENSE) file.
