package status_pages

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListMaintenancesCmd = &cobra.Command{
	Use: "list-maintenances",

	Short: "List maintenances",
	Long: `List maintenances
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#list-maintenances`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MaintenanceArray
		var err error

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListMaintenances(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-maintenances")

		fmt.Println(cmdutil.FormatJSON(res, "maintenances"))
	},
}

func init() {

	Cmd.AddCommand(ListMaintenancesCmd)
}
