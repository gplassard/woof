package status_pages

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var GetDegradationCmd = &cobra.Command{
	Use: "get-degradation [page_id] [degradation_id]",

	Short: "Get degradation",
	Long: `Get degradation
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#get-degradation`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Degradation
		var err error

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDegradation(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), uuid.MustParse(args[1]))
		cmdutil.HandleError(err, "failed to get-degradation")

		fmt.Println(cmdutil.FormatJSON(res, "degradations"))
	},
}

func init() {

	Cmd.AddCommand(GetDegradationCmd)
}
