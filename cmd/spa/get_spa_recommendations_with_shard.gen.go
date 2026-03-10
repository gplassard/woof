package spa

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSPARecommendationsWithShardCmd = &cobra.Command{
	Use:     "get-spa-recommendations-with-shard [shard] [service]",
	Aliases: []string{"get-recommendations-with-shard"},
	Short:   "Get SPA Recommendations with a shard parameter",
	Long: `Get SPA Recommendations with a shard parameter
Documentation: https://docs.datadoghq.com/api/latest/spa/#get-spa-recommendations-with-shard`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RecommendationDocument
		var err error

		api := datadogV2.NewSpaApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSPARecommendationsWithShard(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-spa-recommendations-with-shard")

		fmt.Println(cmdutil.FormatJSON(res, "recommendation"))
	},
}

func init() {

	Cmd.AddCommand(GetSPARecommendationsWithShardCmd)
}
