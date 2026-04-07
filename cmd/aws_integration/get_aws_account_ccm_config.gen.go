package aws_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAWSAccountCCMConfigCmd = &cobra.Command{
	Use: "get-aws-account-ccm-config [aws_account_config_id]",

	Short: "Get AWS CCM config",
	Long: `Get AWS CCM config
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#get-aws-account-ccm-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSCcmConfigResponse
		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAWSAccountCCMConfig(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-aws-account-ccm-config")

		fmt.Println(cmdutil.FormatJSON(res, "ccm_config"))
	},
}

func init() {

	Cmd.AddCommand(GetAWSAccountCCMConfigCmd)
}
