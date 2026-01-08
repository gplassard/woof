package gcp_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListGCPSTSAccountsCmd = &cobra.Command{
	Use: "list-gcp-sts-accounts",

	Short: "List all GCP STS-enabled service accounts",
	Long: `List all GCP STS-enabled service accounts
Documentation: https://docs.datadoghq.com/api/latest/g-c-p-integration/#list-gcp-sts-accounts`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GCPSTSServiceAccountsResponse
		var err error

		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListGCPSTSAccounts(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-gcp-sts-accounts")

		fmt.Println(cmdutil.FormatJSON(res, "g_c_p_s_t_s_account"))
	},
}

func init() {

	Cmd.AddCommand(ListGCPSTSAccountsCmd)
}
