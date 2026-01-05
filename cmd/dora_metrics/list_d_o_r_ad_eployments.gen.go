package dora_metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ListDORADeploymentsCmd = &cobra.Command{
	Use: "list-d-o-r-ad-eployments [payload]",

	Short: "Get a list of deployment events",
	Long: `Get a list of deployment events
Documentation: https://docs.datadoghq.com/api/latest/dora-metrics/#list-d-o-r-ad-eployments`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DORAListResponse
		var err error

		var body datadogV2.DORAListDeploymentsRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err = api.ListDORADeployments(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to list-d-o-r-ad-eployments")

		cmd.Println(cmdutil.FormatJSON(res, "dora_metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListDORADeploymentsCmd)
}
