package incidents

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateGlobalIncidentSettingsCmd = &cobra.Command{
	Use:     "update-global-incident-settings",
	Aliases: []string{"update-global-settings"},
	Short:   "Update global incident settings",
	Long: `Update global incident settings
Documentation: https://docs.datadoghq.com/api/latest/incidents/#update-global-incident-settings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GlobalIncidentSettingsResponse
		var err error

		var body datadogV2.GlobalIncidentSettingsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewIncidentsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateGlobalIncidentSettings(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to update-global-incident-settings")

		fmt.Println(cmdutil.FormatJSON(res, "incidents_global_settings"))
	},
}

func init() {

	UpdateGlobalIncidentSettingsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateGlobalIncidentSettingsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateGlobalIncidentSettingsCmd)
}
