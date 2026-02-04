package auth_n_mappings

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAuthNMappingCmd = &cobra.Command{
	Use:     "get-auth-n-mapping [authn_mapping_id]",
	Aliases: []string{"get"},
	Short:   "Get an AuthN Mapping by UUID",
	Long: `Get an AuthN Mapping by UUID
Documentation: https://docs.datadoghq.com/api/latest/auth-n-mappings/#get-auth-n-mapping`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuthNMappingResponse
		var err error

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAuthNMapping(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-auth-n-mapping")

		fmt.Println(cmdutil.FormatJSON(res, "auth_n_mapping"))
	},
}

func init() {

	Cmd.AddCommand(GetAuthNMappingCmd)
}
