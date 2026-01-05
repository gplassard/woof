package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdatePriorityCmd = &cobra.Command{
	Use: "update-priority [case_id]",

	Short: "Update case priority",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UpdatePriority(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseUpdatePriorityRequest{})
		cmdutil.HandleError(err, "failed to update-priority")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {
	Cmd.AddCommand(UpdatePriorityCmd)
}
