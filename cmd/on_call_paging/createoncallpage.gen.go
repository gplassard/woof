package on_call_paging

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOnCallPageCmd = &cobra.Command{
	Use:   "createoncallpage",
	Short: "Create On-Call Page",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		res, _, err := api.CreateOnCallPage(client.NewContext(apiKey, appKey, site), datadogV2.CreatePageRequest{})
		if err != nil {
			log.Fatalf("failed to createoncallpage: %v", err)
		}

		cmdutil.PrintJSON(res, "on_call_paging")
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallPageCmd)
}
