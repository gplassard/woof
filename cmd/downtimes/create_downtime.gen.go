package downtimes

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateDowntimeCmd = &cobra.Command{
	Use:     "create-downtime [payload]",
	Aliases: []string{"create"},
	Short:   "Schedule a downtime",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.DowntimeCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		res, _, err := api.CreateDowntime(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-downtime")

		cmd.Println(cmdutil.FormatJSON(res, "downtime"))
	},
}

func init() {
	Cmd.AddCommand(CreateDowntimeCmd)
}
