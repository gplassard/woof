package case_management_type

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCaseTypeCmd = &cobra.Command{
	Use:   "createcasetype",
	Short: "Create a case type",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCaseManagementTypeApi(client.NewAPIClient())
		res, _, err := api.CreateCaseType(client.NewContext(apiKey, appKey, site), datadogV2.CaseTypeCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createcasetype: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management_type")
	},
}

func init() {
	Cmd.AddCommand(CreateCaseTypeCmd)
}
