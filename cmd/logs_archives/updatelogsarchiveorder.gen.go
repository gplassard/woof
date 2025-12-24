package logs_archives

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateLogsArchiveOrderCmd = &cobra.Command{
	Use:   "updatelogsarchiveorder",
	Short: "Update archive order",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.UpdateLogsArchiveOrder(client.NewContext(apiKey, appKey, site), datadogV2.LogsArchiveOrder{})
		if err != nil {
			log.Fatalf("failed to updatelogsarchiveorder: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_archives")
	},
}

func init() {
	Cmd.AddCommand(UpdateLogsArchiveOrderCmd)
}
