package cmd

import (
    "os"

    "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use: "tem-ctl",
    Long: `tem-ctl is a command-line tool that integrates project scaffolding, 
annotated code generation, and microservice governance capabilities`,
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getTargetDir(target string) (dir string, err error) {

    if target == "" {
        dir, err = os.Getwd()
    } else {
        dir = target
    }

    return dir, err
}
