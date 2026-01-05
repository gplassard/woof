package rum

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateRUMApplicationCmd = &cobra.Command{
	Use:     "create-rum-application [payload]",
	Aliases: []string{"create-application"},
	Short:   "Create a new RUM application",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.RUMApplicationCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.CreateRUMApplication(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-rum-application")

		cmd.Println(cmdutil.FormatJSON(res, "rum_application"))
	},
}

func init() {
	Cmd.AddCommand(CreateRUMApplicationCmd)
}
