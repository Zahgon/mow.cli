package cli

import (
	"io"
	"os"

	"github.com/jawher/mow.cli/internal/container"
	"github.com/jawher/mow.cli/internal/flow"
)

/*
Cli represents the structure of a CLI app. It should be constructed using the App() function
*/
type Cli struct {
	*Cmd
	version *cliVersion
}

type cliVersion struct {
	version string
	option  *container.Container
}

/*
App creates a new and empty CLI app configured with the passed name and description.

name and description will be used to construct the help message for the app:

	Usage: $name [OPTIONS] COMMAND [arg...]

	$desc
*/
func App(name, desc string) *Cli { _ = "STUB: not implemented"; return nil }

/*
Version sets the version string of the CLI app together with the options that can be used to trigger
printing the version string via the CLI.

	Usage: appName --$name
	$version
*/
func (cli *Cli) Version(name, version string) { _ = "STUB: not implemented"; return }

func (cli *Cli) parse(args []string, entry, inFlow, outFlow *flow.Step) error {
	_ = "STUB: not implemented"
	// We overload Cmd.parse() and handle cases that only apply to the CLI command, like versioning
	// After that, we just call Cmd.parse() for the default behavior
	return nil
}

func (cli *Cli) versionSetAndRequested(args []string) bool { _ = "STUB: not implemented"; return false }

/*
PrintVersion prints the CLI app's version.
In most cases the library users won't need to call this method, unless
a more complex validation is needed.
*/
func (cli *Cli) PrintVersion() { _ = "STUB: not implemented"; return }

/*
Run uses the app configuration (specs, commands, ...) to parse the args slice
and to execute the matching command.

In case of an incorrect usage, and depending on the configured ErrorHandling policy,
it may return an error, panic or exit
*/
func (cli *Cli) Run(args []string) error { _ = "STUB: not implemented"; return nil }

/*
ActionCommand is a convenience function to configure a command with an action.

cmd.ActionCommand(_, _, myFunc) is equivalent to cmd.Command(_, _, func(cmd *cli.Cmd) { cmd.Action = myFunc })
*/
func ActionCommand(action func()) CmdInitializer {
	_ = "STUB: not implemented"
	return *new(CmdInitializer)
}

/*
Exit causes the app the exit with the specified exit code while giving the After interceptors a chance to run.
This should be used instead of os.Exit.
*/
func Exit(code int) { _ = "STUB: not implemented"; return }

var exiter = func(code int) {
	os.Exit(code)
}

var (
	stdOut io.Writer = os.Stdout
	stdErr io.Writer = os.Stderr
)
