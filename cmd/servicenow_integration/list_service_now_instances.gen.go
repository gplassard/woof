package servicenow_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListServiceNowInstancesCmd = &cobra.Command{
	Use: "list-service-now-instances",

	Short: "List ServiceNow instances",
	Long: `List ServiceNow instances
Documentation: https://docs.datadoghq.com/api/latest/servicenow-integration/#list-service-now-instances`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceNowInstancesResponse
		var err error

		api := datadogV2.NewServiceNowIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListServiceNowInstances(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-service-now-instances")

		fmt.Println(cmdutil.FormatJSON(res, "instance"))
	},
}

func init() {

	Cmd.AddCommand(ListServiceNowInstancesCmd)
}
