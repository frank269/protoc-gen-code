package template

import (
	"fmt"
	"text/template"

	"github.com/frank269/protoc-gen-code/proto/extension"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := goSharedFuncs{pgsgo.InitContext(params)}

	tpl.Funcs(map[string]interface{}{
		"tag":            fns.Tags,
		"envName":        fns.EnvName,
		"upperCamelCase": pgsgo.PGGUpperCamelCase,
		"enabled":        fns.Enabled,
	})

	template.Must(tpl.New("struct").Parse(structTpl))
	template.Must(tpl.New("function").Parse(functionTpl))
}

type goSharedFuncs struct{ pgsgo.Context }

func (*goSharedFuncs) Tags(f pgs.Field) string {
	option := &extension.EnviromentOptions{}
	ok, err := f.Extension(extension.E_Env, &option)
	if err != nil {
		return fmt.Sprintf("`%s`", err.Error())
	}
	if ok {
		name := fmt.Sprintf("env:\"%s\"", option.Name)
		defaultValue := ""
		if option.DefaultValue != "" {
			defaultValue = fmt.Sprintf(" envDefault:\"%s\"", option.DefaultValue)
		}
		return fmt.Sprintf("`%s%s`", name, defaultValue)
	}
	return ""
}

func (fns goSharedFuncs) GoName(name pgs.Name) string {
	return string(pgsgo.PGGUpperCamelCase(name))
}

func (fns goSharedFuncs) Enabled(msg pgs.Message) (enabled bool, err error) {
	_, err = msg.Extension(extension.E_Enabled, &enabled)
	return
}

func (fns goSharedFuncs) EnvName(msg pgs.Message) pgs.Name {
	return msg.Name().UpperSnakeCase() + "Env"
}
