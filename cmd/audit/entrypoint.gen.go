package audit

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "audit",
	Short: "audit endpoints",
}
