package change_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetChangeRequestCmd = &cobra.Command{
	Use: "get-change-request [change_request_id]",

	Short: "Get a change request",
	Long: `Get a change request
Documentation: https://docs.datadoghq.com/api/latest/change-management/#get-change-request`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ChangeRequestResponse
		var err error

		api := datadogV2.NewChangeManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetChangeRequest(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-change-request")

		fmt.Println(cmdutil.FormatJSON(res, "change_request"))
	},
}

func init() {

	Cmd.AddCommand(GetChangeRequestCmd)
}
