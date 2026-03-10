package incidents

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteGlobalIncidentHandleCmd = &cobra.Command{
	Use:     "delete-global-incident-handle",
	Aliases: []string{"delete-global-handle"},
	Short:   "Delete global incident handle",
	Long: `Delete global incident handle
Documentation: https://docs.datadoghq.com/api/latest/incidents/#delete-global-incident-handle`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteGlobalIncidentHandle(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to delete-global-incident-handle")

	},
}

func init() {

	Cmd.AddCommand(DeleteGlobalIncidentHandleCmd)
}
