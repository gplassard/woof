package case_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateCaseDescriptionCmd = &cobra.Command{
	Use:   "updatecasedescription [case_id]",
	Short: "Update case description",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateCaseDescription(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CaseUpdateDescriptionRequest{})
		if err != nil {
			log.Fatalf("failed to updatecasedescription: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management")
	},
}

func init() {
	Cmd.AddCommand(UpdateCaseDescriptionCmd)
}
