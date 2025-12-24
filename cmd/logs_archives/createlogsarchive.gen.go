package logs_archives

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateLogsArchiveCmd = &cobra.Command{
	Use:   "createlogsarchive",
	Short: "Create an archive",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsArchivesApi(client.NewAPIClient())
		res, _, err := api.CreateLogsArchive(client.NewContext(apiKey, appKey, site), datadogV2.LogsArchiveCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createlogsarchive: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_archives")
	},
}

func init() {
	Cmd.AddCommand(CreateLogsArchiveCmd)
}
