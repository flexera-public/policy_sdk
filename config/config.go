package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/spf13/viper"
)

type (
	ConfigViper struct {
		*viper.Viper
		Account  *Account
		Accounts map[string]*Account
	}

	Account struct {
		Flexera      *bool
		Host         string
		ID           int
		RefreshToken string `mapstructure:"refresh_token" yaml:"refresh_token"`
	}
)

var (
	Config        ConfigViper
	boolRegexp    = regexp.MustCompile(`^(?i:true)$`)
	hostRegexp    = regexp.MustCompile(`^(?:governance-(\d+)\.(test\.)?rightscale\.com|eu-central-1\.policy-eu\.flexeraeng\.com|api\.eu-central-1\.policy-eu\.flexeraeng\.com)$`)
	hostRegexpDep = regexp.MustCompile(`^(?:eu-central-1\.policy-eu\.flexeraeng\.com)$`)
)

func init() {
	Config.Viper = viper.New()
	Config.SetDefault("login", map[string]interface{}{"accounts": make(map[string]interface{})})
	Config.SetDefault("update", map[string]interface{}{"check": true})
	Config.SetEnvPrefix("fpt")
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	Config.AutomaticEnv()
}

func ReadConfig(configFile, account string) error {
	Config.SetConfigFile(configFile)
	err := Config.ReadInConfig()
	if err != nil {
		if _, ok := err.(*os.PathError); !(ok &&
			Config.IsSet("login.account.id") &&
			Config.IsSet("login.account.host") &&
			Config.IsSet("login.account.refresh_token")) {
			return err
		}
	}

	err = Config.UnmarshalKey("login.accounts", &Config.Accounts)
	if err != nil {
		return fmt.Errorf("%s: %s", configFile, err)
	}

	if Config.IsSet("login.account.id") &&
		Config.IsSet("login.account.host") &&
		Config.IsSet("login.account.refresh_token") {
		Config.Account = &Account{
			ID:           Config.GetInt("login.account.id"),
			Host:         Config.GetString("login.account.host"),
			RefreshToken: Config.GetString("login.account.refresh_token"),
		}
		if Config.IsSet("login.account.flexera") {
			flexera := Config.GetBool("login.account.flexera")
			Config.Account.Flexera = &flexera
		}
	} else {
		var ok bool
		if account == "" {
			defaultAccount := Config.GetString("login.default_account")
			Config.Account, ok = Config.Accounts[strings.ToLower(defaultAccount)]
			if !ok {
				return fmt.Errorf("%s: could not find default account: %s", configFile, defaultAccount)
			}
		} else {
			Config.Account, ok = Config.Accounts[strings.ToLower(account)]
			if !ok {
				return fmt.Errorf("%s: could not find account: %s", configFile, account)
			}
		}
	}

	return nil
}

func (config *ConfigViper) GetAccount(id int, host string) (*Account, error) {
	for _, account := range config.Accounts {
		if account.ID == id && account.Host == host {
			return account, nil
		}
	}

	return nil, fmt.Errorf("Could not find account for account/host: %d %s", id, host)
}

