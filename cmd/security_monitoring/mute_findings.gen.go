package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var MuteFindingsCmd = &cobra.Command{
	Use: "mute-findings",

	Short: "Mute or unmute a batch of findings",
	Long: `Mute or unmute a batch of findings
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#mute-findings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.BulkMuteFindingsResponse
		var err error

		var body datadogV2.BulkMuteFindingsRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.MuteFindings(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to mute-findings")

		cmd.Println(cmdutil.FormatJSON(res, "finding"))
	},
}

func init() {

	MuteFindingsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	MuteFindingsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(MuteFindingsCmd)
}
