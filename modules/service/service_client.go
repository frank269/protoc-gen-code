package service

import (
	"html/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

const serviceClientName = "service"

type serviceClientModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func ServiceClient() pgs.Module {
	return &serviceClientModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *serviceClientModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("service").Funcs(map[string]interface{}{
		"package": m.ctx.PackageName,
		"name":    m.ctx.Name,
	})

	m.tpl = template.Must(tpl.Parse(serviceClientTpl))
}

func (m *serviceClientModule) Name() string {
	return serviceClientName
}

func (m *serviceClientModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		m.gen(t)
	}
	return m.Artifacts()
}

func (m *serviceClientModule) gen(f pgs.File) {
	if len(f.Services()) == 0 {
		return
	}

	name := m.ctx.OutputPath(f).SetBase(f.InputPath().BaseName()).SetExt(".svc.client.pb.go")
	m.AddGeneratorTemplateFile(name.String(), m.tpl, f)
}

const serviceClientTpl = `// Code generated by protoc-gen-code. DO NOT EDIT.
// source: {{ .Name }}

package {{ package . }}
`
