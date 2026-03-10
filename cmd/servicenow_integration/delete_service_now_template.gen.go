package servicenow_integration

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteServiceNowTemplateCmd = &cobra.Command{
	Use: "delete-service-now-template [template_id]",

	Short: "Delete ServiceNow template",
	Long: `Delete ServiceNow template
Documentation: https://docs.datadoghq.com/api/latest/servicenow-integration/#delete-service-now-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewServiceNowIntegrationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteServiceNowTemplate(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to delete-service-now-template")

	},
}

func init() {

	Cmd.AddCommand(DeleteServiceNowTemplateCmd)
}
