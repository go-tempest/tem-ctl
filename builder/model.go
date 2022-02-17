package builder

import "errors"

const modelDirName string = "model"

type ModelBuilder struct {
    DefaultSTRBuilder
}

func (b *ModelBuilder) Build(paths ...string) (string, error) {

    if len(paths) == 0 || paths[0] == "" {
        return "", errors.New("the parameter [paths] or the first element in [paths] is empty")
    }

    return b.DefaultSTRBuilder.Build(paths[0], modelDirName)
}