package static_analysis

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateSCAResolveVulnerableSymbolsCmd = &cobra.Command{
	Use:   "createscaresolvevulnerablesymbols",
	Short: "POST request to resolve vulnerable symbols",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		res, _, err := api.CreateSCAResolveVulnerableSymbols(client.NewContext(apiKey, appKey, site), datadogV2.ResolveVulnerableSymbolsRequest{})
		if err != nil {
			log.Fatalf("failed to createscaresolvevulnerablesymbols: %v", err)
		}

		cmdutil.PrintJSON(res, "static_analysis")
	},
}

func init() {
	Cmd.AddCommand(CreateSCAResolveVulnerableSymbolsCmd)
}
