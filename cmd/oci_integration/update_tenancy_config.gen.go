package oci_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTenancyConfigCmd = &cobra.Command{
	Use: "update-tenancy-config [tenancy_ocid]",

	Short: "Update tenancy config",
	Long: `Update tenancy config
Documentation: https://docs.datadoghq.com/api/latest/oci-integration/#update-tenancy-config`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TenancyConfig
		var err error

		var body datadogV2.UpdateTenancyConfigRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOCIIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateTenancyConfig(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-tenancy-config")

		fmt.Println(cmdutil.FormatJSON(res, "oci_tenancy"))
	},
}

func init() {

	UpdateTenancyConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateTenancyConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateTenancyConfigCmd)
}
