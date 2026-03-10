package oci_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTenancyConfigsCmd = &cobra.Command{
	Use: "get-tenancy-configs",

	Short: "Get tenancy configs",
	Long: `Get tenancy configs
Documentation: https://docs.datadoghq.com/api/latest/oci-integration/#get-tenancy-configs`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TenancyConfigList
		var err error

		api := datadogV2.NewOCIIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTenancyConfigs(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-tenancy-configs")

		fmt.Println(cmdutil.FormatJSON(res, "oci_tenancy"))
	},
}

func init() {

	Cmd.AddCommand(GetTenancyConfigsCmd)
}
