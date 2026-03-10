package cloud_authentication

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAWSCloudAuthPersonaMappingCmd = &cobra.Command{
	Use: "create-aws-cloud-auth-persona-mapping",

	Short: "Create an AWS cloud authentication persona mapping",
	Long: `Create an AWS cloud authentication persona mapping
Documentation: https://docs.datadoghq.com/api/latest/cloud-authentication/#create-aws-cloud-auth-persona-mapping`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSCloudAuthPersonaMappingResponse
		var err error

		var body datadogV2.AWSCloudAuthPersonaMappingCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudAuthenticationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateAWSCloudAuthPersonaMapping(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-aws-cloud-auth-persona-mapping")

		fmt.Println(cmdutil.FormatJSON(res, "aws_cloud_auth_config"))
	},
}

func init() {

	CreateAWSCloudAuthPersonaMappingCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAWSCloudAuthPersonaMappingCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAWSCloudAuthPersonaMappingCmd)
}
