package processors

import (
	"fmt"
	"time"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

func Copyright(owner string) pgs.PostProcessor {
	return &copyrightPostProcessor{owner}
}

type copyrightPostProcessor struct {
	owner string
}

func (p *copyrightPostProcessor) Match(a pgs.Artifact) bool {
	switch a.(type) {
	case pgs.GeneratorFile, pgs.GeneratorTemplateFile, pgs.CustomFile, pgs.CustomTemplateFile:
		return true
	default:
		return false
	}
}

func (p *copyrightPostProcessor) Process(in []byte) ([]byte, error) {
	cmt := fmt.Sprintf("// Copyright © %d %s. All rights reserved\n",
		time.Now().Year(),
		p.owner)

	return append([]byte(cmt), in...), nil
}
