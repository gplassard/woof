package on_call_paging

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOnCallPageCmd = &cobra.Command{
	Use: "create-on-call-page",

	Short: "Create On-Call Page",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOnCallPagingApi(client.NewAPIClient())
		res, _, err := api.CreateOnCallPage(client.NewContext(apiKey, appKey, site), datadogV2.CreatePageRequest{})
		cmdutil.HandleError(err, "failed to create-on-call-page")

		cmd.Println(cmdutil.FormatJSON(res, "pages"))
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallPageCmd)
}
