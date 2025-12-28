package cloud_cost_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ValidateQueryCmd = &cobra.Command{
	Use:   "validate_query",
	Short: "Validate query",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ValidateQuery(client.NewContext(apiKey, appKey, site), datadogV2.RulesValidateQueryRequest{})
		if err != nil {
			log.Fatalf("failed to validate_query: %v", err)
		}

		cmdutil.PrintJSON(res, "validate_response")
	},
}

func init() {
	Cmd.AddCommand(ValidateQueryCmd)
}
