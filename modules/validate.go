package modules

import (
	"github.com/envoyproxy/protoc-gen-validate/templates"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

const validatorName = "validator"

type validateModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func Validator() pgs.Module {
	return &validateModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *validateModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *validateModule) Name() string {
	return validatorName
}

func (m *validateModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		m.gen(t)
	}
	return m.Artifacts()
}

func (m *validateModule) gen(f pgs.File) {
	if len(f.Services()) == 0 || templates.Template(m.Parameters())["go"] == nil || len(templates.Template(m.Parameters())["go"]) == 0 {
		return
	}

	name := m.ctx.OutputPath(f).SetBase(f.InputPath().BaseName()).SetExt(".validate.pb.go")
	m.AddGeneratorTemplateFile(name.String(), templates.Template(m.Parameters())["go"][0], f)
}
