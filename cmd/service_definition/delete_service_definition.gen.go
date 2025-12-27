package service_definition

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteServiceDefinitionCmd = &cobra.Command{
	Use:   "delete_service_definition [service_name]",
	Short: "Delete a single service definition",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		_, err := api.DeleteServiceDefinition(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_service_definition: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteServiceDefinitionCmd)
}
