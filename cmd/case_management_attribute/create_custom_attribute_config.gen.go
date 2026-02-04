package case_management_attribute

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateCustomAttributeConfigCmd = &cobra.Command{
	Use: "create-custom-attribute-config [case_type_id]",

	Short: "Create custom attribute config for a case type",
	Long: `Create custom attribute config for a case type
Documentation: https://docs.datadoghq.com/api/latest/case-management-attribute/#create-custom-attribute-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomAttributeConfigResponse
		var err error

		var body datadogV2.CustomAttributeConfigCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCaseManagementAttributeApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateCustomAttributeConfig(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-custom-attribute-config")

		fmt.Println(cmdutil.FormatJSON(res, "custom_attribute_config"))
	},
}

func init() {

	CreateCustomAttributeConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateCustomAttributeConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateCustomAttributeConfigCmd)
}
