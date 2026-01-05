package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var DetachCaseCmd = &cobra.Command{
	Use: "detach-case [payload]",

	Short: "Detach security findings from their case",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.DetachCaseRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err = api.DetachCase(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to detach-case")

	},
}

func init() {
	Cmd.AddCommand(DetachCaseCmd)
}
