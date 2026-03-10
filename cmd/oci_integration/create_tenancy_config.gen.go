package oci_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTenancyConfigCmd = &cobra.Command{
	Use: "create-tenancy-config",

	Short: "Create tenancy config",
	Long: `Create tenancy config
Documentation: https://docs.datadoghq.com/api/latest/oci-integration/#create-tenancy-config`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TenancyConfig
		var err error

		var body datadogV2.CreateTenancyConfigRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOCIIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateTenancyConfig(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-tenancy-config")

		fmt.Println(cmdutil.FormatJSON(res, "oci_tenancy"))
	},
}

func init() {

	CreateTenancyConfigCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateTenancyConfigCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateTenancyConfigCmd)
}
