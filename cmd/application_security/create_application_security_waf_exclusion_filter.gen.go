package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:     "create-application-security-waf-exclusion-filter [payload]",
	Aliases: []string{"create-waf-exclusion-filter"},
	Short:   "Create a WAF exclusion filter",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.ApplicationSecurityWafExclusionFilterCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err := api.CreateApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-application-security-waf-exclusion-filter")

		cmd.Println(cmdutil.FormatJSON(res, "exclusion_filter"))
	},
}

func init() {
	Cmd.AddCommand(CreateApplicationSecurityWafExclusionFilterCmd)
}
