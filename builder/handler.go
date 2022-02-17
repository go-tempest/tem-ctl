package builder

import "errors"

const handlerDirName string = "handler"

type HandlerBuilder struct {
    DefaultSTRBuilder
}

func (b *HandlerBuilder) Build(paths ...string) (string, error) {

    if len(paths) == 0 || paths[0] == "" {
        return "", errors.New("the parameter [paths] or the first element in [paths] is empty")
    }

    return b.DefaultSTRBuilder.Build(paths[0], handlerDirName)
}
