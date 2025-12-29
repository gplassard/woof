package case_management_attribute

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCustomAttributeConfigCmd = &cobra.Command{
	Use:   "create-custom-attribute-config [case_type_id]",
	
	Short: "Create custom attribute config for a case type",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementAttributeApi(client.NewAPIClient())
		res, _, err := api.CreateCustomAttributeConfig(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CustomAttributeConfigCreateRequest{})
		cmdutil.HandleError(err, "failed to create-custom-attribute-config")

		cmd.Println(cmdutil.FormatJSON(res, "custom_attribute"))
	},
}

func init() {
	Cmd.AddCommand(CreateCustomAttributeConfigCmd)
}
