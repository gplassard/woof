package static_analysis

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateSCAResultCmd = &cobra.Command{
	Use: "create-sca-result [payload]",

	Short: "Post dependencies for analysis",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.ScaRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		_, err = api.CreateSCAResult(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-sca-result")

	},
}

func init() {
	Cmd.AddCommand(CreateSCAResultCmd)
}
