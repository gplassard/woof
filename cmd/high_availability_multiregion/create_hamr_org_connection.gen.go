package high_availability_multiregion

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateHamrOrgConnectionCmd = &cobra.Command{
	Use: "create-hamr-org-connection",

	Short: "Create or update HAMR organization connection",
	Long: `Create or update HAMR organization connection
Documentation: https://docs.datadoghq.com/api/latest/high-availability-multiregion/#create-hamr-org-connection`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.HamrOrgConnectionResponse
		var err error

		var body datadogV2.HamrOrgConnectionRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewHighAvailabilityMultiRegionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateHamrOrgConnection(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-hamr-org-connection")

		fmt.Println(cmdutil.FormatJSON(res, "hamr_org_connections"))
	},
}

func init() {

	CreateHamrOrgConnectionCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateHamrOrgConnectionCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateHamrOrgConnectionCmd)
}
