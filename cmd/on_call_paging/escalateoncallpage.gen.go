package on_call_paging

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var EscalateOnCallPageCmd = &cobra.Command{
	Use:   "escalateoncallpage [page_id]",
	Short: "Escalate On-Call Page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		_, err := api.EscalateOnCallPage(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatalf("failed to escalateoncallpage: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(EscalateOnCallPageCmd)
}
