package builder

import "errors"

const resourcesDirName string = "resources"

type ResourcesBuilder struct {
    DefaultSTRBuilder
}

func (b *ResourcesBuilder) Build(paths ...string) (string, error) {

    if len(paths) == 0 || paths[0] == "" {
        return "", errors.New("the parameter [paths] or the first element in [paths] is empty")
    }

    return b.DefaultSTRBuilder.Build(paths[0], resourcesDirName)
}