package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/jobtalk/pnzr/subcmd/deploy"
	"github.com/jobtalk/pnzr/subcmd/update"
	"github.com/jobtalk/pnzr/subcmd/vault"
	"github.com/jobtalk/pnzr/subcmd/vault_edit"
	"github.com/jobtalk/pnzr/subcmd/vault_view"
	"github.com/jobtalk/pnzr/vars"
	"github.com/mitchellh/cli"
)

var (
	VERSION    string
	BUILD_DATE string
	BUILD_OS   string
)

func generateBuildInfo() string {
	ret := fmt.Sprintf("Build version: %s\n", VERSION)
	ret += fmt.Sprintf("Go version: %s\n", runtime.Version())
	ret += fmt.Sprintf("Build Date: %s\n", BUILD_DATE)
	ret += fmt.Sprintf("Build OS: %s\n", BUILD_OS)
	return ret
}

func init() {
	if VERSION == "" {
		VERSION = "unknown"
	}
	vars.VERSION = VERSION
	vars.BUILD_DATE = BUILD_DATE
	vars.BUILD_OS = BUILD_OS

	VERSION = generateBuildInfo()
	log.SetFlags(log.Llongfile)

}

func main() {
	c := cli.NewCLI("pnzr", VERSION)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"deploy": func() (cli.Command, error) {
			return &deploy.Deploy{}, nil
		},
		"vault": func() (cli.Command, error) {
			return &vault.Vault{}, nil
		},
		"update": func() (cli.Command, error) {
			return &update.Update{}, nil
		},
		"vault-edit": func() (cli.Command, error) {
			return &vedit.VaultEdit{}, nil
		},
		"vault-view": func() (cli.Command, error) {
			return &vview.VaultView{}, nil
		},
	}
	exitCode, err := c.Run()
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(exitCode)
}
