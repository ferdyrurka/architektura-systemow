package main

import (
	"ferdyrurka/architektura-systemow/action"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "intel-8086-simulator"}
	rootCmd.AddCommand(action.InitCli())
	rootCmd.AddCommand(action.InitSti())
	rootCmd.AddCommand(action.InitMov())
	rootCmd.AddCommand(action.InitXchg())
	rootCmd.Execute()
}