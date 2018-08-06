package cmd

import (
	"fmt"
	"os"
	// "syscall"

	"github.com/spf13/cobra"
)

func change(cmd *cobra.Command, args []string) error {
	file, source, err := loadSource(os.O_RDONLY)
	if err != nil {
		return fmt.Errorf("change: %v", err)
	}
	defer file.Close()

	_, path, err := source.match(args[0])
	if err != nil {
		return fmt.Errorf("change: %v", err)
	}
	cmd.Print(path)
	// if err := os.Chdir(path); err != nil {
	// 	return err
	// }
	// shell := os.Getenv("SHELL")
	// if err := syscall.Exec(shell, []string{shell}, syscall.Environ()); err != nil {
	// 	return err
	// }

	return nil
}

func cmdChange() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "change",
		Short: "Change the working directory with the second name",
		Args:  cobra.MinimumNArgs(1),
		RunE:  change,
	}

	return cmd
}
