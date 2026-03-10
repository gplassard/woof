package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCaseNotebookCmd = &cobra.Command{
	Use: "create-case-notebook [case_id]",

	Short: "Create investigation notebook for case",
	Long: `Create investigation notebook for case
Documentation: https://docs.datadoghq.com/api/latest/case-management/#create-case-notebook`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.NotebookCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.CreateCaseNotebook(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-case-notebook")

	},
}

func init() {

	CreateCaseNotebookCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCaseNotebookCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCaseNotebookCmd)
}
