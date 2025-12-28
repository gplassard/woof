package static_analysis

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateSCAResolveVulnerableSymbolsCmd = &cobra.Command{
	Use:   "create_s_c_a_resolve_vulnerable_symbols",
	Short: "POST request to resolve vulnerable symbols",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		res, _, err := api.CreateSCAResolveVulnerableSymbols(client.NewContext(apiKey, appKey, site), datadogV2.ResolveVulnerableSymbolsRequest{})
		if err != nil {
			log.Fatalf("failed to create_s_c_a_resolve_vulnerable_symbols: %v", err)
		}

		cmdutil.PrintJSON(res, "resolve-vulnerable-symbols-response")
	},
}

func init() {
	Cmd.AddCommand(CreateSCAResolveVulnerableSymbolsCmd)
}
