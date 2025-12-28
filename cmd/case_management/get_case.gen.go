package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetCaseCmd = &cobra.Command{
	Use:   "get-case [case_id]",
	
	Short: "Get the details of a case",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.GetCase(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-case: %v", err)
		}

		cmdutil.PrintJSON(res, "case")
	},
}

func init() {
	Cmd.AddCommand(GetCaseCmd)
}
