package service_definition

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOrUpdateServiceDefinitionsCmd = &cobra.Command{
	Use:   "createorupdateservicedefinitions",
	Short: "Create or update service definition",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewServiceDefinitionApi(client.NewAPIClient())
		res, _, err := api.CreateOrUpdateServiceDefinitions(client.NewContext(apiKey, appKey, site), datadogV2.ServiceDefinitionsCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createorupdateservicedefinitions: %v", err)
		}

		cmdutil.PrintJSON(res, "service_definition")
	},
}

func init() {
	Cmd.AddCommand(CreateOrUpdateServiceDefinitionsCmd)
}
