package entity_risk_scores

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListEntityRiskScoresCmd = &cobra.Command{
	Use:     "list-entity-risk-scores",
	Aliases: []string{"list"},
	Short:   "List Entity Risk Scores",
	Long: `List Entity Risk Scores
Documentation: https://docs.datadoghq.com/api/latest/entity-risk-scores/#list-entity-risk-scores`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityEntityRiskScoresResponse
		var err error

		api := datadogV2.NewEntityRiskScoresApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListEntityRiskScores(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-entity-risk-scores")

		fmt.Println(cmdutil.FormatJSON(res, "security_entity_risk_score"))
	},
}

func init() {

	Cmd.AddCommand(ListEntityRiskScoresCmd)
}
