package rum

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateRUMApplicationCmd = &cobra.Command{
	Use:     "update-rum-application [id]",
	Aliases: []string{"update-application"},
	Short:   "Update a RUM application",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.UpdateRUMApplication(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RUMApplicationUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-rum-application")

		cmd.Println(cmdutil.FormatJSON(res, "rum_application"))
	},
}

func init() {
	Cmd.AddCommand(UpdateRUMApplicationCmd)
}
