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

var GetStatusPageCmd = &cobra.Command{
	Use:     "get-status-page [page_id]",
	Aliases: []string{"get"},
	Short:   "Get status page",
	Long: `Get status page
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#get-status-page`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.StatusPage
		var err error

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetStatusPage(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to get-status-page")

		fmt.Println(cmdutil.FormatJSON(res, "status_pages"))
	},
}

func init() {

	Cmd.AddCommand(GetStatusPageCmd)
}
