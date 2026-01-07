package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRolesCmd = &cobra.Command{
	Use:     "list-roles",
	Aliases: []string{"list"},
	Short:   "List roles",
	Long: `List roles
Documentation: https://docs.datadoghq.com/api/latest/roles/#list-roles`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RolesResponse
		var err error

		optionalParams := datadogV2.NewListRolesOptionalParameters()

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

		if cmd.Flags().Changed("filter-id") {
			val, _ := cmd.Flags().GetString("filter-id")
			optionalParams.WithFilterId(val)
		}

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListRoles(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-roles")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {

	ListRolesCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListRolesCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListRolesCmd.Flags().String("sort", "", "Sort roles depending on the given field. Sort order is **ascending** by default. Sort order is **descending** if the field is prefixed by a negative sign, for example: 'sort=-name'.")

	ListRolesCmd.Flags().String("filter", "", "Filter all roles by the given string.")

	ListRolesCmd.Flags().String("filter-id", "", "Filter all roles by the given list of role IDs.")

	Cmd.AddCommand(ListRolesCmd)
}
