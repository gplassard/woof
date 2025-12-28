package case_management_type

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAllCaseTypesCmd = &cobra.Command{
	Use:   "get_all_case_types",
	Short: "Get all case types",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementTypeApi(client.NewAPIClient())
		res, _, err := api.GetAllCaseTypes(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get_all_case_types: %v", err)
		}

		cmdutil.PrintJSON(res, "case_type")
	},
}

func init() {
	Cmd.AddCommand(GetAllCaseTypesCmd)
}
