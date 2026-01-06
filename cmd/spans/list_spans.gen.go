package spans

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSpansCmd = &cobra.Command{
	Use:     "list-spans",
	Aliases: []string{"list"},
	Short:   "Search spans",
	Long: `Search spans
Documentation: https://docs.datadoghq.com/api/latest/spans/#list-spans`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SpansListResponse
		var err error

		var body datadogV2.SpansListRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSpansApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListSpans(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to list-spans")

		cmd.Println(cmdutil.FormatJSON(res, "spans"))
	},
}

func init() {

	ListSpansCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ListSpansCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ListSpansCmd)
}
