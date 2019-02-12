# right_pt

`right_pt` is a command line tool to aid in the development and testing of RightScale Policies. The tool is able to syntax check, upload, and run Policies.

[![Travis CI Build Status](https://travis-ci.com/rightscale/right_pt.svg?token=6Udhsz2ZbD68aBb7ApTx&branch=master)](https://travis-ci.com/rightscale/right_pt)

* [Installation](#installation)
  * [Configuration](#configuration)
* [Usage](#usage)
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
right_pt check <file>...
  Check syntax for a Policy Template.

right_pt upload <file>...
  Upload Policy Template.

right_pt run [<flags>] <file> [<options>...]
  Uploads and applies the PolicyTemplate.

  Execution of the policy will then be followed. Execution log will be tailed and followed and incident printed out.

  Options are user-supplied values for "parameters" defined in the PolicyTemplate language. Options must be in the form of "<name>=<value>". For arrays,
  values must be a comma separated list.

  Example: right_pt run max_snapshots.pt regions=us-east-1,us-west-2 max_snapshot_count=100

  -n, --no-log           Do not print policy execution log.
  -k, --keep             Keep applied policy running at end, for inspection in UI. Normally policy is terminated at the end.
  -r, --run-escalations  If set, escalations will be run. Normally dry_run is set to avoid running any escalations.

right_pt retrieve_data [<flags>] <file> [<options>...]
  Retrieve data from a Policy Template.

  Executes a policy once and retrieve generated datasources, saving them to disk.

  Example: right_pt retrieve_data my_policy.pt --names instances

  -n, --names=NAMES ...      Names of resources/datasources to retrieve. By default, all datasources will be retrieved.
  -o, --output-dir=OUTPUTDIR  Directory to store retrieved datasources.

right_pt script [<flags>] <file> [<parameters>...]
  Run the body of a script locally.

  Example: right_pt script max_snapshots.pt --result snapshots volumes=@ec2_volumes.json instances=@ec2_instances.json max_count=50

  -o, --out="out.json"  Script output file. Defaults to out.json
  -r, --result=RESULT   Name of variable holding final result to extract. Required if supplying a raw JavaScript.
  -n, --name=NAME       Name of script to run, if multiple exist.
```

## Contributors

This tool is maintained by [Douglas Thrift (douglaswth)](https://github.com/douglaswth),
[Peter Schroeter (psschroeter)](https://github.com/psschroeter),
[Avinash Bhashyam (avinashbhashyam-rs)](https://github.com/avinashbhashyam-rs)


## License

The `right_pt` source code is subject to the MIT license, see the
[LICENSE](https://github.com/rightscale/right_pt/blob/master/LICENSE) file.
