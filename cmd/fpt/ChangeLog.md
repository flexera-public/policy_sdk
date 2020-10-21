v1.2.0 / 2020-10-21
-------------------
* Add support for Flexera One refresh tokens using an optional `flexera` boolean parameter in account entries in
  the configuration file

v1.1.4 / 2020-06-30
-------------------
* Remove double read of access token response for non-200 codes so the actual error message comes through

v1.1.3 / 2020-06-26
-------------------
* Fix an incorrect use of `fmt.Printf` without a format to use `fmt.Print` instead
* Handle the `'EOS'` and `"EOS"` forms of [HEREDOC](https://ruby-doc.org/core-2.2.7/doc/syntax/literals_rdoc.html#label-Here+Documents)s when detecting JavaScript `code` blocks in the `fpt script` subcommand
* Actually check if `--result`/`-r` is being passed when executing raw JavaScript with the `fpt script` subcommand
* Handle Policy Template (Ruby) comments correctly when parsing for the `fpt script` subcommand

v1.1.2 / 2020-05-21
-------------------
* Actually check for any updates if `update.check` is set to `true`

v1.1.1 / 2020-05-20
-------------------
* Clean up JSON output and output files so they are more readable without escaped `<`, `>`, and `&` characters
* Clean up `console.log` and `console.dir` output so brackets and braces line up correctly
* Fix a bug where the format for `fmt.Printf` included user input unnecessarily in JavaScript `console.log`/`console.dir` functions
* Only write `fpt script` JSON output if there is no error

v1.1.0 / 2020-05-05
-------------------
* Print all JavaScript compilation errors in the `fpt script` command (#21)

v1.0.9 / 2020-03-12
-------------------
* Bump timeout for http client to 5 minutes
* Fix panic while passing non-existent options (#17)

v1.0.8 / 2020-02-28
-------------------
* Add short flag (`-C`) for credentials to go along with `--credentials`.

v1.0.7 / 2020-02-28
-------------------
* Fix errors with incident index.

v1.0.6 / 2020-02-27
-------------------
* Add credential support on `fpt run` or `fpt retrieve_data`.

v1.0.5 / 2020-02-10
-------------------
* Do not require parameters with defaults when running `fpt run` or `fpt retrieve_data`.

v1.0.4 / 2019-08-07
-------------------
* Fix script command with CRLF line endings.

v1.0.3 / 2019-03-14
-------------------
* Fix errors for bad hostname in config.
* Fix bad_request error for template in subdirectories.

v1.0.2 / 2019-02-15
-------------------
* Add way to past list param to script command
* Renamed right_pt to fpt (flexera policy tool)

v1.0.1 / 2019-02-12
-------------------
* Add in update code

v1.0.0 / 2019-02-12
-------------------
* Initial release
