package seats

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UnassignSeatsUserCmd = &cobra.Command{
	Use:     "unassign-seats-user",
	Aliases: []string{"unassign-user"},
	Short:   "Unassign seats from users",
	Long: `Unassign seats from users
Documentation: https://docs.datadoghq.com/api/latest/seats/#unassign-seats-user`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.UnassignSeatsUserRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSeatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.UnassignSeatsUser(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to unassign-seats-user")

	},
}

func init() {

	UnassignSeatsUserCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UnassignSeatsUserCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UnassignSeatsUserCmd)
}
