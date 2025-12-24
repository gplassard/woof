package users

func init() {
	Cmd.AddCommand(
		listUsers,
		listUserOrganizations,
	)
}
