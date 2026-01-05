package apm_retention_filters

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ReorderApmRetentionFiltersCmd = &cobra.Command{
	Use:     "reorder-apm-retention-filters [payload]",
	Aliases: []string{"reorder"},
	Short:   "Re-order retention filters",
	Long: `Re-order retention filters
Documentation: https://docs.datadoghq.com/api/latest/apm-retention-filters/#reorder-apm-retention-filters`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.ReorderRetentionFiltersRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		_, err = api.ReorderApmRetentionFilters(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to reorder-apm-retention-filters")

	},
}

func init() {
	Cmd.AddCommand(ReorderApmRetentionFiltersCmd)
}
