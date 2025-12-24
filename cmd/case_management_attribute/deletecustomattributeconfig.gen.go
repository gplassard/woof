package case_management_attribute

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCustomAttributeConfigCmd = &cobra.Command{
	Use:   "deletecustomattributeconfig [case_type_id] [custom_attribute_id]",
	Short: "Delete custom attributes config",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCaseManagementAttributeApi(client.NewAPIClient())
		_, err := api.DeleteCustomAttributeConfig(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to deletecustomattributeconfig: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCustomAttributeConfigCmd)
}
