package seats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AssignSeatsUserCmd = &cobra.Command{
	Use:     "assign-seats-user",
	Aliases: []string{"assign-user"},
	Short:   "Assign seats to users",
	Long: `Assign seats to users
Documentation: https://docs.datadoghq.com/api/latest/seats/#assign-seats-user`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AssignSeatsUserResponse
		var err error

		var body datadogV2.AssignSeatsUserRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSeatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.AssignSeatsUser(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to assign-seats-user")

		fmt.Println(cmdutil.FormatJSON(res, "seat-assignments"))
	},
}

func init() {

	AssignSeatsUserCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AssignSeatsUserCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AssignSeatsUserCmd)
}
