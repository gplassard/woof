package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CommentCaseCmd = &cobra.Command{
	Use: "comment-case [case_id] [payload]",

	Short: "Comment case",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TimelineResponse
		var err error

		var body datadogV2.CaseCommentRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err = api.CommentCase(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to comment-case")

		cmd.Println(cmdutil.FormatJSON(res, "timeline_cell"))
	},
}

func init() {
	Cmd.AddCommand(CommentCaseCmd)
}
