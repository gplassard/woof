package case_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UnassignCaseCmd = &cobra.Command{
	Use:   "unassigncase [case_id]",
	Short: "Unassign case",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UnassignCase(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseEmptyRequest{})
		if err != nil {
			log.Fatalf("failed to unassigncase: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management")
	},
}

func init() {
	Cmd.AddCommand(UnassignCaseCmd)
}
