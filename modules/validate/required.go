package validate

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if {{ accessor . }} == nil {
			return {{ err . "value is required" }}
		}
	{{ end }}
`
