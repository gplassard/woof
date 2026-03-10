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

var UpdateStatusPageCmd = &cobra.Command{
	Use:     "update-status-page [page_id]",
	Aliases: []string{"update"},
	Short:   "Update status page",
	Long: `Update status page
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#update-status-page`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.StatusPage
		var err error

		var body datadogV2.PatchStatusPageRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateStatusPage(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-status-page")

		fmt.Println(cmdutil.FormatJSON(res, "status_pages"))
	},
}

func init() {

	UpdateStatusPageCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateStatusPageCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateStatusPageCmd)
}
