package dora_metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDORAIncidentCmd = &cobra.Command{
	Use: "create-d-o-r-ai-ncident",

	Short: "Send an incident event",
	Long: `Send an incident event
Documentation: https://docs.datadoghq.com/api/latest/dora-metrics/#create-d-o-r-ai-ncident`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORAFailureResponse
		var err error

		var body datadogV2.DORAFailureRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateDORAIncident(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-d-o-r-ai-ncident")

		fmt.Println(cmdutil.FormatJSON(res, "dora_failure"))
	},
}

func init() {

	CreateDORAIncidentCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDORAIncidentCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDORAIncidentCmd)
}
