package static_analysis

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateSCAResolveVulnerableSymbolsCmd = &cobra.Command{
	Use: "create-sca-resolve-vulnerable-symbols",

	Short: "POST request to resolve vulnerable symbols",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		res, _, err := api.CreateSCAResolveVulnerableSymbols(client.NewContext(apiKey, appKey, site), datadogV2.ResolveVulnerableSymbolsRequest{})
		cmdutil.HandleError(err, "failed to create-sca-resolve-vulnerable-symbols")

		cmd.Println(cmdutil.FormatJSON(res, "resolve-vulnerable-symbols-response"))
	},
}

func init() {
	Cmd.AddCommand(CreateSCAResolveVulnerableSymbolsCmd)
}
