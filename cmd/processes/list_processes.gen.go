package processes

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListProcessesCmd = &cobra.Command{
	Use:   "list-processes",
	Short: "Get all processes",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewProcessesApi(client.NewAPIClient())
		res, _, err := api.ListProcesses(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-processes: %v", err)
		}

		cmdutil.PrintJSON(res, "process")
	},
}

func init() {
	Cmd.AddCommand(ListProcessesCmd)
}
