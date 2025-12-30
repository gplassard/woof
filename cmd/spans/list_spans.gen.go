package spans

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSpansCmd = &cobra.Command{
	Use:     "list-spans",
	Aliases: []string{"list"},
	Short:   "Search spans",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpansApi(client.NewAPIClient())
		res, _, err := api.ListSpans(client.NewContext(apiKey, appKey, site), datadogV2.SpansListRequest{})
		cmdutil.HandleError(err, "failed to list-spans")

		cmd.Println(cmdutil.FormatJSON(res, "spans"))
	},
}

func init() {
	Cmd.AddCommand(ListSpansCmd)
}
