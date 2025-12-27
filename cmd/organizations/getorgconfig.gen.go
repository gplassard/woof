package organizations

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetOrgConfigCmd = &cobra.Command{
	Use:   "getorgconfig [org_config_name]",
	Short: "Get a specific Org Config value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOrganizationsApi(client.NewAPIClient())
		res, _, err := api.GetOrgConfig(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getorgconfig: %v", err)
		}

		cmdutil.PrintJSON(res, "organizations")
	},
}

func init() {
	Cmd.AddCommand(GetOrgConfigCmd)
}
