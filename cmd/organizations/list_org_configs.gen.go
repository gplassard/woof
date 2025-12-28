package organizations

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListOrgConfigsCmd = &cobra.Command{
	Use:   "list-org-configs",
	
	Short: "List Org Configs",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		res, _, err := api.ListOrgConfigs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-org-configs")

		cmdutil.PrintJSON(res, "org_configs")
	},
}

func init() {
	Cmd.AddCommand(ListOrgConfigsCmd)
}
