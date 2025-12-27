package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ValidateExistingMonitorUserTemplateCmd = &cobra.Command{
	Use:   "validateexistingmonitorusertemplate [template_id]",
	Short: "Validate an existing monitor user template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		_, err := api.ValidateExistingMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MonitorUserTemplateUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to validateexistingmonitorusertemplate: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(ValidateExistingMonitorUserTemplateCmd)
}
