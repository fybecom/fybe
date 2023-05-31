package main

import (
	"fybe.com/cli/fybe/cmd"
	_ "fybe.com/cli/fybe/cmd/buckets"
	_ "fybe.com/cli/fybe/cmd/datacenters"
	_ "fybe.com/cli/fybe/cmd/images"
	_ "fybe.com/cli/fybe/cmd/instanceActions"
	_ "fybe.com/cli/fybe/cmd/instances"
	_ "fybe.com/cli/fybe/cmd/objectStorage"
	_ "fybe.com/cli/fybe/cmd/objects"
	_ "fybe.com/cli/fybe/cmd/roles"
	_ "fybe.com/cli/fybe/cmd/secrets"
	_ "fybe.com/cli/fybe/cmd/tagAssignment"
	_ "fybe.com/cli/fybe/cmd/tags"
	_ "fybe.com/cli/fybe/cmd/users"
	_ "fybe.com/cli/fybe/cmd/vpc"
)

func main() {
	cmd.Execute()
}