// Obtain input via STDIN then print out to config file
// Example of config file
// login:
//   default_account: acct1
//   accounts:
//     acct1:
//       id: 67972
//       host: us-3.rightscale.com
//       refresh_token: abc123abc123abc123abc123abc123abc123abc1
//     acct2:
//       id: 60073
//       host: us-4.rightscale.com
//       refresh_token: zxy987zxy987zxy987zxy987xzy987zxy987xzy9
func (config *ConfigViper) SetAccount(name string, setDefault bool, input io.Reader, output io.Writer) error {
	// if the default account isn't set we should set it to the account we are setting
	if !config.IsSet("login.default_account") {
		setDefault = true
	}

	// get the settings and specifically the login settings into a map we can manipulate and marshal to YAML unhindered
	// by the meddling of the Viper
	settings := config.AllSettings()
	if _, ok := settings["login"]; !ok {
		settings["login"] = map[string]interface{}{"accounts": make(map[string]interface{})}
	}
	loginSettings := settings["login"].(map[string]interface{})

	// set the default account if we want or need to
	if setDefault {
		loginSettings["default_account"] = name
	}

	// get the previous value for the named account if it exists and construct a new account to populate
	oldAccount, ok := config.Accounts[name]
	newAccount := &Account{}

	// prompt for the account ID and use the old value if nothing is entered
	fmt.Fprint(output, "Account ID")
	if ok {
		fmt.Fprintf(output, " (%d)", oldAccount.ID)
	}
	fmt.Fprint(output, ": ")
	fmt.Fscanln(input, &newAccount.ID)
	if ok && newAccount.ID == 0 {
		newAccount.ID = oldAccount.ID
	}

	// prompt for the API endpoint host and use the old value if nothing is entered
	fmt.Fprint(output, "API endpoint host")
	if ok {
		fmt.Fprintf(output, " (%s)", oldAccount.Host)
	}
	fmt.Fprint(output, ": ")
	fmt.Fscanln(input, &newAccount.Host)
	if ok && newAccount.Host == "" {
		newAccount.Host = oldAccount.Host
	}

	// prompt for the refresh token and use the old value if nothing is entered
	fmt.Fprint(output, "Refresh token")
	if ok {
		fmt.Fprintf(output, " (%s)", oldAccount.RefreshToken)
	}
	fmt.Fprint(output, ": ")
	fmt.Fscanln(input, &newAccount.RefreshToken)
	if ok && newAccount.RefreshToken == "" {
		newAccount.RefreshToken = oldAccount.RefreshToken
	}

	// prompt for whether the refresh token is from Flexera One
	fmt.Fprint(output, "Flexera One (true if the refresh token is from the Flexera One platform, false if it is from the RightScale dashboard)")
	if ok {
		flexera := false
		if oldAccount.Flexera != nil {
			flexera = *oldAccount.Flexera
		}
		fmt.Fprintf(output, " (%t)", flexera)
	}
	fmt.Fprint(output, ": ")
	var flexera string
	fmt.Fscanln(input, &flexera)
	if ok && flexera == "" {
		newAccount.Flexera = oldAccount.Flexera
	} else {
		newAccount.Flexera = new(bool)
		*newAccount.Flexera = boolRegexp.MatchString(flexera)
	}

	// add the new account to the map of accounts overwriting any old value
	accounts := loginSettings["accounts"].(map[string]interface{})
	accounts[name] = newAccount

	// make sure all accounts have flexera set
	for _, v := range accounts {
		switch account := v.(type) {
		case *Account:
			if account.Flexera == nil {
				account.Flexera = new(bool) // the zero value of a bool is false
			}
		case map[string]interface{}:
			if _, ok := account["flexera"]; !ok {
				account["flexera"] = false
			}
		}
	}

	// render the settings map as YAML
	yml, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}

	configPath := config.ConfigFileUsed()
	// back up the current config file before writing a new one or if one does not exist, make sure the directory exists
	if err := os.Rename(configPath, configPath+".bak"); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(filepath.Dir(configPath), 0700); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// create a new config file which only the current user can read or write
	configFile, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer configFile.Close()

	// write the YAML into the config file
	if _, err := configFile.Write(yml); err != nil {
		return err
	}

	return nil
}

func (config *ConfigViper) ShowConfiguration(output io.Writer) error {
	// Check if config file exists
	if _, err := os.Stat(config.ConfigFileUsed()); err != nil {
		return err
	}

	yml, err := yaml.Marshal(config.AllSettings())
	if err != nil {
		return err
	}
	output.Write(yml)

	return nil
}

func (a *Account) Validate() error {
	if !hostRegexp.MatchString(a.Host) {
		return fmt.Errorf("invalid host: must be of form governance-<shard number>.rightscale.com or api.eu-central-1.policy-eu.flexeraeng.com")
	}
	if hostRegexpDep.MatchString(a.Host) {
		fmt.Fprintln(os.Stderr, "Warning: API endpoint host eu-central-1.policy-eu.flexeraeng.com is deprecated and will be removed soon. Change profile API endpoint host to api.eu-central-1.policy-eu.flexeraeng.com before May 23rd")
	}
	return nil
}

func (a *Account) AuthHost() string {
	matches := hostRegexp.FindStringSubmatch(a.Host)
	if len(matches) == 0 {
		return ""
	}
	shardNum := matches[1]
	testHost := matches[2]
	if shardNum == "" {
		return "login.flexera.eu"
	} else if a.Flexera != nil && *a.Flexera {
		if testHost != "" {
			return "login.flexeratest.com"
		}
		return "login.flexera.com"
	} else {
		prefix := "us"
		if shardNum == "10" {
			prefix = "telstra"
		}
		if testHost != "" {
			prefix = "moo"
		}
		return fmt.Sprintf("%s-%s.%srightscale.com", prefix, shardNum, testHost)
	}
}

func (a *Account) AppHostAndIsFlexera() (string, bool) {
	matches := hostRegexp.FindStringSubmatch(a.Host)
	if len(matches) == 0 {
		return "", false
	}
	shardNum := matches[1]
	testHost := matches[2]
	if shardNum == "" {
		return "app.flexera.eu", true
	} else if a.Flexera != nil && *a.Flexera {
		if testHost != "" {
			return "app.flexeratest.com", true
		}
		return "app.flexera.com", true
	} else {
		if testHost != "" {
			return "governance.test.rightscale.com", false
		}
		return "governance.rightscale.com", false
	}
}
