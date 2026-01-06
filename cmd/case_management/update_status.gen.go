package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateStatusCmd = &cobra.Command{
	Use: "update-status [case_id]",

	Short: "Update case status",
	Long: `Update case status
Documentation: https://docs.datadoghq.com/api/latest/case-management/#update-status`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseResponse
		var err error

		var body datadogV2.CaseUpdateStatusRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateStatus(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-status")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {

	UpdateStatusCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateStatusCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateStatusCmd)
}
