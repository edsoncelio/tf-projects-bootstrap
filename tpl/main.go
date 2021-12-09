package tpl

func VersionsTemplate() []byte {
	return []byte(`terraform {
	required_version = "~> {{ .Version}}" 
}`)
}
