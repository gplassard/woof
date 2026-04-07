package aws_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAWSAccountCCMConfigCmd = &cobra.Command{
	Use: "create-aws-account-ccm-config [aws_account_config_id]",

	Short: "Create AWS CCM config",
	Long: `Create AWS CCM config
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#create-aws-account-ccm-config`,
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
		res, _, err = api.CreateAWSAccountCCMConfig(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-aws-account-ccm-config")

		fmt.Println(cmdutil.FormatJSON(res, "ccm_config"))
	},
}

func init() {

	CreateAWSAccountCCMConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAWSAccountCCMConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAWSAccountCCMConfigCmd)
}
