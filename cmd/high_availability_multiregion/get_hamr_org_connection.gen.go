package high_availability_multiregion

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetHamrOrgConnectionCmd = &cobra.Command{
	Use: "get-hamr-org-connection",

	Short: "Get HAMR organization connection",
	Long: `Get HAMR organization connection
Documentation: https://docs.datadoghq.com/api/latest/high-availability-multiregion/#get-hamr-org-connection`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.HamrOrgConnectionResponse
		var err error

		api := datadogV2.NewHighAvailabilityMultiRegionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetHamrOrgConnection(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-hamr-org-connection")

		fmt.Println(cmdutil.FormatJSON(res, "hamr_org_connections"))
	},
}

func init() {

	Cmd.AddCommand(GetHamrOrgConnectionCmd)
}
