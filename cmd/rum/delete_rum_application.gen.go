package rum

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteRUMApplicationCmd = &cobra.Command{
	Use:     "delete-rum-application [id]",
	Aliases: []string{"delete-application"},
	Short:   "Delete a RUM application",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		_, err := api.DeleteRUMApplication(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-rum-application")

	},
}

func init() {
	Cmd.AddCommand(DeleteRUMApplicationCmd)
}
