package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

const (
	defaultKeyfileName = "keyfile.json"
)

var (
	gitCommit = "" // Git SHA1 commit hash of the release (set via linker flags)

	app *cli.App // the main app instance
)

var ( // Commonly used command line flags.
	passphraseFlag = cli.StringFlag{
		Name:  "passwordfile",
		Usage: "the file that contains the passphrase for the keyfile",
	}

	jsonFlag = cli.BoolFlag{
		Name:  "json",
		Usage: "output JSON instead of human-readable format",
	}

	messageFlag = cli.StringFlag{
		Name:  "message",
		Usage: "the file that contains the message to sign/verify",
	}
)

// Configure the app instance.
func init() {
	app = utils.NewApp(gitCommit, "an Ethereum key manager")
	app.Commands = []cli.Command{
		commandGenerate,
		commandInspect,
		commandSignMessage,
		commandVerifyMessage,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
