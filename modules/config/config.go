package config

import (
	"text/template"

	psg "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

const moduleName = "model"

type configModule struct {
	*psg.ModuleBase
	ctx pgsgo.Context
}

func Module() psg.Module {
	return &configModule{ModuleBase: &psg.ModuleBase{}}
}

func (m *configModule) Name() string {
	return moduleName
}

func (m *configModule) InitContext(c psg.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *configModule) Execute(targets map[string]psg.File, packages map[string]psg.Package) []psg.Artifact {
	for _, t := range targets {
		m.generate(t)
	}

	return m.Artifacts()
}

func (m *configModule) generate(f psg.File) {
	if len(f.Messages()) == 0 {
		return
	}
	name := m.ctx.OutputPath(f).SetBase(f.InputPath().BaseName()).SetExt(".config.pb.go")
	m.AddGeneratorTemplateFile(name.String(), Register(template.New(moduleName), m.Parameters()), f)
}
