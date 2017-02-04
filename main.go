/* bub

   BitBucket / Stash version of Hub

   Add the following functionality to the git command.

   * pull-request
   * create
   * browse
   * compare
   * fork ???
   * ci-status ???

   All other commands pass-through to the git command normally.

*/
package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "bub"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:   "pull-request",
			Usage:  "create a pull-request",
			Action: cmdPullRequest,
		},
	}

	app.Run(os.Args)
}
