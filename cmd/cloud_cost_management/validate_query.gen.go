package cloud_cost_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ValidateQueryCmd = &cobra.Command{
	Use:   "validate-query",
	
	Short: "Validate query",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ValidateQuery(client.NewContext(apiKey, appKey, site), datadogV2.RulesValidateQueryRequest{})
		cmdutil.HandleError(err, "failed to validate-query")

		cmd.Println(cmdutil.FormatJSON(res, "validate_response"))
	},
}

func init() {
	Cmd.AddCommand(ValidateQueryCmd)
}
