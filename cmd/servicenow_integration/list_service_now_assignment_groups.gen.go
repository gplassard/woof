package servicenow_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var ListServiceNowAssignmentGroupsCmd = &cobra.Command{
	Use: "list-service-now-assignment-groups [instance_id]",

	Short: "List ServiceNow assignment groups",
	Long: `List ServiceNow assignment groups
Documentation: https://docs.datadoghq.com/api/latest/servicenow-integration/#list-service-now-assignment-groups`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceNowAssignmentGroupsResponse
		var err error

		api := datadogV2.NewServiceNowIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListServiceNowAssignmentGroups(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to list-service-now-assignment-groups")

		fmt.Println(cmdutil.FormatJSON(res, "assignment_groups"))
	},
}

func init() {

	Cmd.AddCommand(ListServiceNowAssignmentGroupsCmd)
}
