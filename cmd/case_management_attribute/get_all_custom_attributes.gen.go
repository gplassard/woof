package case_management_attribute

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAllCustomAttributesCmd = &cobra.Command{
	Use: "get-all-custom-attributes",

	Short: "Get all custom attributes",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementAttributeApi(client.NewAPIClient())
		res, _, err := api.GetAllCustomAttributes(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-all-custom-attributes")

		cmd.Println(cmdutil.FormatJSON(res, "custom_attribute"))
	},
}

func init() {
	Cmd.AddCommand(GetAllCustomAttributesCmd)
}
