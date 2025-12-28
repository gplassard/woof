package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCaseCommentCmd = &cobra.Command{
	Use:   "delete-case-comment [case_id] [cell_id]",
	
	Short: "Delete case comment",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		_, err := api.DeleteCaseComment(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to delete-case-comment: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCaseCommentCmd)
}
