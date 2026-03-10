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

var GetMaintenanceCmd = &cobra.Command{
	Use: "get-maintenance [page_id] [maintenance_id]",

	Short: "Get maintenance",
	Long: `Get maintenance
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#get-maintenance`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Maintenance
		var err error

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetMaintenance(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), uuid.MustParse(args[1]))
		cmdutil.HandleError(err, "failed to get-maintenance")

		fmt.Println(cmdutil.FormatJSON(res, "maintenances"))
	},
}

func init() {

	Cmd.AddCommand(GetMaintenanceCmd)
}
