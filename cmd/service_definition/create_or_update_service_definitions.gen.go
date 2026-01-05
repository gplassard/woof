package service_definition

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateOrUpdateServiceDefinitionsCmd = &cobra.Command{
	Use:     "create-or-update-service-definitions",
	Aliases: []string{"create-or-update"},
	Short:   "Create or update service definition",
	Long: `Create or update service definition
Documentation: https://docs.datadoghq.com/api/latest/service-definition/#create-or-update-service-definitions`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ServiceDefinitionCreateResponse
		var err error

		var body datadogV2.ServiceDefinitionsCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		res, _, err = api.CreateOrUpdateServiceDefinitions(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-or-update-service-definitions")

		cmd.Println(cmdutil.FormatJSON(res, "service_definition"))
	},
}

func init() {

	CreateOrUpdateServiceDefinitionsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateOrUpdateServiceDefinitionsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateOrUpdateServiceDefinitionsCmd)
}
