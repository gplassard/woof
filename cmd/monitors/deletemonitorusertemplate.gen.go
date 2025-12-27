package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteMonitorUserTemplateCmd = &cobra.Command{
	Use:   "deletemonitorusertemplate [template_id]",
	Short: "Delete a monitor user template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err := api.DeleteMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletemonitorusertemplate: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteMonitorUserTemplateCmd)
}
