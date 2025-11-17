package loaders

import (
	"github.com/SBPH-Matthew/testosterone-tracker/graph/model"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	UserLoader *dataloadgen.Loader[string, *model.User]
}
