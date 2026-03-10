package status_pages

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteStatusPageCmd = &cobra.Command{
	Use:     "delete-status-page [page_id]",
	Aliases: []string{"delete"},
	Short:   "Delete status page",
	Long: `Delete status page
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#delete-status-page`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteStatusPage(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to delete-status-page")

	},
}

func init() {

	Cmd.AddCommand(DeleteStatusPageCmd)
}
