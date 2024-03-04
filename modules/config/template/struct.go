package template

const structTpl = `
type {{ envName . }} struct {
	{{ range .Fields }}
	{{ upperCamelCase .Name }}	{{ typ . }} {{ tag . }}{{ end }}
}`
