package authn_mappings

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAuthNMappingCmd = &cobra.Command{
	Use: "get-auth-n-mapping [authn_mapping_id]",

	Short: "Get an AuthN Mapping by UUID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuthNMappingResponse
		var err error

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		res, _, err = api.GetAuthNMapping(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-auth-n-mapping")

		cmd.Println(cmdutil.FormatJSON(res, "authn_mappings"))
	},
}

func init() {
	Cmd.AddCommand(GetAuthNMappingCmd)
}
