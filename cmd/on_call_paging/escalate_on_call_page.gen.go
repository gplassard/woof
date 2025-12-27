package on_call_paging

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var EscalateOnCallPageCmd = &cobra.Command{
	Use:   "escalate_on_call_page [page_id]",
	Short: "Escalate On-Call Page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		_, err := api.EscalateOnCallPage(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatalf("failed to escalate_on_call_page: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(EscalateOnCallPageCmd)
}
