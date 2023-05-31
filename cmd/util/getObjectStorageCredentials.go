package util

import (
	"context"
	_nethttp "net/http"

	apiClient "fybe.com/cli/fybe/apiclient"
	"fybe.com/cli/fybe/client"
	uuid "github.com/satori/go.uuid"
)

func GetObjectStorageCredentials(
	userId string,
	objectStorageId string,
) (apiClient.ListCredentialResponse, *_nethttp.Response, error) {
	ApiGetObjectStorageCredentialsRequest := client.ApiClient().UsersApi.ListObjectStorageCredentials(context.Background(), userId).
		XRequestId(uuid.NewV4().String())
	if objectStorageId != "" {
		ApiGetObjectStorageCredentialsRequest = ApiGetObjectStorageCredentialsRequest.ObjectStorageId(objectStorageId)
	}

	retrieveCredentialResponse, httpResp, err := ApiGetObjectStorageCredentialsRequest.Execute()
	return retrieveCredentialResponse, httpResp, err
}
