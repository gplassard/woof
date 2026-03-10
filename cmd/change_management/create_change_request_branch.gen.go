package change_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateChangeRequestBranchCmd = &cobra.Command{
	Use: "create-change-request-branch [change_request_id]",

	Short: "Create a change request branch",
	Long: `Create a change request branch
Documentation: https://docs.datadoghq.com/api/latest/change-management/#create-change-request-branch`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ChangeRequestResponse
		var err error

		var body datadogV2.ChangeRequestBranchCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewChangeManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateChangeRequestBranch(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-change-request-branch")

		fmt.Println(cmdutil.FormatJSON(res, "change_request"))
	},
}

func init() {

	CreateChangeRequestBranchCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateChangeRequestBranchCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateChangeRequestBranchCmd)
}
