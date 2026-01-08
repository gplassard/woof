package auth_n_mappings

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteAuthNMappingCmd = &cobra.Command{
	Use:     "delete-auth-n-mapping [authn_mapping_id]",
	Aliases: []string{"delete"},
	Short:   "Delete an AuthN Mapping",
	Long: `Delete an AuthN Mapping
Documentation: https://docs.datadoghq.com/api/latest/auth-n-mappings/#delete-auth-n-mapping`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteAuthNMapping(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-auth-n-mapping")

	},
}

func init() {

	Cmd.AddCommand(DeleteAuthNMappingCmd)
}
