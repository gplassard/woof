package aws_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAWSAccountCmd = &cobra.Command{
	Use: "get-aws-account [aws_account_config_id]",

	Short: "Get an AWS integration by config ID",
	Long: `Get an AWS integration by config ID
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#get-aws-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSAccountResponse
		var err error

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAWSAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-aws-account")

		fmt.Println(cmdutil.FormatJSON(res, "account"))
	},
}

func init() {

	Cmd.AddCommand(GetAWSAccountCmd)
}
