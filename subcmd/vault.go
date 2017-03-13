package subcmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/ieee0824/thor/vault"
)

type vaultParam struct {
	Pass *string
	Path *string
}

func (v *vaultParam) Validate() error {
	if v.Pass == nil {
		return errors.New("pass is nil")
	} else if v.Path == nil {
		return errors.New("path is nil")
	}
	return nil
}

func parseVaultArgs(args []string) (*vaultParam, error) {
	var result = &vaultParam{}
	passPram, err := getValFromArgs(args, "-p")
	if err != nil {
		return nil, err
	}
	if 2 <= len(passPram) {
		return nil, errors.New("'-p' parameter can not be specified more than once.")
	} else if 0 == len(passPram) {
		return nil, errors.New("'-p' parameter is empty")
	}
	result.Pass = passPram[0]

	pathParam, err := getValFromArgs(args, "-f")
	if err != nil {
		return nil, err
	}
	if 2 <= len(pathParam) {
		return nil, errors.New("'-f' parameter can not be specified more than once.")
	} else if 0 == len(pathParam) {
		return nil, errors.New("'-f' parameter is empty")
	}
	result.Path = pathParam[0]
	return result, nil
}

type Vault struct{}

func (c *Vault) Help() string {
	return ""
}

func (c *Vault) Run(args []string) int {
	param, err := parseVaultArgs(args)
	if err != nil {
		log.Fatalln(err)
	}
	if err := param.Validate(); err != nil {
		log.Fatalln(err)
	}
	bin, err := ioutil.ReadFile(*param.Path)
	if err != nil {
		log.Fatalln(err)
	}
	vaulter := vault.New(bin)
	if err := vaulter.Encrypt(*param.Pass); err != nil {
		log.Fatalln(err)
	}
	vaultedJSON, err := json.Marshal(vaulter)
	if err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile(*param.Path, vaultedJSON, 0644); err != nil {
		log.Fatalln(err)
	}

	return 0
}

func (c *Vault) Synopsis() string {
	return ""
}
