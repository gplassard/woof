package authn_mappings

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAuthNMappingCmd = &cobra.Command{
	Use:   "create-auth-n-mapping",
	
	Short: "Create an AuthN Mapping",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		res, _, err := api.CreateAuthNMapping(client.NewContext(apiKey, appKey, site), datadogV2.AuthNMappingCreateRequest{})
		cmdutil.HandleError(err, "failed to create-auth-n-mapping")

		cmdutil.PrintJSON(res, "authn_mappings")
	},
}

func init() {
	Cmd.AddCommand(CreateAuthNMappingCmd)
}
