package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Cmd) newRemoveCmd() *cobra.Command {
	removeCmd := &cobra.Command{
		Use:     "remove [tracking number]",
		Aliases: []string{"rm"},
		Short:   "Remove the package",
		Long:    "Remove a package from the list.",
		Example: "  nimotsu remove 112233445566",
		Args:    cobra.ExactValidArgs(1),
		RunE:    c.execRemoveCmd,
	}

	removeCmd.AddCommand(c.newRemoveAllCmd())

	return removeCmd
}

func (c *Cmd) execRemoveCmd(cmd *cobra.Command, args []string) error {
	err := c.list.RemoveItem(args[0])
	if err != nil {
		return err
	}

	c.list.Save()

	fmt.Println("Success: removed " + args[0])
	return nil
}
