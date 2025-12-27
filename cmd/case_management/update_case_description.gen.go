package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateCaseDescriptionCmd = &cobra.Command{
	Use:   "update_case_description [case_id]",
	Short: "Update case description",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCaseDescription(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseUpdateDescriptionRequest{})
		if err != nil {
			log.Fatalf("failed to update_case_description: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management")
	},
}

func init() {
	Cmd.AddCommand(UpdateCaseDescriptionCmd)
}
