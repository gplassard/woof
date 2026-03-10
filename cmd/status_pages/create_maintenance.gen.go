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

var CreateMaintenanceCmd = &cobra.Command{
	Use: "create-maintenance [page_id]",

	Short: "Schedule maintenance",
	Long: `Schedule maintenance
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#create-maintenance`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Maintenance
		var err error

		var body datadogV2.CreateMaintenanceRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateMaintenance(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to create-maintenance")

		fmt.Println(cmdutil.FormatJSON(res, "maintenances"))
	},
}

func init() {

	CreateMaintenanceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateMaintenanceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateMaintenanceCmd)
}
