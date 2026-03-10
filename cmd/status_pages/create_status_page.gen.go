package status_pages

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateStatusPageCmd = &cobra.Command{
	Use:     "create-status-page",
	Aliases: []string{"create"},
	Short:   "Create status page",
	Long: `Create status page
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#create-status-page`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.StatusPage
		var err error

		var body datadogV2.CreateStatusPageRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateStatusPage(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-status-page")

		fmt.Println(cmdutil.FormatJSON(res, "status_pages"))
	},
}

func init() {

	CreateStatusPageCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateStatusPageCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateStatusPageCmd)
}
