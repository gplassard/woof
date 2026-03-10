package spa

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSPARecommendationsCmd = &cobra.Command{
	Use:     "get-spa-recommendations [service]",
	Aliases: []string{"get-recommendations"},
	Short:   "Get SPA Recommendations",
	Long: `Get SPA Recommendations
Documentation: https://docs.datadoghq.com/api/latest/spa/#get-spa-recommendations`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RecommendationDocument
		var err error

		api := datadogV2.NewSpaApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSPARecommendations(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-spa-recommendations")

		fmt.Println(cmdutil.FormatJSON(res, "recommendation"))
	},
}

func init() {

	Cmd.AddCommand(GetSPARecommendationsCmd)
}
