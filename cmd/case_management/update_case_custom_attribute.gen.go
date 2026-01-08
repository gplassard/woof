package case_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateCaseCustomAttributeCmd = &cobra.Command{
	Use: "update-case-custom-attribute [case_id] [custom_attribute_key]",

	Short: "Update case custom attribute",
	Long: `Update case custom attribute
Documentation: https://docs.datadoghq.com/api/latest/case-management/#update-case-custom-attribute`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CaseResponse
		var err error

		var body datadogV2.CaseUpdateCustomAttributeRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateCaseCustomAttribute(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-case-custom-attribute")

		fmt.Println(cmdutil.FormatJSON(res, "case_custom_attribute"))
	},
}

func init() {

	UpdateCaseCustomAttributeCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateCaseCustomAttributeCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateCaseCustomAttributeCmd)
}
