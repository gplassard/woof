package change_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateChangeRequestCmd = &cobra.Command{
	Use: "create-change-request",

	Short: "Create a change request",
	Long: `Create a change request
Documentation: https://docs.datadoghq.com/api/latest/change-management/#create-change-request`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ChangeRequestResponse
		var err error

		var body datadogV2.ChangeRequestCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewChangeManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateChangeRequest(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-change-request")

		fmt.Println(cmdutil.FormatJSON(res, "change_request"))
	},
}

func init() {

	CreateChangeRequestCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateChangeRequestCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateChangeRequestCmd)
}
