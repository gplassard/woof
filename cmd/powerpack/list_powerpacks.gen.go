package powerpack

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListPowerpacksCmd = &cobra.Command{
	Use:     "list-powerpacks",
	Aliases: []string{"list"},
	Short:   "Get all powerpacks",
	Long: `Get all powerpacks
Documentation: https://docs.datadoghq.com/api/latest/powerpack/#list-powerpacks`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListPowerpacksResponse
		var err error

		optionalParams := datadogV2.NewListPowerpacksOptionalParameters()

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		api := datadogV2.NewPowerpackApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListPowerpacks(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-powerpacks")

		fmt.Println(cmdutil.FormatJSON(res, "powerpack"))
	},
}

func init() {

	ListPowerpacksCmd.Flags().Int64("page-limit", 0, "Maximum number of powerpacks in the response.")

	ListPowerpacksCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	Cmd.AddCommand(ListPowerpacksCmd)
}
