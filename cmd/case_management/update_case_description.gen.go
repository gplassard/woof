package case_management

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCaseDescriptionCmd = &cobra.Command{
	Use: "update-case-description [case_id]",

	Short: "Update case description",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCaseDescription(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseUpdateDescriptionRequest{})
		cmdutil.HandleError(err, "failed to update-case-description")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {
	Cmd.AddCommand(UpdateCaseDescriptionCmd)
}
