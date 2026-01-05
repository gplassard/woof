package powerpack

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListPowerpacksCmd = &cobra.Command{
	Use:     "list-powerpacks",
	Aliases: []string{"list"},
	Short:   "Get all powerpacks",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		res, _, err := api.ListPowerpacks(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-powerpacks")

		cmd.Println(cmdutil.FormatJSON(res, "powerpack"))
	},
}

func init() {
	Cmd.AddCommand(ListPowerpacksCmd)
}
