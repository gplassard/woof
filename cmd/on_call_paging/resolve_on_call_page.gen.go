package on_call_paging

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var ResolveOnCallPageCmd = &cobra.Command{
	Use: "resolve-on-call-page [page_id]",

	Short: "Resolve On-Call Page",
	Long: `Resolve On-Call Page
Documentation: https://docs.datadoghq.com/api/latest/on-call-paging/#resolve-on-call-page`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.ResolveOnCallPage(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to resolve-on-call-page")

	},
}

func init() {

	Cmd.AddCommand(ResolveOnCallPageCmd)
}
