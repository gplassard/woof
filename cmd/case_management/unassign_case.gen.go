package case_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UnassignCaseCmd = &cobra.Command{
	Use: "unassign-case [case_id]",

	Short: "Unassign case",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UnassignCase(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseEmptyRequest{})
		cmdutil.HandleError(err, "failed to unassign-case")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {
	Cmd.AddCommand(UnassignCaseCmd)
}
