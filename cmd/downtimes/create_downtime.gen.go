package downtimes

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDowntimeCmd = &cobra.Command{
	Use:     "create-downtime",
	Aliases: []string{"create"},
	Short:   "Schedule a downtime",
	Long: `Schedule a downtime
Documentation: https://docs.datadoghq.com/api/latest/downtimes/#create-downtime`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DowntimeResponse
		var err error

		var body datadogV2.DowntimeCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDowntimesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateDowntime(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-downtime")

		fmt.Println(cmdutil.FormatJSON(res, "downtime"))
	},
}

func init() {

	CreateDowntimeCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDowntimeCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDowntimeCmd)
}
