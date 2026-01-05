package authn_mappings

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteAuthNMappingCmd = &cobra.Command{
	Use: "delete-auth-n-mapping [authn_mapping_id]",

	Short: "Delete an AuthN Mapping",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		_, err := api.DeleteAuthNMapping(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-auth-n-mapping")

	},
}

func init() {
	Cmd.AddCommand(DeleteAuthNMappingCmd)
}
