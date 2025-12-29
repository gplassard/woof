package case_management_type

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCaseTypeCmd = &cobra.Command{
	Use:   "create-case-type",
	
	Short: "Create a case type",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCaseManagementTypeApi(client.NewAPIClient())
		res, _, err := api.CreateCaseType(client.NewContext(apiKey, appKey, site), datadogV2.CaseTypeCreateRequest{})
		cmdutil.HandleError(err, "failed to create-case-type")

		cmd.Println(cmdutil.FormatJSON(res, "case_type"))
	},
}

func init() {
	Cmd.AddCommand(CreateCaseTypeCmd)
}
