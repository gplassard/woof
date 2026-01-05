package on_call_paging

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var AcknowledgeOnCallPageCmd = &cobra.Command{
	Use: "acknowledge-on-call-page [page_id]",

	Short: "Acknowledge On-Call Page",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		_, err := api.AcknowledgeOnCallPage(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to acknowledge-on-call-page")

	},
}

func init() {
	Cmd.AddCommand(AcknowledgeOnCallPageCmd)
}
