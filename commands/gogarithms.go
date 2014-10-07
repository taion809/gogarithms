package commands

import (
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{Use: "gogarithms"}

func Execute() {
	initCommand.Execute()
}
