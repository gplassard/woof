package aws_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateAWSAccountCCMConfigCmd = &cobra.Command{
	Use: "update-aws-account-ccm-config [aws_account_config_id]",

	Short: "Update AWS CCM config",
	Long: `Update AWS CCM config
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#update-aws-account-ccm-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSCcmConfigResponse
		var err error

		var body datadogV2.AWSCcmConfigRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateAWSAccountCCMConfig(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-aws-account-ccm-config")

		fmt.Println(cmdutil.FormatJSON(res, "ccm_config"))
	},
}

func init() {

	UpdateAWSAccountCCMConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateAWSAccountCCMConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateAWSAccountCCMConfigCmd)
}
