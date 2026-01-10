package data_deletion

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDataDeletionRequestsCmd = &cobra.Command{
	Use:     "get-data-deletion-requests",
	Aliases: []string{"get-requests"},
	Short:   "Gets a list of data deletion requests",
	Long: `Gets a list of data deletion requests
Documentation: https://docs.datadoghq.com/api/latest/data-deletion/#get-data-deletion-requests`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetDataDeletionsResponseBody
		var err error

		optionalParams := datadogV2.NewGetDataDeletionRequestsOptionalParameters()

		if cmd.Flags().Changed("next-page") {
			val, _ := cmd.Flags().GetString("next-page")
			optionalParams.WithNextPage(val)
		}

		if cmd.Flags().Changed("product") {
			val, _ := cmd.Flags().GetString("product")
			optionalParams.WithProduct(val)
		}

		if cmd.Flags().Changed("query") {
			val, _ := cmd.Flags().GetString("query")
			optionalParams.WithQuery(val)
		}

		if cmd.Flags().Changed("status") {
			val, _ := cmd.Flags().GetString("status")
			optionalParams.WithStatus(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		api := datadogV2.NewDataDeletionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDataDeletionRequests(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to get-data-deletion-requests")

		fmt.Println(cmdutil.FormatJSON(res, "data_deletion"))
	},
}

func init() {

	GetDataDeletionRequestsCmd.Flags().String("next-page", "", "The next page of the previous search. If the next_page parameter is included, the rest of the query elements are ignored.")

	GetDataDeletionRequestsCmd.Flags().String("product", "", "Retrieve only the requests related to the given product.")

	GetDataDeletionRequestsCmd.Flags().String("query", "", "Retrieve only the requests that matches the given query.")

	GetDataDeletionRequestsCmd.Flags().String("status", "", "Retrieve only the requests with the given status.")

	GetDataDeletionRequestsCmd.Flags().Int64("page-size", 0, "Sets the page size of the search.")

	Cmd.AddCommand(GetDataDeletionRequestsCmd)
}
