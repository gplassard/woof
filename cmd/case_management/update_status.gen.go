package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateStatusCmd = &cobra.Command{
	Use:   "update_status [case_id]",
	Short: "Update case status",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateStatus(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseUpdateStatusRequest{})
		if err != nil {
			log.Fatalf("failed to update_status: %v", err)
		}

		cmdutil.PrintJSON(res, "case")
	},
}

func init() {
	Cmd.AddCommand(UpdateStatusCmd)
}
