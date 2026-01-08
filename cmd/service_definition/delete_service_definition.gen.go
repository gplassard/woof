package service_definition

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteServiceDefinitionCmd = &cobra.Command{
	Use:     "delete-service-definition [service_name]",
	Aliases: []string{"delete"},
	Short:   "Delete a single service definition",
	Long: `Delete a single service definition
Documentation: https://docs.datadoghq.com/api/latest/service-definition/#delete-service-definition`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteServiceDefinition(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-service-definition")

	},
}

func init() {

	Cmd.AddCommand(DeleteServiceDefinitionCmd)
}
