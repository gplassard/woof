package service_definition

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateOrUpdateServiceDefinitionsCmd = &cobra.Command{
	Use:     "create-or-update-service-definitions [payload]",
	Aliases: []string{"create-or-update"},
	Short:   "Create or update service definition",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.ServiceDefinitionsCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		res, _, err := api.CreateOrUpdateServiceDefinitions(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-or-update-service-definitions")

		cmd.Println(cmdutil.FormatJSON(res, "service_definition"))
	},
}

func init() {
	Cmd.AddCommand(CreateOrUpdateServiceDefinitionsCmd)
}
