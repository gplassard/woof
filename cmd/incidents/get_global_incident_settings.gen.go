package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetGlobalIncidentSettingsCmd = &cobra.Command{
	Use:     "get-global-incident-settings",
	Aliases: []string{"get-global-settings"},
	Short:   "Get global incident settings",
	Long: `Get global incident settings
Documentation: https://docs.datadoghq.com/api/latest/incidents/#get-global-incident-settings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GlobalIncidentSettingsResponse
		var err error

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetGlobalIncidentSettings(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-global-incident-settings")

		fmt.Println(cmdutil.FormatJSON(res, "incidents_global_settings"))
	},
}

func init() {

	Cmd.AddCommand(GetGlobalIncidentSettingsCmd)
}
