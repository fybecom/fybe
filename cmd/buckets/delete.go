package cmd

import (
	"context"
	"fmt"
	"net/url"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	"fybe.com/cli/fybe/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteBucketsCmd = &cobra.Command{
	Use:     "bucket --storageId [objectStorageId] --name [bucketName]",
	Short:   "Delete a specific bucket.",
	Long:    `Delete a bucket you created by its name.`,
	Example: `fybe delete bucket --storageId 1de451d4-9737-487d-9e5d-8f50a6c8d558 --name m123`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveObjectStorageRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveObjectStorage(context.Background(), deleteBucketObjectStorageId).
			XRequestId(uuid.NewV4().String())

		objStorageRetrieveResponse, httpResp, err := ApiRetrieveObjectStorageRequest.Execute()
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while retrieving object storage with id : %v", deleteBucketObjectStorageId))

		if len(objStorageRetrieveResponse.Data) == 0 {
			log.Fatal(fmt.Sprintf("No Object Storage could be found with id : %v", deleteBucketObjectStorageId))
		}

		objStorage := objStorageRetrieveResponse.Data[0]

		keycloakId := util.GetKeycloakId(config.Conf.Oauth2User)

		// get user credentials
		retrieveCredentialResponse, httpResp, err := util.GetObjectStorageCredentials(keycloakId, deleteBucketObjectStorageId)
		util.HandleErrors(err, httpResp, fmt.Sprintf("Error while getting credentials for object storage with id : %v", deleteBucketObjectStorageId))
		awsAccessKeyCred := retrieveCredentialResponse.Data[0].AccessKey
		awsSecretKeyCred := retrieveCredentialResponse.Data[0].SecretKey

		s3Url, err := url.Parse(objStorage.S3Url)
		if err != nil {
			log.Fatal(err)
		}

		client, err := minio.New(s3Url.Host, &minio.Options{
			Creds: credentials.NewStaticV4(
				awsAccessKeyCred,
				awsSecretKeyCred,
				"",
			),
			Secure: true,
		})
		if err != nil {
			log.Fatal(err)
		}

		err = client.RemoveBucket(context.Background(), deleteBucketName)

		if err != nil {
			fmt.Println(err)
			log.Fatal("Could not delete bucket " + deleteBucketName)
		}

		fmt.Printf("Bucket " + deleteBucketName + " deleted.\n")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("storageId", cmd.Flags().Lookup("storageId"))
		deleteBucketObjectStorageId = viper.GetString("storageId")

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		deleteBucketName = viper.GetString("name")

		if deleteBucketObjectStorageId == "" {
			cmd.Help()
			log.Fatal("storageId missing.")
		}
		if deleteBucketName == "" {
			cmd.Help()
			log.Fatal("bucketName missing.")
		}

		return nil
	},
}

func init() {
	cliCmd.DeleteCmd.AddCommand(deleteBucketsCmd)
	deleteBucketsCmd.Flags().StringVarP(&deleteBucketObjectStorageId, "storageId", "s", "", `Id of the objectStorage from which the bucket will be deleted`)
	deleteBucketsCmd.Flags().StringVarP(&deleteBucketName, "name", "n", "", `Name of the bucket that will be deleted`)
}
