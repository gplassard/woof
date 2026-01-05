package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateCaseDescriptionCmd = &cobra.Command{
	Use: "update-case-description [case_id] [payload]",

	Short: "Update case description",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseResponse
		var err error

		var body datadogV2.CaseUpdateDescriptionRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err = api.UpdateCaseDescription(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-case-description")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCaseDescriptionCmd)
}
