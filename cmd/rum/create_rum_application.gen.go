package rum

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateRUMApplicationCmd = &cobra.Command{
	Use:     "create-rum-application",
	Aliases: []string{"create-application"},
	Short:   "Create a new RUM application",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.CreateRUMApplication(client.NewContext(apiKey, appKey, site), datadogV2.RUMApplicationCreateRequest{})
		cmdutil.HandleError(err, "failed to create-rum-application")

		cmd.Println(cmdutil.FormatJSON(res, "rum_application"))
	},
}

func init() {
	Cmd.AddCommand(CreateRUMApplicationCmd)
}
