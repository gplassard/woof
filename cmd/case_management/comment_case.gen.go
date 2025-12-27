package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CommentCaseCmd = &cobra.Command{
	Use:   "comment_case [case_id]",
	Short: "Comment case",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.CommentCase(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseCommentRequest{})
		if err != nil {
			log.Fatalf("failed to comment_case: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management")
	},
}

func init() {
	Cmd.AddCommand(CommentCaseCmd)
}
