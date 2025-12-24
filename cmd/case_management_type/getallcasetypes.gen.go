package case_management_type

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAllCaseTypesCmd = &cobra.Command{
	Use:   "getallcasetypes",
	Short: "Get all case types",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCaseManagementTypeApi(client.NewAPIClient())
		res, _, err := api.GetAllCaseTypes(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getallcasetypes: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management_type")
	},
}

func init() {
	Cmd.AddCommand(GetAllCaseTypesCmd)
}
