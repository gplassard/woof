package change_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateChangeRequestCmd = &cobra.Command{
	Use: "update-change-request [change_request_id]",

	Short: "Update a change request",
	Long: `Update a change request
Documentation: https://docs.datadoghq.com/api/latest/change-management/#update-change-request`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ChangeRequestResponse
		var err error

		var body datadogV2.ChangeRequestUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewChangeManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateChangeRequest(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-change-request")

		fmt.Println(cmdutil.FormatJSON(res, "change_request"))
	},
}

func init() {

	UpdateChangeRequestCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateChangeRequestCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateChangeRequestCmd)
}
