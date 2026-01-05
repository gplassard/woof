package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:     "update-application-security-waf-exclusion-filter [exclusion_filter_id] [payload]",
	Aliases: []string{"update-waf-exclusion-filter"},
	Short:   "Update a WAF exclusion filter",
	Long: `Update a WAF exclusion filter
Documentation: https://docs.datadoghq.com/api/latest/application-security/#update-application-security-waf-exclusion-filter`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationSecurityWafExclusionFilterResponse
		var err error

		var body datadogV2.ApplicationSecurityWafExclusionFilterUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err = api.UpdateApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-application-security-waf-exclusion-filter")

		cmd.Println(cmdutil.FormatJSON(res, "exclusion_filter"))
	},
}

func init() {
	Cmd.AddCommand(UpdateApplicationSecurityWafExclusionFilterCmd)
}
