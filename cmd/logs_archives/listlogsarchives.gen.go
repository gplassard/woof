package logs_archives

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListLogsArchivesCmd = &cobra.Command{
	Use:   "listlogsarchives",
	Short: "Get all archives",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.ListLogsArchives(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listlogsarchives: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_archives")
	},
}

func init() {
	Cmd.AddCommand(ListLogsArchivesCmd)
}
