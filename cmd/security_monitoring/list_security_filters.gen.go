package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListSecurityFiltersCmd = &cobra.Command{
	Use: "list-security-filters",

	Short: "Get all security filters",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListSecurityFilters(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-security-filters")

		cmd.Println(cmdutil.FormatJSON(res, "security_filters"))
	},
}

func init() {
	Cmd.AddCommand(ListSecurityFiltersCmd)
}
