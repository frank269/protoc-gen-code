package validate

import (
	"text/template"

	goTemplate "github.com/frank269/protoc-gen-code/modules/validate/template"
	"github.com/frank269/protoc-gen-code/shared"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

func Register(tpl *template.Template, params pgs.Parameters) *template.Template {
	shared.RegisterFunctions(tpl, params)
	goTemplate.Register(tpl, params)
	template.Must(tpl.Parse(fileTpl))
	template.Must(tpl.New("required").Parse(requiredTpl))
	template.Must(tpl.New("timestamp").Parse(timestampTpl))
	template.Must(tpl.New("duration").Parse(durationTpl))
	template.Must(tpl.New("message").Parse(messageTpl))
	return tpl
}
