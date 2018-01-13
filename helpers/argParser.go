package helpers

import "flag"

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

	return Options{
		*manifestPath,
		*postfix,
		*domain,
	}, nil
}
