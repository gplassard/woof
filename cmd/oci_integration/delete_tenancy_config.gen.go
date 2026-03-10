package oci_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTenancyConfigCmd = &cobra.Command{
	Use: "delete-tenancy-config [tenancy_ocid]",

	Short: "Delete tenancy config",
	Long: `Delete tenancy config
Documentation: https://docs.datadoghq.com/api/latest/oci-integration/#delete-tenancy-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewOCIIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteTenancyConfig(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-tenancy-config")

	},
}

func init() {

	Cmd.AddCommand(DeleteTenancyConfigCmd)
}
