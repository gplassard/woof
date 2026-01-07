package okta_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListOktaAccountsCmd = &cobra.Command{
	Use: "list-okta-accounts",

	Short: "List Okta accounts",
	Long: `List Okta accounts
Documentation: https://docs.datadoghq.com/api/latest/okta-integration/#list-okta-accounts`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OktaAccountsResponse
		var err error

		api := datadogV2.NewOktaIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListOktaAccounts(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-okta-accounts")

		fmt.Println(cmdutil.FormatJSON(res, "okta-accounts"))
	},
}

func init() {

	Cmd.AddCommand(ListOktaAccountsCmd)
}
