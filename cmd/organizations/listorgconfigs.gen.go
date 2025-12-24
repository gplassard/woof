package organizations

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListOrgConfigsCmd = &cobra.Command{
	Use:   "listorgconfigs",
	Short: "List Org Configs",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		res, _, err := api.ListOrgConfigs(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listorgconfigs: %v", err)
		}

		cmdutil.PrintJSON(res, "organizations")
	},
}

func init() {
	Cmd.AddCommand(ListOrgConfigsCmd)
}
