package case_management_type

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAllCaseTypesCmd = &cobra.Command{
	Use:   "get-all-case-types",
	
	Short: "Get all case types",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementTypeApi(client.NewAPIClient())
		res, _, err := api.GetAllCaseTypes(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-all-case-types")

		cmd.Println(cmdutil.FormatJSON(res, "case_type"))
	},
}

func init() {
	Cmd.AddCommand(GetAllCaseTypesCmd)
}
