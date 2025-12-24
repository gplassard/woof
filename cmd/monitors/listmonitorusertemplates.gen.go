package monitors

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListMonitorUserTemplatesCmd = &cobra.Command{
	Use:   "listmonitorusertemplates",
	Short: "Get all monitor user templates",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.ListMonitorUserTemplates(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listmonitorusertemplates: %v", err)
		}

		cmdutil.PrintJSON(res, "monitors")
	},
}

func init() {
	Cmd.AddCommand(ListMonitorUserTemplatesCmd)
}
