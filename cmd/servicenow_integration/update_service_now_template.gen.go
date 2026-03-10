package servicenow_integration

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var UpdateServiceNowTemplateCmd = &cobra.Command{
	Use: "update-service-now-template [template_id]",

	Short: "Update ServiceNow template",
	Long: `Update ServiceNow template
Documentation: https://docs.datadoghq.com/api/latest/servicenow-integration/#update-service-now-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceNowTemplateResponse
		var err error

		var body datadogV2.ServiceNowTemplateUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceNowIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateServiceNowTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to update-service-now-template")

		fmt.Println(cmdutil.FormatJSON(res, "servicenow_templates"))
	},
}

func init() {

	UpdateServiceNowTemplateCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateServiceNowTemplateCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateServiceNowTemplateCmd)
}
