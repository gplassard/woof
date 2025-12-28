package spa

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSPARecommendationsCmd = &cobra.Command{
	Use:   "get-s-p-a-recommendations [shard] [service]",
	
	Short: "Get SPA Recommendations",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSpaApi(client.NewAPIClient())
		res, _, err := api.GetSPARecommendations(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to get-s-p-a-recommendations: %v", err)
		}

		cmdutil.PrintJSON(res, "recommendation")
	},
}

func init() {
	Cmd.AddCommand(GetSPARecommendationsCmd)
}
