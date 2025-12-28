package case_management_type

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCaseTypeCmd = &cobra.Command{
	Use:   "delete-case-type [case_type_id]",
	Short: "Delete a case type",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementTypeApi(client.NewAPIClient())
		_, err := api.DeleteCaseType(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete-case-type: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCaseTypeCmd)
}
