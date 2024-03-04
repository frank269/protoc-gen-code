package template

import (
	"text/template"

	"github.com/frank269/protoc-gen-code/modules/validate/common"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	tpl.Funcs(map[string]interface{}{
		"disabled":    common.Disabled,
		"required":    common.RequiredOneOf,
		"context":     common.RulesContext,
		"render":      common.Render(tpl),
		"has":         common.Has,
		"needs":       common.Needs,
		"accessor":    common.Accessor,
		"err":         common.Err,
		"errCause":    common.ErrCause,
		"errIdx":      common.ErrIdx,
		"errIdxCause": common.ErrIdxCause,
		"errname":     common.ErrName,
		"unwrap":      common.Unwrap,
	})

	template.Must(tpl.New("msg").Parse(msgTpl))
	template.Must(tpl.New("const").Parse(constTpl))
	template.Must(tpl.New("ltgt").Parse(ltgtTpl))
	template.Must(tpl.New("in").Parse(inTpl))

	template.Must(tpl.New("none").Parse(noneTpl))
	template.Must(tpl.New("float").Parse(numTpl))
	template.Must(tpl.New("double").Parse(numTpl))
	template.Must(tpl.New("int32").Parse(numTpl))
	template.Must(tpl.New("int64").Parse(numTpl))
	template.Must(tpl.New("uint32").Parse(numTpl))
	template.Must(tpl.New("uint64").Parse(numTpl))
	template.Must(tpl.New("sint32").Parse(numTpl))
	template.Must(tpl.New("sint64").Parse(numTpl))
	template.Must(tpl.New("fixed32").Parse(numTpl))
	template.Must(tpl.New("fixed64").Parse(numTpl))
	template.Must(tpl.New("sfixed32").Parse(numTpl))
	template.Must(tpl.New("sfixed64").Parse(numTpl))

	template.Must(tpl.New("bool").Parse(constTpl))
	template.Must(tpl.New("string").Parse(strTpl))
	template.Must(tpl.New("bytes").Parse(bytesTpl))

	template.Must(tpl.New("email").Parse(emailTpl))
	template.Must(tpl.New("hostname").Parse(hostTpl))
	template.Must(tpl.New("address").Parse(hostTpl))
	template.Must(tpl.New("uuid").Parse(uuidTpl))

	template.Must(tpl.New("enum").Parse(enumTpl))
	template.Must(tpl.New("repeated").Parse(repTpl))
	template.Must(tpl.New("map").Parse(mapTpl))

	template.Must(tpl.New("any").Parse(anyTpl))
	template.Must(tpl.New("timestampcmp").Parse(timestampcmpTpl))
	template.Must(tpl.New("durationcmp").Parse(durationcmpTpl))

	template.Must(tpl.New("wrapper").Parse(wrapperTpl))
}
