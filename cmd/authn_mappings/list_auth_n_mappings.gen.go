package authn_mappings

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListAuthNMappingsCmd = &cobra.Command{
	Use: "list-auth-n-mappings",

	Short: "List all AuthN Mappings",
	Long: `List all AuthN Mappings
Documentation: https://docs.datadoghq.com/api/latest/authn-mappings/#list-auth-n-mappings`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.AuthNMappingsResponse
		var err error

		optionalParams := datadogV2.NewListAuthNMappingsOptionalParameters()

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("filter") {
			val, _ := cmd.Flags().GetString("filter")
			optionalParams.WithFilter(val)
		}

		if cmd.Flags().Changed("resource-type") {
			val, _ := cmd.Flags().GetString("resource-type")
			optionalParams.WithResourceType(val)
		}

		api := datadogV2.NewAuthNMappingsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListAuthNMappings(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-auth-n-mappings")

		cmd.Println(cmdutil.FormatJSON(res, "authn_mappings"))
	},
}

func init() {

	ListAuthNMappingsCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListAuthNMappingsCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListAuthNMappingsCmd.Flags().String("sort", "", "Sort AuthN Mappings depending on the given field.")

	ListAuthNMappingsCmd.Flags().String("filter", "", "Filter all mappings by the given string.")

	ListAuthNMappingsCmd.Flags().String("resource-type", "", "Filter by mapping resource type. Defaults to \"role\" if not specified.")

	Cmd.AddCommand(ListAuthNMappingsCmd)
}
