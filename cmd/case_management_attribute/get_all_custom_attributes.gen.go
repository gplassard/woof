package case_management_attribute

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAllCustomAttributesCmd = &cobra.Command{
	Use:   "get_all_custom_attributes",
	Short: "Get all custom attributes",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementAttributeApi(client.NewAPIClient())
		res, _, err := api.GetAllCustomAttributes(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get_all_custom_attributes: %v", err)
		}

		cmdutil.PrintJSON(res, "custom_attribute")
	},
}

func init() {
	Cmd.AddCommand(GetAllCustomAttributesCmd)
}
