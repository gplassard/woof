package users

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListUsersCmd = &cobra.Command{
	Use:     "list-users",
	Aliases: []string{"list"},
	Short:   "List all users",
	Long: `List all users
Documentation: https://docs.datadoghq.com/api/latest/users/#list-users`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UsersResponse
		var err error

		optionalParams := datadogV2.NewListUsersOptionalParameters()

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

		if cmd.Flags().Changed("sort-dir") {
			val, _ := cmd.Flags().GetString("sort-dir")
			optionalParams.WithSortDir(val)
		}

		if cmd.Flags().Changed("filter") {
			val, _ := cmd.Flags().GetString("filter")
			optionalParams.WithFilter(val)
		}

		if cmd.Flags().Changed("filter-status") {
			val, _ := cmd.Flags().GetString("filter-status")
			optionalParams.WithFilterStatus(val)
		}

		api := datadogV2.NewUsersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListUsers(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-users")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {

	ListUsersCmd.Flags().Int64("page-size", 0, "Size for a given page. The maximum allowed value is 100.")

	ListUsersCmd.Flags().Int64("page-number", 0, "Specific page number to return.")

	ListUsersCmd.Flags().String("sort", "", "User attribute to order results by. Sort order is ascending by default. Sort order is descending if the field is prefixed by a negative sign, for example 'sort=-name'. Options: 'name', 'modified_at', 'user_count'.")

	ListUsersCmd.Flags().String("sort-dir", "", "Direction of sort. Options: 'asc', 'desc'.")

	ListUsersCmd.Flags().String("filter", "", "Filter all users by the given string. Defaults to no filtering.")

	ListUsersCmd.Flags().String("filter-status", "", "Filter on status attribute. Comma separated list, with possible values 'Active', 'Pending', and 'Disabled'. Defaults to no filtering.")

	Cmd.AddCommand(ListUsersCmd)
}
