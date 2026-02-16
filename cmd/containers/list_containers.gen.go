package containers

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListContainersCmd = &cobra.Command{
	Use:     "list-containers",
	Aliases: []string{"list"},
	Short:   "Get All Containers",
	Long: `Get All Containers
Documentation: https://docs.datadoghq.com/api/latest/containers/#list-containers`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ContainersResponse
		var err error

		optionalParams := datadogV2.NewListContainersOptionalParameters()

		if cmd.Flags().Changed("filter-tags") {
			val, _ := cmd.Flags().GetString("filter-tags")
			optionalParams.WithFilterTags(val)
		}

		if cmd.Flags().Changed("group-by") {
			val, _ := cmd.Flags().GetString("group-by")
			optionalParams.WithGroupBy(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("page-cursor") {
			val, _ := cmd.Flags().GetString("page-cursor")
			optionalParams.WithPageCursor(val)
		}

		api := datadogV2.NewContainersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListContainers(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-containers")

		fmt.Println(cmdutil.FormatJSON(res, "containers"))
	},
}

func init() {

	ListContainersCmd.Flags().String("filter-tags", "", "Comma-separated list of tags to filter containers by.")

	ListContainersCmd.Flags().String("group-by", "", "Comma-separated list of tags to group containers by.")

	ListContainersCmd.Flags().String("sort", "", "Attribute to sort containers by.")

	ListContainersCmd.Flags().Int64("page-size", 0, "Maximum number of results returned.")

	ListContainersCmd.Flags().String("page-cursor", "", "String to query the next page of results. This key is provided with each valid response from the API in 'meta.pagination.next_cursor'.")

	Cmd.AddCommand(ListContainersCmd)
}
