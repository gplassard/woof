package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ArchiveCaseCmd = &cobra.Command{
	Use:   "archive-case [case_id]",
	
	Short: "Archive case",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.ArchiveCase(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseEmptyRequest{})
		if err != nil {
			log.Fatalf("failed to archive-case: %v", err)
		}

		cmdutil.PrintJSON(res, "case")
	},
}

func init() {
	Cmd.AddCommand(ArchiveCaseCmd)
}
