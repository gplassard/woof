package servicenow_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateServiceNowTemplateCmd = &cobra.Command{
	Use: "create-service-now-template",

	Short: "Create ServiceNow template",
	Long: `Create ServiceNow template
Documentation: https://docs.datadoghq.com/api/latest/servicenow-integration/#create-service-now-template`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceNowTemplateResponse
		var err error

		var body datadogV2.ServiceNowTemplateCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceNowIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateServiceNowTemplate(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-service-now-template")

		fmt.Println(cmdutil.FormatJSON(res, "servicenow_templates"))
	},
}

func init() {

	CreateServiceNowTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateServiceNowTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateServiceNowTemplateCmd)
}
