package change_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateChangeRequestDecisionCmd = &cobra.Command{
	Use: "update-change-request-decision [change_request_id] [decision_id]",

	Short: "Update a change request decision",
	Long: `Update a change request decision
Documentation: https://docs.datadoghq.com/api/latest/change-management/#update-change-request-decision`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ChangeRequestResponse
		var err error

		var body datadogV2.ChangeRequestDecisionUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewChangeManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateChangeRequestDecision(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-change-request-decision")

		fmt.Println(cmdutil.FormatJSON(res, "change_request"))
	},
}

func init() {

	UpdateChangeRequestDecisionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateChangeRequestDecisionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateChangeRequestDecisionCmd)
}
