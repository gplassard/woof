package aws_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteAWSAccountCCMConfigCmd = &cobra.Command{
	Use: "delete-aws-account-ccm-config [aws_account_config_id]",

	Short: "Delete AWS CCM config",
	Long: `Delete AWS CCM config
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#delete-aws-account-ccm-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteAWSAccountCCMConfig(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-aws-account-ccm-config")

	},
}

func init() {

	Cmd.AddCommand(DeleteAWSAccountCCMConfigCmd)
}
