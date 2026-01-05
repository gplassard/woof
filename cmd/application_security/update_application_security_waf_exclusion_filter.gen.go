package application_security

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateApplicationSecurityWafExclusionFilterCmd = &cobra.Command{
	Use:     "update-application-security-waf-exclusion-filter [exclusion_filter_id]",
	Aliases: []string{"update-waf-exclusion-filter"},
	Short:   "Update a WAF exclusion filter",
	Long: `Update a WAF exclusion filter
Documentation: https://docs.datadoghq.com/api/latest/application-security/#update-application-security-waf-exclusion-filter`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ApplicationSecurityWafExclusionFilterResponse
		var err error

		var body datadogV2.ApplicationSecurityWafExclusionFilterUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewApplicationSecurityApi(client.NewAPIClient())
		res, _, err = api.UpdateApplicationSecurityWafExclusionFilter(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-application-security-waf-exclusion-filter")

		cmd.Println(cmdutil.FormatJSON(res, "exclusion_filter"))
	},
}

func init() {

	UpdateApplicationSecurityWafExclusionFilterCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateApplicationSecurityWafExclusionFilterCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateApplicationSecurityWafExclusionFilterCmd)
}
