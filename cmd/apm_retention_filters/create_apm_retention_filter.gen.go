package apm_retention_filters

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateApmRetentionFilterCmd = &cobra.Command{
	Use:     "create-apm-retention-filter",
	Aliases: []string{"create"},
	Short:   "Create a retention filter",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.CreateApmRetentionFilter(client.NewContext(apiKey, appKey, site), datadogV2.RetentionFilterCreateRequest{})
		cmdutil.HandleError(err, "failed to create-apm-retention-filter")

		cmd.Println(cmdutil.FormatJSON(res, "apm_retention_filter"))
	},
}

func init() {
	Cmd.AddCommand(CreateApmRetentionFilterCmd)
}
