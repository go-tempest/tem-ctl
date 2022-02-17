package cmd

import (
    "github.com/spf13/cobra"
)

const (
    dirFlag         string = "dir"
    dirShorthand    string = "d"
    dirFlagUsage    string = "Specifies the run directory"
    dirDefaultValue string = ""
)

var targetDir string

// scaffoldCmd 脚手架子命令
var scaffoldCmd = &cobra.Command{
    Use:   "scaffold",
    Short: "Scaffolding tools",
    Long:  `Command-line tools for project scaffolding related operations`,
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
    rootCmd.AddCommand(scaffoldCmd)
    scaffoldCmd.PersistentFlags().StringVarP(&targetDir, dirFlag, dirShorthand,
        dirDefaultValue, dirFlagUsage)
}
