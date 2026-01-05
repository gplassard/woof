package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var ValidateQueryCmd = &cobra.Command{
	Use: "validate-query [payload]",

	Short: "Validate query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RulesValidateQueryResponse
		var err error

		var body datadogV2.RulesValidateQueryRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.ValidateQuery(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to validate-query")

		cmd.Println(cmdutil.FormatJSON(res, "validate_response"))
	},
}

func init() {
	Cmd.AddCommand(ValidateQueryCmd)
}
