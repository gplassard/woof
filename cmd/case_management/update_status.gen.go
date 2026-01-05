package case_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateStatusCmd = &cobra.Command{
	Use: "update-status [case_id]",

	Short: "Update case status",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateStatus(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseUpdateStatusRequest{})
		cmdutil.HandleError(err, "failed to update-status")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {
	Cmd.AddCommand(UpdateStatusCmd)
}
