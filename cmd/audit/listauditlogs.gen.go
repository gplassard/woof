package audit

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListAuditLogsCmd = &cobra.Command{
	Use:   "listauditlogs",
	Short: "Get a list of Audit Logs events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewAuditApi(client.NewAPIClient())
		res, _, err := api.ListAuditLogs(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listauditlogs: %v", err)
		}

		cmdutil.PrintJSON(res, "audit")
	},
}

func init() {
	Cmd.AddCommand(ListAuditLogsCmd)
}
