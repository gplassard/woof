package case_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var SearchCasesCmd = &cobra.Command{
	Use:   "searchcases",
	Short: "Search cases",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.SearchCases(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to searchcases: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management")
	},
}

func init() {
	Cmd.AddCommand(SearchCasesCmd)
}
