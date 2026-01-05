package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCaseDescriptionCmd = &cobra.Command{
	Use: "update-case-description [case_id]",

	Short: "Update case description",
	Long: `Update case description
Documentation: https://docs.datadoghq.com/api/latest/case-management/#update-case-description`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseResponse
		var err error

		var body datadogV2.CaseUpdateDescriptionRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err = api.UpdateCaseDescription(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-case-description")

		cmd.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {

	UpdateCaseDescriptionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCaseDescriptionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCaseDescriptionCmd)
}
