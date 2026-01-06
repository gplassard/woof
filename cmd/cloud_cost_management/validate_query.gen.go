package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ValidateQueryCmd = &cobra.Command{
	Use: "validate-query",

	Short: "Validate query",
	Long: `Validate query
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#validate-query`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RulesValidateQueryResponse
		var err error

		var body datadogV2.RulesValidateQueryRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ValidateQuery(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-query")

		cmd.Println(cmdutil.FormatJSON(res, "validate_response"))
	},
}

func init() {

	ValidateQueryCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ValidateQueryCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ValidateQueryCmd)
}
