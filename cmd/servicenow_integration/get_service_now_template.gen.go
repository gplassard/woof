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

var GetServiceNowTemplateCmd = &cobra.Command{
	Use: "get-service-now-template [template_id]",

	Short: "Get ServiceNow template",
	Long: `Get ServiceNow template
Documentation: https://docs.datadoghq.com/api/latest/servicenow-integration/#get-service-now-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceNowTemplateResponse
		var err error

		api := datadogV2.NewServiceNowIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetServiceNowTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to get-service-now-template")

		fmt.Println(cmdutil.FormatJSON(res, "servicenow_templates"))
	},
}

func init() {

	Cmd.AddCommand(GetServiceNowTemplateCmd)
}
