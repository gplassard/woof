package case_management_attribute

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAllCustomAttributeConfigsByCaseTypeCmd = &cobra.Command{
	Use: "get-all-custom-attribute-configs-by-case-type [case_type_id]",

	Short: "Get all custom attributes config of case type",
	Long: `Get all custom attributes config of case type
Documentation: https://docs.datadoghq.com/api/latest/case-management-attribute/#get-all-custom-attribute-configs-by-case-type`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomAttributeConfigsResponse
		var err error

		api := datadogV2.NewCaseManagementAttributeApi(client.NewAPIClient())
		res, _, err = api.GetAllCustomAttributeConfigsByCaseType(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-all-custom-attribute-configs-by-case-type")

		cmd.Println(cmdutil.FormatJSON(res, "custom_attribute"))
	},
}

func init() {

	Cmd.AddCommand(GetAllCustomAttributeConfigsByCaseTypeCmd)
}
