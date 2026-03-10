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

var ListServiceNowBusinessServicesCmd = &cobra.Command{
	Use: "list-service-now-business-services [instance_id]",

	Short: "List ServiceNow business services",
	Long: `List ServiceNow business services
Documentation: https://docs.datadoghq.com/api/latest/servicenow-integration/#list-service-now-business-services`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceNowBusinessServicesResponse
		var err error

		api := datadogV2.NewServiceNowIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListServiceNowBusinessServices(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to list-service-now-business-services")

		fmt.Println(cmdutil.FormatJSON(res, "servicenow_integration"))
	},
}

func init() {

	Cmd.AddCommand(ListServiceNowBusinessServicesCmd)
}
