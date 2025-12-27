package api_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var DeleteOpenAPICmd = &cobra.Command{
	Use:   "delete_open_a_p_i [id]",
	Short: "Delete an API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		_, err := api.DeleteOpenAPI(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatalf("failed to delete_open_a_p_i: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteOpenAPICmd)
}
