package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchSecurityFindingsCmd = &cobra.Command{
	Use: "search-security-findings",

	Short: "Search security findings",
	Long: `Search security findings
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#search-security-findings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListSecurityFindingsResponse
		var err error

		var body datadogV2.SecurityFindingsSearchRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchSecurityFindings(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to search-security-findings")

		fmt.Println(cmdutil.FormatJSON(res, "finding"))
	},
}

func init() {

	SearchSecurityFindingsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	SearchSecurityFindingsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(SearchSecurityFindingsCmd)
}
