package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateCaseCustomAttributeCmd = &cobra.Command{
	Use:   "updatecasecustomattribute [case_id] [custom_attribute_key]",
	Short: "Update case custom attribute",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCaseCustomAttribute(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.CaseUpdateCustomAttributeRequest{})
		if err != nil {
			log.Fatalf("failed to updatecasecustomattribute: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management")
	},
}

func init() {
	Cmd.AddCommand(UpdateCaseCustomAttributeCmd)
}
