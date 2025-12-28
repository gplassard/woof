package gcp_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListGCPSTSAccountsCmd = &cobra.Command{
	Use:   "list_g_c_p_s_t_s_accounts",
	Short: "List all GCP STS-enabled service accounts",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListGCPSTSAccounts(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_g_c_p_s_t_s_accounts: %v", err)
		}

		cmdutil.PrintJSON(res, "gcp_service_account")
	},
}

func init() {
	Cmd.AddCommand(ListGCPSTSAccountsCmd)
}
