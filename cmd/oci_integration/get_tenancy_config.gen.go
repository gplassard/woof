package oci_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetTenancyConfigCmd = &cobra.Command{
	Use: "get-tenancy-config [tenancy_ocid]",

	Short: "Get tenancy config",
	Long: `Get tenancy config
Documentation: https://docs.datadoghq.com/api/latest/oci-integration/#get-tenancy-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TenancyConfig
		var err error

		api := datadogV2.NewOCIIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetTenancyConfig(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-tenancy-config")

		fmt.Println(cmdutil.FormatJSON(res, "oci_tenancy"))
	},
}

func init() {

	Cmd.AddCommand(GetTenancyConfigCmd)
}
