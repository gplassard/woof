package static_analysis

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateSCAResolveVulnerableSymbolsCmd = &cobra.Command{
	Use: "create-sca-resolve-vulnerable-symbols [payload]",

	Short: "POST request to resolve vulnerable symbols",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ResolveVulnerableSymbolsResponse
		var err error

		var body datadogV2.ResolveVulnerableSymbolsRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewStaticAnalysisApi(client.NewAPIClient())
		res, _, err = api.CreateSCAResolveVulnerableSymbols(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-sca-resolve-vulnerable-symbols")

		cmd.Println(cmdutil.FormatJSON(res, "resolve-vulnerable-symbols-response"))
	},
}

func init() {
	Cmd.AddCommand(CreateSCAResolveVulnerableSymbolsCmd)
}
