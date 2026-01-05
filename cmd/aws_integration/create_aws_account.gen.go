package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAWSAccountCmd = &cobra.Command{
	Use: "create-aws-account",

	Short: "Create an AWS integration",
	Long: `Create an AWS integration
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#create-aws-account`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSAccountResponse
		var err error

		var body datadogV2.AWSAccountCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		res, _, err = api.CreateAWSAccount(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-aws-account")

		cmd.Println(cmdutil.FormatJSON(res, "account"))
	},
}

func init() {

	CreateAWSAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAWSAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAWSAccountCmd)
}
