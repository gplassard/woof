package aws_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateAWSAccountCmd = &cobra.Command{
	Use: "update-aws-account [aws_account_config_id]",

	Short: "Update an AWS integration",
	Long: `Update an AWS integration
Documentation: https://docs.datadoghq.com/api/latest/a-w-s-integration/#update-aws-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSAccountResponse
		var err error

		var body datadogV2.AWSAccountUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateAWSAccount(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-aws-account")

		fmt.Println(cmdutil.FormatJSON(res, "a_w_s_account"))
	},
}

func init() {

	UpdateAWSAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateAWSAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateAWSAccountCmd)
}
