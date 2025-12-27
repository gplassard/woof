package on_call_paging

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOnCallPageCmd = &cobra.Command{
	Use:   "create_on_call_page",
	Short: "Create On-Call Page",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		res, _, err := api.CreateOnCallPage(client.NewContext(apiKey, appKey, site), datadogV2.CreatePageRequest{})
		if err != nil {
			log.Fatalf("failed to create_on_call_page: %v", err)
		}

		cmdutil.PrintJSON(res, "on_call_paging")
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallPageCmd)
}
