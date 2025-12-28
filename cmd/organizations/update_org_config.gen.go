package organizations

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateOrgConfigCmd = &cobra.Command{
	Use:   "update_org_config [org_config_name]",
	Short: "Update a specific Org Config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		res, _, err := api.UpdateOrgConfig(client.NewContext(apiKey, appKey, site), args[0], datadogV2.OrgConfigWriteRequest{})
		if err != nil {
			log.Fatalf("failed to update_org_config: %v", err)
		}

		cmdutil.PrintJSON(res, "org_configs")
	},
}

func init() {
	Cmd.AddCommand(UpdateOrgConfigCmd)
}
