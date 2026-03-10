package seats

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetSeatsUsersCmd = &cobra.Command{
	Use:     "get-seats-users [product_code]",
	Aliases: []string{"get-users"},
	Short:   "Get users with seats",
	Long: `Get users with seats
Documentation: https://docs.datadoghq.com/api/latest/seats/#get-seats-users`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.SeatUserDataArray
		var err error

		api := datadogV2.NewSeatsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetSeatsUsers(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-seats-users")

		fmt.Println(cmdutil.FormatJSON(res, "seat-users"))
	},
}

func init() {

	Cmd.AddCommand(GetSeatsUsersCmd)
}
