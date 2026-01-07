package usage_metering

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetProjectedCostCmd = &cobra.Command{
	Use: "get-projected-cost",

	Short: "Get projected cost across your account",
	Long: `Get projected cost across your account
Documentation: https://docs.datadoghq.com/api/latest/usage-metering/#get-projected-cost`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ProjectedCostResponse
		var err error

		optionalParams := datadogV2.NewGetProjectedCostOptionalParameters()

		if cmd.Flags().Changed("view") {
			val, _ := cmd.Flags().GetString("view")
			optionalParams.WithView(val)
		}

		if cmd.Flags().Changed("include-connected-accounts") {
			val, _ := cmd.Flags().GetString("include-connected-accounts")
			optionalParams.WithIncludeConnectedAccounts(val)
		}

		api := datadogV2.NewUsageMeteringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetProjectedCost(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to get-projected-cost")

		cmd.Println(cmdutil.FormatJSON(res, "projected_cost"))
	},
}

func init() {

	GetProjectedCostCmd.Flags().String("view", "", "String to specify whether cost is broken down at a parent-org level or at the sub-org level. Available views are 'summary' and 'sub-org'. Defaults to 'summary'.")

	GetProjectedCostCmd.Flags().String("include-connected-accounts", "", "Boolean to specify whether to include accounts connected to the current account as partner customers in the Datadog partner network program. Defaults to 'false'. ")

	Cmd.AddCommand(GetProjectedCostCmd)
}
