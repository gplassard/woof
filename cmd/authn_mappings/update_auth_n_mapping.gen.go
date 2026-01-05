package authn_mappings

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateAuthNMappingCmd = &cobra.Command{
	Use: "update-auth-n-mapping [authn_mapping_id] [payload]",

	Short: "Edit an AuthN Mapping",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuthNMappingResponse
		var err error

		var body datadogV2.AuthNMappingUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		res, _, err = api.UpdateAuthNMapping(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-auth-n-mapping")

		cmd.Println(cmdutil.FormatJSON(res, "authn_mappings"))
	},
}

func init() {
	Cmd.AddCommand(UpdateAuthNMappingCmd)
}
