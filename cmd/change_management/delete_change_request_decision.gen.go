package change_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteChangeRequestDecisionCmd = &cobra.Command{
	Use: "delete-change-request-decision [change_request_id] [decision_id]",

	Short: "Delete a change request decision",
	Long: `Delete a change request decision
Documentation: https://docs.datadoghq.com/api/latest/change-management/#delete-change-request-decision`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ChangeRequestResponse
		var err error

		api := datadogV2.NewChangeManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.DeleteChangeRequestDecision(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-change-request-decision")

		fmt.Println(cmdutil.FormatJSON(res, "change_request"))
	},
}

func init() {

	Cmd.AddCommand(DeleteChangeRequestDecisionCmd)
}
