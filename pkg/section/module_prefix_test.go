package section

import (
	"testing"

	"github.com/daixiang0/gci/pkg/specificity"
)

func TestModulePrefixx(t *testing.T) {
	testCases := []specificityTestData{
		{`"foo/pkg/bar"`, ModulePrefix{"foo"}, specificity.ModulePrefix{Length: 3}},
		{`""`, Module{"foo"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, ModulePrefix{"bar"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, ModulePrefix{"github.com/foo/bar"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, ModulePrefix{"github.com/foo"}, specificity.MisMatch{}},
		{`"foo/pkg/bar"`, ModulePrefix{"github.com/bar"}, specificity.MisMatch{}},
	}
	testSpecificity(t, testCases)
}
