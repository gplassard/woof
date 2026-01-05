package powerpack

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdatePowerpackCmd = &cobra.Command{
	Use:     "update-powerpack [powerpack_id]",
	Aliases: []string{"update"},
	Short:   "Update a powerpack",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		res, _, err := api.UpdatePowerpack(client.NewContext(apiKey, appKey, site), args[0], datadogV2.Powerpack{})
		cmdutil.HandleError(err, "failed to update-powerpack")

		cmd.Println(cmdutil.FormatJSON(res, "powerpack"))
	},
}

func init() {
	Cmd.AddCommand(UpdatePowerpackCmd)
}
