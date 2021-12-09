package tpl

func VersionsTemplate() []byte {
	return []byte(`terraform {
	# Terraform required versions
	required_version = "~> {{ .Version}}"
	# Provider required versions (aws example below)
	#required_providers {
	#	aws = {
	#	  source  = "hashicorp/aws"
	#	  version = "~> 3.0"
	#	}
	#  } 
}`)
}

func EditorCfgTemplate() []byte {
	return []byte(`# EditorConfig is awesome: http://EditorConfig.org
# Uses editorconfig to maintain consistent coding styles
	
# top-most EditorConfig file
root = true
	
# Unix-style newlines with a newline ending every file
[*]
charset = utf-8
end_of_line = lf
indent_size = 2
indent_style = space
insert_final_newline = true
max_line_length = 80
trim_trailing_whitespace = true
	
[*.{tf,tfvars}]
indent_size = 2
indent_style = space
`)
}

func BackendTemplate() []byte {
	return []byte(`terraform {
	backend "s3" {
		  #bucket = "mybucket" Your bucket name, must be already created!
		  #key    = "path/to/my/key" # you state key in the bucket
		  #region = "us-east-1" # the region
	}
}`)
}

func VariablesTemplate() []byte {
	return []byte(`# Variables
#variable "my_variable" {
	#description = "My awesome variable"
	#type    = string
	#default = "my awesome variable"
#}`)
}

func OutputsTemplate() []byte {
	return []byte(`# Outputs
#output "my_output" {
	#description = "My awesome output"
	#value = "my awesome output"
#}`)
}
