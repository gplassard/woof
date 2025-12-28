package on_call_paging

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
	
	
	"github.com/spf13/cobra"
	
)

var ResolveOnCallPageCmd = &cobra.Command{
	Use:   "resolve-on-call-page [page_id]",
	
	Short: "Resolve On-Call Page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		_, err := api.ResolveOnCallPage(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to resolve-on-call-page")

		
	},
}

func init() {
	Cmd.AddCommand(ResolveOnCallPageCmd)
}
