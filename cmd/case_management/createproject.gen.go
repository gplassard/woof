package case_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateProjectCmd = &cobra.Command{
	Use:   "createproject",
	Short: "Create a project",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.CreateProject(client.NewContext(apiKey, appKey, site), datadogV2.ProjectCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createproject: %v", err)
		}

		cmdutil.PrintJSON(res, "case_management")
	},
}

func init() {
	Cmd.AddCommand(CreateProjectCmd)
}
