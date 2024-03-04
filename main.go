package main

import (
	"github.com/frank269/protoc-gen-code/modules/config"
	"github.com/frank269/protoc-gen-code/modules/jsonify"
	"github.com/frank269/protoc-gen-code/modules/validate"
	"github.com/frank269/protoc-gen-code/processors"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

func main() {
	pgs.Init(
	// pgs.DebugEnv("DEBUG"),
	// pgs.DebugMode(),
	).RegisterModule(
		jsonify.JSONify(),
		// service.ServiceClient(),
		config.Module(),
		validate.Validator(),
	).RegisterPostProcessor(
		pgsgo.GoImports(),
		pgsgo.GoFmt(),
		processors.Version("1.0.0"),
		// processors.Copyright("MPT"),
		// processors.Author("Tiendc", "tiendc@mpt.com.vn"),
	).Render()
}
