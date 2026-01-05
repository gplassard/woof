package gcp_integration

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListGCPSTSAccountsCmd = &cobra.Command{
	Use: "list-gcp-sts-accounts",

	Short: "List all GCP STS-enabled service accounts",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewGCPIntegrationApi(client.NewAPIClient())
		res, _, err := api.ListGCPSTSAccounts(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-gcp-sts-accounts")

		cmd.Println(cmdutil.FormatJSON(res, "gcp_service_account"))
	},
}

func init() {
	Cmd.AddCommand(ListGCPSTSAccountsCmd)
}
