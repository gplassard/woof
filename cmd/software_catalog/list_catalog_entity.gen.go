package software_catalog

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCatalogEntityCmd = &cobra.Command{
	Use: "list-catalog-entity",

	Short: "Get a list of entities",
	Long: `Get a list of entities
Documentation: https://docs.datadoghq.com/api/latest/software-catalog/#list-catalog-entity`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListEntityCatalogResponse
		var err error

		optionalParams := datadogV2.NewListCatalogEntityOptionalParameters()

		if cmd.Flags().Changed("page-offset") {
			val, _ := cmd.Flags().GetInt64("page-offset")
			optionalParams.WithPageOffset(val)
		}

		if cmd.Flags().Changed("page-limit") {
			val, _ := cmd.Flags().GetInt64("page-limit")
			optionalParams.WithPageLimit(val)
		}

		if cmd.Flags().Changed("filter-id") {
			val, _ := cmd.Flags().GetString("filter-id")
			optionalParams.WithFilterId(val)
		}

		if cmd.Flags().Changed("filter-ref") {
			val, _ := cmd.Flags().GetString("filter-ref")
			optionalParams.WithFilterRef(val)
		}

		if cmd.Flags().Changed("filter-name") {
			val, _ := cmd.Flags().GetString("filter-name")
			optionalParams.WithFilterName(val)
		}

		if cmd.Flags().Changed("filter-kind") {
			val, _ := cmd.Flags().GetString("filter-kind")
			optionalParams.WithFilterKind(val)
		}

		if cmd.Flags().Changed("filter-owner") {
			val, _ := cmd.Flags().GetString("filter-owner")
			optionalParams.WithFilterOwner(val)
		}

		if cmd.Flags().Changed("filter-relation-type") {
			val, _ := cmd.Flags().GetString("filter-relation-type")
			optionalParams.WithFilterRelationType(val)
		}

		if cmd.Flags().Changed("filter-exclude-snapshot") {
			val, _ := cmd.Flags().GetString("filter-exclude-snapshot")
			optionalParams.WithFilterExcludeSnapshot(val)
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
		res, _, err = api.ListCatalogEntity(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-catalog-entity")

		fmt.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {

	ListCatalogEntityCmd.Flags().Int64("page-offset", 0, "Specific offset to use as the beginning of the returned page.")

	ListCatalogEntityCmd.Flags().Int64("page-limit", 0, "Maximum number of entities in the response.")

	ListCatalogEntityCmd.Flags().String("filter-id", "", "Filter entities by UUID.")

	ListCatalogEntityCmd.Flags().String("filter-ref", "", "Filter entities by reference")

	ListCatalogEntityCmd.Flags().String("filter-name", "", "Filter entities by name.")

	ListCatalogEntityCmd.Flags().String("filter-kind", "", "Filter entities by kind.")

	ListCatalogEntityCmd.Flags().String("filter-owner", "", "Filter entities by owner.")

	ListCatalogEntityCmd.Flags().String("filter-relation-type", "", "Filter entities by relation type.")

	ListCatalogEntityCmd.Flags().String("filter-exclude-snapshot", "", "Filter entities by excluding snapshotted entities.")

	ListCatalogEntityCmd.Flags().String("include", "", "Include relationship data.")

	ListCatalogEntityCmd.Flags().String("includediscovered", "", "If true, includes discovered services from APM and USM that do not have entity definitions.")

	Cmd.AddCommand(ListCatalogEntityCmd)
}
