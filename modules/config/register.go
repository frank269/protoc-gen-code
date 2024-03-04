package config

import (
	"text/template"

	goTemplate "github.com/frank269/protoc-gen-code/modules/config/template"
	"github.com/frank269/protoc-gen-code/shared"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

func Register(tpl *template.Template, params pgs.Parameters) *template.Template {
	shared.RegisterFunctions(tpl, params)
	goTemplate.Register(tpl, params)
	template.Must(tpl.Parse(fileTpl))
	return tpl
}
