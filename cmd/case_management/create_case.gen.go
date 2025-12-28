package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCaseCmd = &cobra.Command{
	Use:   "create-case",
	
	Short: "Create a case",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCase(client.NewContext(apiKey, appKey, site), datadogV2.CaseCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-case: %v", err)
		}

		cmdutil.PrintJSON(res, "case")
	},
}

func init() {
	Cmd.AddCommand(CreateCaseCmd)
}
