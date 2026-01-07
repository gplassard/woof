package software_catalog

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCatalogRelationCmd = &cobra.Command{
	Use: "list-catalog-relation",

	Short: "Get a list of entity relations",
	Long: `Get a list of entity relations
Documentation: https://docs.datadoghq.com/api/latest/software-catalog/#list-catalog-relation`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListRelationCatalogResponse
		var err error

		optionalParams := datadogV2.NewListCatalogRelationOptionalParameters()

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("filter-type") {
			val, _ := cmd.Flags().GetString("filter-type")
			optionalParams.WithFilterType(val)
		}

		if cmd.Flags().Changed("filter-from-ref") {
			val, _ := cmd.Flags().GetString("filter-from-ref")
			optionalParams.WithFilterFromRef(val)
		}

		if cmd.Flags().Changed("filter-to-ref") {
			val, _ := cmd.Flags().GetString("filter-to-ref")
			optionalParams.WithFilterToRef(val)
		}

		if cmd.Flags().Changed("include") {
			val, _ := cmd.Flags().GetString("include")
			optionalParams.WithInclude(val)
		}

		if cmd.Flags().Changed("includediscovered") {
			val, _ := cmd.Flags().GetString("includediscovered")
			optionalParams.WithIncludeDiscovered(val)
		}

		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCatalogRelation(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-catalog-relation")

		cmd.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {

	ListCatalogRelationCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	ListCatalogRelationCmd.Flags().Int64("page-limit", 0, "Maximum number of relations in the response.")

	ListCatalogRelationCmd.Flags().String("filter-type", "", "Filter relations by type.")

	ListCatalogRelationCmd.Flags().String("filter-from-ref", "", "Filter relations by the reference of the first entity in the relation.")

	ListCatalogRelationCmd.Flags().String("filter-to-ref", "", "Filter relations by the reference of the second entity in the relation.")

	ListCatalogRelationCmd.Flags().String("include", "", "Include relationship data.")

	ListCatalogRelationCmd.Flags().String("includediscovered", "", "If true, includes relationships discovered by APM and USM.")

	Cmd.AddCommand(ListCatalogRelationCmd)
}
