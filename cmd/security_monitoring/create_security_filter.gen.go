package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateSecurityFilterCmd = &cobra.Command{
	Use: "create-security-filter [payload]",

	Short: "Create a security filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SecurityFilterResponse
		var err error

		var body datadogV2.SecurityFilterCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.CreateSecurityFilter(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-security-filter")

		cmd.Println(cmdutil.FormatJSON(res, "security_filters"))
	},
}

func init() {
	Cmd.AddCommand(CreateSecurityFilterCmd)
}
