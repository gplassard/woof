package case_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCaseCommentCmd = &cobra.Command{
	Use:   "deletecasecomment [case_id] [cell_id]",
	Short: "Delete case comment",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		_, err := api.DeleteCaseComment(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to deletecasecomment: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCaseCommentCmd)
}
