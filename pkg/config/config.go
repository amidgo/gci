package config

import (
	"sort"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/daixiang0/gci/pkg/section"
)

var defaultOrder = map[string]int{
	section.StandardType: 0,
	section.DefaultType:  1,
	section.CustomType:   2,
	section.BlankType:    3,
	section.DotType:      4,
	section.AliasType:    5,
}

type BoolConfig struct {
	NoInlineComments bool `yaml:"no-inlineComments"`
	NoPrefixComments bool `yaml:"no-prefixComments"`
	Debug            bool `yaml:"-"`
	SkipGenerated    bool `yaml:"skipGenerated"`
	SkipVendor       bool `yaml:"skipVendor"`
	CustomOrder      bool `yaml:"customOrder"`
}

type Config struct {
	BoolConfig
	Sections          section.SectionList
	SectionSeparators section.SectionList
}

type YamlConfig struct {
	Cfg                     BoolConfig `yaml:",inline"`
	SectionStrings          []string   `yaml:"sections"`
	SectionSeparatorStrings []string   `yaml:"sectionseparators"`
	ModulePath              string     `yaml:"modulePath"`
}

func (g YamlConfig) Parse() (*Config, error) {
	var err error

	modulePath := g.modulePath()

	sections, err := section.Parse(g.SectionStrings, modulePath)
	if err != nil {
		return nil, err
	}

	if sections == nil {
		sections = section.DefaultSections()
	}

	// if default order sorted sections
	if !g.Cfg.CustomOrder {
		sort.Slice(sections, func(i, j int) bool {
			sectionI, sectionJ := sections[i].Type(), sections[j].Type()

			if strings.Compare(sectionI, sectionJ) == 0 {
				return strings.Compare(sections[i].String(), sections[j].String()) < 0
			}
			return defaultOrder[sectionI] < defaultOrder[sectionJ]
		})
	}

	sectionSeparators, err := section.Parse(g.SectionSeparatorStrings, modulePath)
	if err != nil {
		return nil, err
	}

	if sectionSeparators == nil {
		sectionSeparators = section.DefaultSectionSeparators()
	}

	return &Config{g.Cfg, sections, sectionSeparators}, nil
}

func (g *YamlConfig) modulePath() string {
	if g.ModulePath == "" {
		return "go.mod"
	}

	return g.ModulePath
}

func ParseConfig(in string) (*Config, error) {
	config := YamlConfig{}

	err := yaml.Unmarshal([]byte(in), &config)
	if err != nil {
		return nil, err
	}

	gciCfg, err := config.Parse()
	if err != nil {
		return nil, err
	}

	return gciCfg, nil
}
