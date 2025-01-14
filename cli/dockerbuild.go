package cli

import (
	"fmt"
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/vsoch/uptodate/parsers/docker"
	"github.com/vsoch/uptodate/utils"
)

// Args and flags for generate
type DockerBuildArgs struct {
	Root []string `zero:"true" desc:"A root directory to parse."`
}

type DockerBuildFlags struct {
	Changes bool   `long:"changes" desc:"Only consider changed uptodate files"`
	Branch  string `long:"branch" desc:"Branch to compare HEAD against, defaults to main"`
}

// Dockerfile updates one or more Dockerfile
var DockerBuild = cmd.Sub{
	Name:  "dockerbuild",
	Alias: "db",
	Short: "Generate a matrix of builds for GitHub Actions or Similar.",
	Flags: &DockerBuildFlags{},
	Args:  &DockerBuildArgs{},
	Run:   RunDockerBuild,
}

func init() {
	cmd.Register(&DockerBuild)
}

// RunDockerfile updates one or more Dockerfile
func RunDockerBuild(r *cmd.Root, c *cmd.Sub) {

	args := c.Args.(*DockerBuildArgs)
	flags := c.Flags.(*DockerBuildFlags)

	// If no root provided, assume parsing the PWD
	if len(args.Root) == 0 {
		args.Root = []string{utils.GetPwd()}
	}

	// Print the logo!
	fmt.Println(utils.GetLogo() + "                     dockerbuild\n")

	// Set default branch
	if flags.Branch == "" {
		flags.Branch = "main"
	}

	// Update the dockerfiles with a Dockerfile parser
	parser := docker.DockerBuildParser{}
	parser.Parse(args.Root[0], flags.Changes, flags.Branch)

}
