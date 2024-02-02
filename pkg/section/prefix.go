package section

import (
	"fmt"
	"strings"

	"github.com/daixiang0/gci/pkg/parse"
	"github.com/daixiang0/gci/pkg/specificity"
)

type Custom struct {
	Prefix string
}

// CustomSeparator allows you to group multiple custom prefix together in the same section
// gci diff -s standard -s default -s prefix(github.com/company,gitlab.com/company,companysuffix)
const CustomSeparator = ","

const CustomType = "custom"

func (c Custom) MatchSpecificity(spec *parse.GciImports) specificity.MatchSpecificity {
	path := strings.Trim(spec.Path, `"`)

	for _, prefix := range strings.Split(c.Prefix, CustomSeparator) {
		prefix = strings.TrimSpace(prefix)

		if prefix == "" {
			continue
		}

		if strings.HasPrefix(path, prefix) {
			return specificity.Match{Length: len(prefix)}
		}
	}

	return specificity.MisMatch{}
}

func (c Custom) String() string {
	return fmt.Sprintf("prefix(%s)", c.Prefix)
}

func (c Custom) Type() string {
	return CustomType
}
