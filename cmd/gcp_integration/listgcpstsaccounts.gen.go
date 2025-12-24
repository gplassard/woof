package gcp_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListGCPSTSAccountsCmd = &cobra.Command{
	Use:   "listgcpstsaccounts",
	Short: "List all GCP STS-enabled service accounts",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListGCPSTSAccounts(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listgcpstsaccounts: %v", err)
		}

		cmdutil.PrintJSON(res, "gcp_integration")
	},
}

func init() {
	Cmd.AddCommand(ListGCPSTSAccountsCmd)
}
