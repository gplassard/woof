package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateProjectCmd = &cobra.Command{
	Use: "create-project",

	Short: "Create a project",
	Long: `Create a project
Documentation: https://docs.datadoghq.com/api/latest/case-management/#create-project`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ProjectResponse
		var err error

		var body datadogV2.ProjectCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateProject(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-project")

		cmd.Println(cmdutil.FormatJSON(res, "project"))
	},
}

func init() {

	CreateProjectCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateProjectCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateProjectCmd)
}
