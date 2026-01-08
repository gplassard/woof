package case_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ArchiveCaseCmd = &cobra.Command{
	Use: "archive-case [case_id]",

	Short: "Archive case",
	Long: `Archive case
Documentation: https://docs.datadoghq.com/api/latest/case-management/#archive-case`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseResponse
		var err error

		var body datadogV2.CaseEmptyRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ArchiveCase(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to archive-case")

		fmt.Println(cmdutil.FormatJSON(res, "case_management"))
	},
}

func init() {

	ArchiveCaseCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ArchiveCaseCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ArchiveCaseCmd)
}
