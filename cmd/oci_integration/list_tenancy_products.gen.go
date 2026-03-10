package oci_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTenancyProductsCmd = &cobra.Command{
	Use: "list-tenancy-products [productKeys]",

	Short: "List tenancy products",
	Long: `List tenancy products
Documentation: https://docs.datadoghq.com/api/latest/oci-integration/#list-tenancy-products`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TenancyProductsList
		var err error

		api := datadogV2.NewOCIIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListTenancyProducts(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-tenancy-products")

		fmt.Println(cmdutil.FormatJSON(res, "oci_tenancy_product"))
	},
}

func init() {

	Cmd.AddCommand(ListTenancyProductsCmd)
}
