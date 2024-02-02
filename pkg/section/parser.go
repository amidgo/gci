package section

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

var (
	ErrGoModFileNotFound    = errors.New("go module file not found")
	ErrFailedParseGoModFile = errors.New("failed parse go.mod file, wrong format")
)

func Parse(data []string, goModulePath string) (SectionList, error) {
	if len(data) == 0 {
		return nil, nil
	}

	list := make(SectionList, 0, len(data))

	errString := strings.Builder{}
	for _, d := range data {
		s := strings.ToLower(d)
		if len(s) == 0 {
			return nil, nil
		}

		if s == "default" {
			list = append(list, Default{})
		} else if s == "standard" {
			list = append(list, Standard{})
		} else if s == "newline" {
			list = append(list, NewLine{})
		} else if strings.HasPrefix(s, "prefix(") && len(d) > 8 {
			list = append(list, Custom{d[7 : len(d)-1]})
		} else if strings.HasPrefix(s, "commentline(") && len(d) > 13 {
			list = append(list, Custom{d[12 : len(d)-1]})
		} else if s == "dot" {
			list = append(list, Dot{})
		} else if s == "blank" {
			list = append(list, Blank{})
		} else if s == "alias" {
			list = append(list, Alias{})
		} else if s == "module" {
			sections, err := moduleSections(goModulePath)
			if err != nil {
				errString.WriteRune(' ')
				errString.WriteString(err.Error())

				continue
			}

			list = append(list, sections...)
		} else if strings.HasPrefix(s, "module_prefix(") && len(s) > 15 {
			sections, err := modulePrefixSections(goModulePath, s)
			if err != nil {
				errString.WriteRune(' ')
				errString.WriteString(err.Error())

				continue
			}

			list = append(list, sections...)
		} else {
			errString.WriteRune(' ')
			errString.WriteString(s)
		}
	}

	if errString.String() != "" {
		return nil, fmt.Errorf("invalid params:%s", errString.String())
	}

	return list, nil
}

func moduleSections(goModulePath string) ([]Section, error) {
	mdfile, err := modFile(goModulePath)
	if err != nil {
		return nil, err
	}

	sections := make([]Section, 0)

	for _, dep := range mdfile.Require {
		if dep.Indirect {
			continue
		}

		sections = append(sections, Module{Pkg: dep.Mod.Path})
	}

	return sections, nil
}

func modulePrefixSections(goModulePath string, s string) ([]Section, error) {
	mdfile, err := modFile(goModulePath)
	if err != nil {
		return nil, err
	}

	prefix := strings.TrimPrefix(s, "module_prefix(")
	prefix = strings.TrimSuffix(prefix, ")")

	sections := make([]Section, 0)

	for _, dep := range mdfile.Require {
		if dep.Indirect {
			continue
		}

		if !strings.HasPrefix(dep.Mod.Path, prefix) {
			continue
		}

		sections = append(sections, ModulePrefix{Pkg: dep.Mod.Path})
	}

	return sections, nil
}

func modFile(goModulePath string) (*modfile.File, error) {
	goModFileContent, err := os.ReadFile(goModulePath)
	if err != nil {
		return nil, ErrGoModFileNotFound
	}

	modFile, err := modfile.Parse(goModulePath, goModFileContent, nil)
	if err != nil {
		return nil, ErrFailedParseGoModFile
	}

	return modFile, nil
}
