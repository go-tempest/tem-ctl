package builder

import (
    codetpl "code-tpl/scaffold"
    "errors"
    "fmt"
    "os"
    "path"
    "text/template"
)

const (
    modFileName    string = "go.mod"
    modFileTplName        = modFileName + `.tpl`
)

type ModFileBuilder struct {
    DefaultSTRBuilder
}

func (b *ModFileBuilder) Build(parent string, data map[string]interface{}) error {

    if parent == "" {
        return errors.New("the parameter [parent] is empty")
    }

    modFilePath := path.Join(parent, modFileName)

    fmt.Printf(
        ">> The full path to the go.mod file in the project is [%s]\n",
        modFilePath,
    )

    if _, err := os.Stat(modFilePath); os.IsNotExist(err) {

        f, _ := os.Create(modFilePath)
        defer func() { _ = f.Close() }()

        t, _ := template.New(modFileTplName).Parse(codetpl.GoModTpl)
        _ = t.Execute(f, data)
        return nil
    }

    fmt.Println(">> The go.mod file already exists in the project and will not be created")
    return nil
}
