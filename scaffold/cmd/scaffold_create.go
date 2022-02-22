package cmd

import (
    "fmt"
    "github.com/go-tempest/tem-ctl/builder"
    "github.com/spf13/cobra"
    "os"
)

const (
    nameFlag         string = "name"
    nameShorthand    string = "n"
    nameFlagUsage    string = "Specifies the project name"
    nameDefaultValue string = ""

    goVersionFlag      string = "go-version"
    goVersionShorthand string = "v"
    goVersionFlagUsage string = "Specifies the go version that the project depends on"

    goDefaultVersion string = "1.17"

    goVersionPropName string = "goVersion"
    projectPropName   string = "projectName"
)

var name, goVersion string

// createCmd 创建项目的指令
var createCmd = &cobra.Command{
    Use: "create",
    Long: `Used to create a project directory structure based on the tempest specification. For example:

lauevan@lauevan-mbp  ~/tmp tempest-cli scaffold create /tmp/test -n user -v 1.17`,
    Run: func(cmd *cobra.Command, args []string) {
        dir, err := getTargetDir(targetDir)
        checkErr(nil, err)
        createSTR(dir, name)
    },
}

//createSTR 创建项目结构
func createSTR(dir, name string) {

    if dir == "" || name == "" {
        os.Exit(1)
    }

    fmt.Println("[tempest-cli] [scaffold] [create] >> [Start]", "Start creating the project structure...")

    b := new(builder.STRBuilderFactory)
    projectDir, err := b.CreateSTRBuilder(builder.Project).Build(dir, name)
    checkErr(nil, err)

    checkErr(b.CreateSTRBuilder(builder.Handler).Build(projectDir))
    checkErr(b.CreateSTRBuilder(builder.Service).Build(projectDir))
    checkErr(b.CreateSTRBuilder(builder.Model).Build(projectDir))
    checkErr(b.CreateSTRBuilder(builder.Resources).Build(projectDir))

    err = b.CreateFileBuilder(builder.ModFile).Build(projectDir, map[string]interface{}{
        projectPropName:   name,
        goVersionPropName: goVersion,
    })
    checkErr(nil, err)

    fmt.Println("[tempest-cli] [scaffold] [create] >> [Done]", "The project structure is created")
}

func init() {
    scaffoldCmd.AddCommand(createCmd)
    createCmd.Flags().StringVarP(&name, nameFlag, nameShorthand, nameDefaultValue, nameFlagUsage)
    createCmd.Flags().StringVarP(&goVersion, goVersionFlag, goVersionShorthand, goDefaultVersion, goVersionFlagUsage)
    err := createCmd.MarkFlagRequired(nameFlag)
    checkErr(nil, err)
}

func checkErr(_ interface{}, err error) {
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "[tempest-cli] [scaffold] [create] >> "+
            "Failed to create project structure, error is [%v]\n", err)
        os.Exit(1)
    }
}
