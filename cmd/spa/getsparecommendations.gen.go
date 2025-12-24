package spa

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetSPARecommendationsCmd = &cobra.Command{
	Use:   "getsparecommendations",
	Short: "Get SPA Recommendations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/spa/recommendations/{service}/{shard}")
		fmt.Println("OperationID: GetSPARecommendations")
	},
}

func init() {
	Cmd.AddCommand(GetSPARecommendationsCmd)
}
