package section

import (
	"fmt"
	"strings"

	"github.com/daixiang0/gci/pkg/parse"
	"github.com/daixiang0/gci/pkg/specificity"
)

const ModulePrefixType = "module_prefix"

type ModulePrefix struct {
	Pkg string
}

func (d ModulePrefix) MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity {
	path := strings.Trim(spec.Path, `"`)

	if path != "" &&
		strings.HasPrefix(path, d.Pkg) {
		return specificity.ModulePrefix{Length: len(d.Pkg)}
	}

	return specificity.MisMatch{}
}

func (d ModulePrefix) String() string {
	return fmt.Sprintf("module_prefix(%s)", d.Pkg)
}

func (d ModulePrefix) Type() string {
	return ModulePrefixType
}
