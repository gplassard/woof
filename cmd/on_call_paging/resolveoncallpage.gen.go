package on_call_paging

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var ResolveOnCallPageCmd = &cobra.Command{
	Use:   "resolveoncallpage [page_id]",
	Short: "Resolve On-Call Page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		_, err := api.ResolveOnCallPage(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatalf("failed to resolveoncallpage: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(ResolveOnCallPageCmd)
}
