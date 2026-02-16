package aws_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAWSAccountsCmd = &cobra.Command{
	Use: "list-aws-accounts",

	Short: "List all AWS integrations",
	Long: `List all AWS integrations
Documentation: https://docs.datadoghq.com/api/latest/aws-integration/#list-aws-accounts`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AWSAccountsResponse
		var err error

		optionalParams := datadogV2.NewListAWSAccountsOptionalParameters()

		if cmd.Flags().Changed("aws-account-id") {
			val, _ := cmd.Flags().GetString("aws-account-id")
			optionalParams.WithAwsAccountId(val)
		}

		api := datadogV2.NewAWSIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAWSAccounts(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-aws-accounts")

		fmt.Println(cmdutil.FormatJSON(res, "account"))
	},
}

func init() {

	ListAWSAccountsCmd.Flags().String("aws-account-id", "", "Optional query parameter to filter accounts by AWS Account ID. If not provided, all accounts are returned.")

	Cmd.AddCommand(ListAWSAccountsCmd)
}
