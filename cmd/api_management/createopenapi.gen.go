package api_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOpenAPICmd = &cobra.Command{
	Use:   "createopenapi",
	Short: "Create a new API",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		res, _, err := api.CreateOpenAPI(client.NewContext(apiKey, appKey, site), *datadogV2.NewCreateOpenAPIOptionalParameters())
		if err != nil {
			log.Fatalf("failed to createopenapi: %v", err)
		}

		cmdutil.PrintJSON(res, "api_management")
	},
}

func init() {
	Cmd.AddCommand(CreateOpenAPICmd)
}
