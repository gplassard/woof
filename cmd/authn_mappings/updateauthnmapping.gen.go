package authn_mappings

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateAuthNMappingCmd = &cobra.Command{
	Use:   "updateauthnmapping [authn_mapping_id]",
	Short: "Edit an AuthN Mapping",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		res, _, err := api.UpdateAuthNMapping(client.NewContext(apiKey, appKey, site), args[0], datadogV2.AuthNMappingUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updateauthnmapping: %v", err)
		}

		cmdutil.PrintJSON(res, "authn_mappings")
	},
}

func init() {
	Cmd.AddCommand(UpdateAuthNMappingCmd)
}
