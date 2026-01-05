package reference_tables

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
	"strings"
)

var GetRowsByIDCmd = &cobra.Command{
	Use: "get-rows-by-id [id] [row_id]",

	Short: "Get rows by id",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err := api.GetRowsByID(client.NewContext(apiKey, appKey, site), args[0], strings.Split(args[1], ","))
		cmdutil.HandleError(err, "failed to get-rows-by-id")

		cmd.Println(cmdutil.FormatJSON(res, "row"))
	},
}

func init() {
	Cmd.AddCommand(GetRowsByIDCmd)
}
