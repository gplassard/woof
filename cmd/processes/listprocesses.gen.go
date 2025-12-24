package processes

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListProcessesCmd = &cobra.Command{
	Use:   "listprocesses",
	Short: "Get all processes",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewProcessesApi(client.NewAPIClient())
		res, _, err := api.ListProcesses(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listprocesses: %v", err)
		}

		cmdutil.PrintJSON(res, "processes")
	},
}

func init() {
	Cmd.AddCommand(ListProcessesCmd)
}
