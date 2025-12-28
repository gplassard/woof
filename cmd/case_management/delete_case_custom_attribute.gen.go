package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCaseCustomAttributeCmd = &cobra.Command{
	Use:   "delete-case-custom-attribute [case_id] [custom_attribute_key]",
	Short: "Delete custom attribute from case",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.DeleteCaseCustomAttribute(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to delete-case-custom-attribute: %v", err)
		}

		cmdutil.PrintJSON(res, "case")
	},
}

func init() {
	Cmd.AddCommand(DeleteCaseCustomAttributeCmd)
}
