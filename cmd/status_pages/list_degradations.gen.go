package status_pages

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDegradationsCmd = &cobra.Command{
	Use: "list-degradations",

	Short: "List degradations",
	Long: `List degradations
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#list-degradations`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DegradationArray
		var err error

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDegradations(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-degradations")

		fmt.Println(cmdutil.FormatJSON(res, "degradations"))
	},
}

func init() {

	Cmd.AddCommand(ListDegradationsCmd)
}
