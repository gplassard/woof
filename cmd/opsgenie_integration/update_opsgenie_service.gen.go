package opsgenie_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateOpsgenieServiceCmd = &cobra.Command{
	Use: "update-opsgenie-service [integration_service_id]",

	Short: "Update a single service object",
	Long: `Update a single service object
Documentation: https://docs.datadoghq.com/api/latest/opsgenie-integration/#update-opsgenie-service`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.OpsgenieServiceResponse
		var err error

		var body datadogV2.OpsgenieServiceUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateOpsgenieService(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-opsgenie-service")

		fmt.Println(cmdutil.FormatJSON(res, "opsgenie_service"))
	},
}

func init() {

	UpdateOpsgenieServiceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateOpsgenieServiceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateOpsgenieServiceCmd)
}
