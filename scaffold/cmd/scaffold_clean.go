package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "io/ioutil"
    "os"
    "path"
)

const (
    projectsFlag      string = "projects"
    projectsShorthand string = "p"
    projectsFlagUsage string = "Specifies the name of the project to delete(separate multiple projects with ',')"
)

var projects []string

// cleanCmd 清理项目的指令
var cleanCmd = &cobra.Command{
    Use:  "clean",
    Long: `Used to clear project structure files created under a given directory`,
    Run: func(cmd *cobra.Command, args []string) {

        dir, err := getTargetDir(targetDir)
        if err != nil {
            _, _ = fmt.Fprintf(
                os.Stderr,
                "[tempest-cli] [scaffold] [clean] >> Failed to get the target run directory, the error is [%v]\n",
                err)
            os.Exit(1)
        }

        var subProjects []string

        if len(projects) != 0 {
            subProjects = projects
        } else {
            subProjects = getAllProjectsUnder(dir)
        }

        if len(subProjects) != 0 {
            clean(dir, subProjects...)
        }
    },
}

func clean(parent string, projects ...string) {
    for _, p := range projects {
        err := os.RemoveAll(path.Join(parent, p))
        if err != nil {
            _, _ = fmt.Fprintf(
                os.Stderr,
                "[tempest-cli] [scaffold] [clean] >> Failed to delete the specified project [%s], the error is [%v]\n",
                p,
                err)
            continue
        }
    }
}

func getAllProjectsUnder(dir string) []string {

    var ps []string

    if dir == "" {
        return ps
    }

    fss, err := ioutil.ReadDir(dir)
    if err != nil || len(fss) == 0 {
        return ps
    }

    for _, fs := range fss {
        if fs.IsDir() {
            ps = append(ps, fs.Name())
        }
    }

    return ps
}

func init() {
    scaffoldCmd.AddCommand(cleanCmd)
    cleanCmd.Flags().StringSliceVarP(&projects, projectsFlag, projectsShorthand, make([]string, 0), projectsFlagUsage)
}
