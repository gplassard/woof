package case_management_attribute

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAllCustomAttributeConfigsByCaseTypeCmd = &cobra.Command{
	Use:   "getallcustomattributeconfigsbycasetype [case_type_id]",
	Short: "Get all custom attributes config of case type",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCaseManagementAttributeApi(client.NewAPIClient())
		res, _, err := api.GetAllCustomAttributeConfigsByCaseType(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getallcustomattributeconfigsbycasetype: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management_attribute")
	},
}

func init() {
	Cmd.AddCommand(GetAllCustomAttributeConfigsByCaseTypeCmd)
}
