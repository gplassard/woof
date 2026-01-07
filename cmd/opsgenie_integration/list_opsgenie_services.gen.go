package opsgenie_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListOpsgenieServicesCmd = &cobra.Command{
	Use: "list-opsgenie-services",

	Short: "Get all service objects",
	Long: `Get all service objects
Documentation: https://docs.datadoghq.com/api/latest/opsgenie-integration/#list-opsgenie-services`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OpsgenieServicesResponse
		var err error

		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListOpsgenieServices(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-opsgenie-services")

		fmt.Println(cmdutil.FormatJSON(res, "opsgenie-service"))
	},
}

func init() {

	Cmd.AddCommand(ListOpsgenieServicesCmd)
}
