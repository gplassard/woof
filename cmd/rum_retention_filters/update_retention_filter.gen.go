package rum_retention_filters

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateRetentionFilterCmd = &cobra.Command{
	Use: "update-retention-filter [app_id] [rf_id] [payload]",

	Short: "Update a RUM retention filter",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RumRetentionFilterResponse
		var err error

		var body datadogV2.RumRetentionFilterUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err = api.UpdateRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-retention-filter")

		cmd.Println(cmdutil.FormatJSON(res, "retention_filters"))
	},
}

func init() {
	Cmd.AddCommand(UpdateRetentionFilterCmd)
}
