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

var UpdateMaintenanceCmd = &cobra.Command{
	Use: "update-maintenance [page_id] [maintenance_id]",

	Short: "Update maintenance",
	Long: `Update maintenance
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#update-maintenance`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Maintenance
		var err error

		var body datadogV2.PatchMaintenanceRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateMaintenance(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), uuid.MustParse(args[1]), body)
		cmdutil.HandleError(err, "failed to update-maintenance")

		fmt.Println(cmdutil.FormatJSON(res, "maintenances"))
	},
}

func init() {

	UpdateMaintenanceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateMaintenanceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateMaintenanceCmd)
}
