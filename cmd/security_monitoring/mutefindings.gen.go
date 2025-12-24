package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var MuteFindingsCmd = &cobra.Command{
	Use:   "mutefindings",
	Short: "Mute or unmute a batch of findings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/posture_management/findings")
		fmt.Println("OperationID: MuteFindings")
	},
}

func init() {
	Cmd.AddCommand(MuteFindingsCmd)
}
