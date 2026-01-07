package apm

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetServiceListCmd = &cobra.Command{
	Use: "get-service-list",

	Short: "Get service list",
	Long: `Get service list
Documentation: https://docs.datadoghq.com/api/latest/apm/#get-service-list`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceList
		var err error

		api := datadogV2.NewAPMApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetServiceList(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-service-list")

		fmt.Println(cmdutil.FormatJSON(res, "services_list"))
	},
}

func init() {

	Cmd.AddCommand(GetServiceListCmd)
}
