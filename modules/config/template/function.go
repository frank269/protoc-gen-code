package template

const functionTpl = `func (c *{{ envName .}}) FromEnv(filenames ...string) error {
	if len(filenames) == 0 {
		filenames = []string{".env"}
	}
	for _, filename := range filenames {
		err := godotenv.Load(filename)
		if err != nil {
			return err
		}
	}
	if err := env.Parse(c); err != nil {
		return err
	}

	return nil
}

func (c *{{envName .}}) AsProto() (*{{name .}}, error) {
	x := new({{name .}}) {{ range .Fields }}
	x.{{ upperCamelCase .Name }} = c.{{ upperCamelCase .Name }} {{ end }}
	return x, nil
}
`
