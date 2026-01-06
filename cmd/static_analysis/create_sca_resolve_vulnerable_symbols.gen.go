package static_analysis

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSCAResolveVulnerableSymbolsCmd = &cobra.Command{
	Use: "create-sca-resolve-vulnerable-symbols",

	Short: "POST request to resolve vulnerable symbols",
	Long: `POST request to resolve vulnerable symbols
Documentation: https://docs.datadoghq.com/api/latest/static-analysis/#create-sca-resolve-vulnerable-symbols`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ResolveVulnerableSymbolsResponse
		var err error

		var body datadogV2.ResolveVulnerableSymbolsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateSCAResolveVulnerableSymbols(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-sca-resolve-vulnerable-symbols")

		cmd.Println(cmdutil.FormatJSON(res, "resolve-vulnerable-symbols-response"))
	},
}

func init() {

	CreateSCAResolveVulnerableSymbolsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateSCAResolveVulnerableSymbolsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateSCAResolveVulnerableSymbolsCmd)
}
