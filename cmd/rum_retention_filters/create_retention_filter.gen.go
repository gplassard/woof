package rum_retention_filters

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateRetentionFilterCmd = &cobra.Command{
	Use: "create-retention-filter [app_id] [payload]",

	Short: "Create a RUM retention filter",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.RumRetentionFilterCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.CreateRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-retention-filter")

		cmd.Println(cmdutil.FormatJSON(res, "retention_filters"))
	},
}

func init() {
	Cmd.AddCommand(CreateRetentionFilterCmd)
}
