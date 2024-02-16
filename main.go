package main

import (
	modules "protoc-gen-code/modules"
	processors "protoc-gen-code/processors"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

func main() {
	pgs.Init(
	// pgs.DebugEnv("DEBUG"),
	// pgs.DebugMode(),
	).RegisterModule(
		modules.JSONify(),
		modules.Validator(),
		// modules.Service(),
		modules.Config(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
		processors.Version("1.0.0"),
		processors.Copyright("MPT"),
		processors.Author("Tiendc", "tiendc@mpt.com.vn"),
	).Render()
}
