# tf-projects-bootstrap
Tool to help you to create new Terraform projects.


## How to install
> For now, it's with our friend `go build`, for the future I'll use goreleaser   

Clone the project and inside the directory:
1. Build:
`go build`

2. Give permission for execution to binary:
`chmod +x tf-projects-boostrap`

3. Execute:
`./tf-projects-boostrap`
## How to use

1. To create a generic project with default options:
```
$ tf-projects-boostrap init 
initializing your terraform project...
versions.tf created
main.tf created
variables.tf created
outputs.tf created
backend.tf created
```

2. To create a module project with default options:
```
$ tf-projects-boostrap init --type module
initializing your terraform project...
versions.tf created
main.tf created
variables.tf created
outputs.tf created
.editorconfig created
```

3. To create a project setting up a specific Terraform version:
```
$ ./tf-projects-boostrap init --tf-version "1.0.1"
initializing your terraform project...
versions.tf created
main.tf created
variables.tf created
outputs.tf created
backend.tf created
```
This will produce this `versions.tf` file:
```
$cat versions.tf
terraform {
        # Terraform required versions
        required_version = "~> 1.0.1"
        # Provider required versions (aws example below)
        #required_providers {
        #       aws = {
        #         source  = "hashicorp/aws"
        #         version = "~> 3.0"
        #       }
        #  } 
}   
```
