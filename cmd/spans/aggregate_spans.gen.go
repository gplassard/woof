package spans

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AggregateSpansCmd = &cobra.Command{
	Use:     "aggregate-spans",
	Aliases: []string{"aggregate"},
	Short:   "Aggregate spans",
	Long: `Aggregate spans
Documentation: https://docs.datadoghq.com/api/latest/spans/#aggregate-spans`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SpansAggregateResponse
		var err error

		var body datadogV2.SpansAggregateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSpansApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.AggregateSpans(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to aggregate-spans")

		fmt.Println(cmdutil.FormatJSON(res, "spans"))
	},
}

func init() {

	AggregateSpansCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AggregateSpansCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AggregateSpansCmd)
}
