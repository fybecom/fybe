package cmd

import (
	"fmt"
	"os"

	"fybe.com/cli/fybe/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage config files",
	Long: `View and edit config files. Configuration is merged from four different sources in this order:

	1. file /etc/.fybe.yml
	2. file $HOME/.fybe.yml
	3. environment variables starting with FYBE_
	4. arguments passed to command line

	This means that the arguments have the highest priority. You are free to edit the contents of configuration files with your favorite editor.
	For the most convenient usage specify the oauht2 credentials (oauth2-client-secret, oauth2-clientid, oauth2-password, oauth2-user).`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func ReadConfigFile() *config.Configuration {
	config := config.Configuration{}
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		return &config
	}
	// read current yaml file
	yamlFile, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not read file %v. Got error %v", cfgFile, err))
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Errors in yaml file. Could not parse due to %v", err))
	}
	return &config
}

func SaveConfigFile(config *config.Configuration) {
	newYamlFile, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could handle new values as they lead to errors %v", err))
	}
	err = os.WriteFile(cfgFile, newYamlFile, 0600)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not store to %v due to %v", cfgFile, err))
	}
	log.Info(fmt.Sprintf("Successfully saved value to %v", cfgFile))
}
