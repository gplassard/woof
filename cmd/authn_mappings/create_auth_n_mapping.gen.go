package authn_mappings

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateAuthNMappingCmd = &cobra.Command{
	Use: "create-auth-n-mapping",

	Short: "Create an AuthN Mapping",
	Long: `Create an AuthN Mapping
Documentation: https://docs.datadoghq.com/api/latest/authn-mappings/#create-auth-n-mapping`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuthNMappingResponse
		var err error

		var body datadogV2.AuthNMappingCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateAuthNMapping(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-auth-n-mapping")

		cmd.Println(cmdutil.FormatJSON(res, "authn_mappings"))
	},
}

func init() {

	CreateAuthNMappingCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateAuthNMappingCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateAuthNMappingCmd)
}
