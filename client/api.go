package client

import (
	apiClient "fybe.com/cli/fybe/apiclient"
	cmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/config"
	"fybe.com/cli/fybe/oauth2Client"
)

func ApiClient() *apiClient.APIClient {
	configuration := apiClient.NewConfiguration()
	configuration.AddDefaultHeader("x-trace-id", "fybe_cli")
	configuration.HTTPClient = oauth2Client.BearerHttpClient(config.Conf.Oauth2TokenUrl, config.Conf.Oauth2ClientId, config.Conf.Oauth2ClientSecret, config.Conf.Oauth2User, config.Conf.Oauth2Password)
	var server apiClient.ServerConfiguration
	server.URL = config.Conf.Api
	var serverConfigurations []apiClient.ServerConfiguration
	serverConfigurations = append(serverConfigurations, server)
	configuration.Servers = serverConfigurations
	if cmd.DebugLevel == "trace" {
		configuration.Debug = true
	}

	return apiClient.NewAPIClient(configuration)
}
