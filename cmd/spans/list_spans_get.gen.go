package spans

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSpansGetCmd = &cobra.Command{
	Use:     "list-spans-get",
	Aliases: []string{"list-get"},
	Short:   "Get a list of spans",
	Long: `Get a list of spans
Documentation: https://docs.datadoghq.com/api/latest/spans/#list-spans-get`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SpansListResponse
		var err error

		api := datadogV2.NewSpansApi(client.NewAPIClient())
		res, _, err = api.ListSpansGet(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-spans-get")

		cmd.Println(cmdutil.FormatJSON(res, "spans"))
	},
}

func init() {

	Cmd.AddCommand(ListSpansGetCmd)
}
