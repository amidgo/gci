package section

import (
	"strings"

	"github.com/daixiang0/gci/pkg/parse"
	"github.com/daixiang0/gci/pkg/specificity"
)

const ModuleType = "module"

type Module struct {
	Pkg string
}

func (d Module) MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity {
	path := strings.Trim(spec.Path, `"`)

	if strings.HasPrefix(path, d.Pkg) {
		return specificity.Module{Length: len(d.Pkg)}
	}

	return specificity.MisMatch{}
}

func (d Module) String() string {
	return ModuleType
}

func (d Module) Type() string {
	return ModuleType
}
