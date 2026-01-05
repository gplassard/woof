package spa

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSPARecommendationsCmd = &cobra.Command{
	Use: "get-s-p-a-recommendations [shard] [service]",

	Short: "Get SPA Recommendations",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RecommendationDocument
		var err error

		api := datadogV2.NewSpaApi(client.NewAPIClient())
		res, _, err = api.GetSPARecommendations(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-s-p-a-recommendations")

		cmd.Println(cmdutil.FormatJSON(res, "recommendation"))
	},
}

func init() {
	Cmd.AddCommand(GetSPARecommendationsCmd)
}
