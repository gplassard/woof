package reference_tables

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"strings"

	"github.com/spf13/cobra"
)

var GetRowsByIDCmd = &cobra.Command{
	Use: "get-rows-by-id [id] [row_id]",

	Short: "Get rows by id",
	Long: `Get rows by id
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#get-rows-by-id`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.TableRowResourceArray
		var err error

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetRowsByID(client.NewContext(apiKey, appKey, site), args[0], strings.Split(args[1], ", "))
		cmdutil.HandleError(err, "failed to get-rows-by-id")

		fmt.Println(cmdutil.FormatJSON(res, "rows_by_i_d"))
	},
}

func init() {

	Cmd.AddCommand(GetRowsByIDCmd)
}
