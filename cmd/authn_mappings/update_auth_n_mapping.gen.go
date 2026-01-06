package authn_mappings

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateAuthNMappingCmd = &cobra.Command{
	Use: "update-auth-n-mapping [authn_mapping_id]",

	Short: "Edit an AuthN Mapping",
	Long: `Edit an AuthN Mapping
Documentation: https://docs.datadoghq.com/api/latest/authn-mappings/#update-auth-n-mapping`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuthNMappingResponse
		var err error

		var body datadogV2.AuthNMappingUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateAuthNMapping(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-auth-n-mapping")

		cmd.Println(cmdutil.FormatJSON(res, "authn_mappings"))
	},
}

func init() {

	UpdateAuthNMappingCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateAuthNMappingCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateAuthNMappingCmd)
}
