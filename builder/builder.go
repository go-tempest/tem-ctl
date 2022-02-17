package builder

import (
    "errors"
    "os"
    "path"
)

//BType Builder 构建器类型
type BType uint8

const (
    Project BType = iota
    Handler
    Service
    Model
    Resources
    ModFile
)

//STRBuilder 项目结构构建器接口
type STRBuilder interface {
    Build(paths ...string) (string, error)
}

type DefaultSTRBuilder struct {
}

func (b *DefaultSTRBuilder) Build(paths ...string) (string, error) {
    return MKdirs(paths...)
}

type FileBuilder interface {
    Build(parent string, data map[string]interface{}) error
}

type STRBuilderFactory struct {
}

func (f *STRBuilderFactory) CreateSTRBuilder(t BType) (builder STRBuilder) {
    switch t {
    case Project:
        builder = &ProjectBuilder{}
    case Handler:
        builder = &HandlerBuilder{}
    case Service:
        builder = &ServiceBuilder{}
    case Model:
        builder = &ModelBuilder{}
    case Resources:
        builder = &ResourcesBuilder{}
    }
    return
}

func (f *STRBuilderFactory) CreateFileBuilder(t BType) (builder FileBuilder) {
    switch t {
    case ModFile:
        builder = &ModFileBuilder{}
    }
    return
}

func MKdirs(paths ...string) (string, error) {

    if len(paths) == 0 {
        return "", errors.New("argument [paths] is empty")
    }

    d := path.Join(paths...)
    _ = os.MkdirAll(d, os.ModePerm)
    return d, nil
}
