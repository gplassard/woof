package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ValidateMonitorUserTemplateCmd = &cobra.Command{
	Use:   "validatemonitorusertemplate",
	Short: "Validate a monitor user template",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err := api.ValidateMonitorUserTemplate(client.NewContext(apiKey, appKey, site), datadogV2.MonitorUserTemplateCreateRequest{})
		if err != nil {
			log.Fatalf("failed to validatemonitorusertemplate: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(ValidateMonitorUserTemplateCmd)
}
