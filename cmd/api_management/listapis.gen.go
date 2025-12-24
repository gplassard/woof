package api_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAPIsCmd = &cobra.Command{
	Use:   "listapis",
	Short: "List APIs",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		res, _, err := api.ListAPIs(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listapis: %v", err)
		}

		cmdutil.PrintJSON(res, "api_management")
	},
}

func init() {
	Cmd.AddCommand(ListAPIsCmd)
}
