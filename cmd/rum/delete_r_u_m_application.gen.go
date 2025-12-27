package rum

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteRUMApplicationCmd = &cobra.Command{
	Use:   "delete_r_u_m_application [id]",
	Short: "Delete a RUM application",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		_, err := api.DeleteRUMApplication(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_r_u_m_application: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteRUMApplicationCmd)
}
