package processors

import (
	"fmt"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

func Author(author string, email string) pgs.PostProcessor {
	return &authorPostProcessor{author, email}
}

type authorPostProcessor struct {
	author string
	email  string
}

func (p *authorPostProcessor) Match(a pgs.Artifact) bool {
	switch a.(type) {
	case pgs.GeneratorFile, pgs.GeneratorTemplateFile, pgs.CustomFile, pgs.CustomTemplateFile:
		return true
	default:
		return false
	}
}

func (p *authorPostProcessor) Process(in []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("// Author %s (%s)\n%s", p.author, p.email, in)), nil
}
