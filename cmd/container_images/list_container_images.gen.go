package container_images

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListContainerImagesCmd = &cobra.Command{
	Use:     "list-container-images",
	Aliases: []string{"list"},
	Short:   "Get all Container Images",
	Long: `Get all Container Images
Documentation: https://docs.datadoghq.com/api/latest/container-images/#list-container-images`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ContainerImagesResponse
		var err error

		optionalParams := datadogV2.NewListContainerImagesOptionalParameters()

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

		api := datadogV2.NewContainerImagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListContainerImages(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-container-images")

		fmt.Println(cmdutil.FormatJSON(res, "container_images"))
	},
}

func init() {

	ListContainerImagesCmd.Flags().String("filter-tags", "", "Comma-separated list of tags to filter Container Images by.")

	ListContainerImagesCmd.Flags().String("group-by", "", "Comma-separated list of tags to group Container Images by.")

	ListContainerImagesCmd.Flags().String("sort", "", "Attribute to sort Container Images by.")

	ListContainerImagesCmd.Flags().Int64("page-size", 0, "Maximum number of results returned.")

	ListContainerImagesCmd.Flags().String("page-cursor", "", "String to query the next page of results. This key is provided with each valid response from the API in 'meta.pagination.next_cursor'.")

	Cmd.AddCommand(ListContainerImagesCmd)
}
