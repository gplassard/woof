package processes

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListProcessesCmd = &cobra.Command{
	Use:     "list-processes",
	Aliases: []string{"list"},
	Short:   "Get all processes",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewProcessesApi(client.NewAPIClient())
		res, _, err := api.ListProcesses(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-processes")

		cmd.Println(cmdutil.FormatJSON(res, "process"))
	},
}

func init() {
	Cmd.AddCommand(ListProcessesCmd)
}
