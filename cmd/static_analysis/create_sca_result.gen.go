package static_analysis

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSCAResultCmd = &cobra.Command{
	Use: "create-sca-result",

	Short: "Post dependencies for analysis",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		_, err := api.CreateSCAResult(client.NewContext(apiKey, appKey, site), datadogV2.ScaRequest{})
		cmdutil.HandleError(err, "failed to create-sca-result")

	},
}

func init() {
	Cmd.AddCommand(CreateSCAResultCmd)
}
