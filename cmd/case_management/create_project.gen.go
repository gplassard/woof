package case_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateProjectCmd = &cobra.Command{
	Use:   "create-project",
	
	Short: "Create a project",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		res, _, err := api.CreateProject(client.NewContext(apiKey, appKey, site), datadogV2.ProjectCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-project: %v", err)
		}

		cmdutil.PrintJSON(res, "project")
	},
}

func init() {
	Cmd.AddCommand(CreateProjectCmd)
}
