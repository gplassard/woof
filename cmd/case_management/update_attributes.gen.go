package case_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateAttributesCmd = &cobra.Command{
	Use: "update-attributes [case_id]",

	Short: "Update case attributes",
	Long: `Update case attributes
Documentation: https://docs.datadoghq.com/api/latest/case-management/#update-attributes`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseResponse
		var err error

		var body datadogV2.CaseUpdateAttributesRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateAttributes(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-attributes")

		fmt.Println(cmdutil.FormatJSON(res, "case"))
	},
}

func init() {

	UpdateAttributesCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateAttributesCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateAttributesCmd)
}
