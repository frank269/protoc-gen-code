package modules

import (
	"bytes"
	"fmt"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

type reportModule struct {
	*pgs.ModuleBase
}

func Report() pgs.Module {
	return &reportModule{&pgs.ModuleBase{}}
}

func (m *reportModule) Name() string {
	return "reporter"
}

func (m *reportModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}
	for _, f := range targets {
		m.Push(f.Name().String()).Debug("reporting")

		fmt.Fprintf(buf, "------ %v ------\n", f.Name())

		for i, msg := range f.AllMessages() {
			fmt.Fprintf(buf, "%03d. %v\n", i, msg.Name())
		}

		m.Pop()
	}

	m.OverwriteCustomFile("report.txt", buf.String(), 0644)

	return m.Artifacts()
}
