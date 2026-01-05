package case_management_attribute

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCustomAttributeConfigCmd = &cobra.Command{
	Use: "delete-custom-attribute-config [case_type_id] [custom_attribute_id]",

	Short: "Delete custom attributes config",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewCaseManagementAttributeApi(client.NewAPIClient())
		_, err := api.DeleteCustomAttributeConfig(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-custom-attribute-config")

	},
}

func init() {
	Cmd.AddCommand(DeleteCustomAttributeConfigCmd)
}
