package static_analysis

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSCAResultCmd = &cobra.Command{
	Use: "create-sca-result",

	Short: "",
	Long: `
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#create-sca-result`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.ScaRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.CreateSCAResult(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-sca-result")

	},
}

func init() {

	CreateSCAResultCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSCAResultCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSCAResultCmd)
}
