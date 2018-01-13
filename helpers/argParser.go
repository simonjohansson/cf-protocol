package helpers

import (
	"flag"
	"regexp"
	"code.cloudfoundry.org/cli/cf/errors"
)

type Options struct {
	ManifestPath string
	Postfix      string
	Domain       string
}

func ParseArgs(args []string) (Options, error) {
	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	manifestPath := flagSet.String("manifest", "", "Path to the manifest")
	postfix := flagSet.String("postfix", "", "Postfix to use push")
	domain := flagSet.String("domain", "", "Domain to use when pushing")

	err := flagSet.Parse(args[1:])
	if err != nil {
		return Options{}, err
	}

	mustBeANumber, _ := regexp.Compile("[0-9]+")
	if !mustBeANumber.MatchString(*postfix) {
		return Options{}, errors.New("postfix must be a number!")
	}

	return Options{
		*manifestPath,
		*postfix,
		*domain,
	}, nil
}
