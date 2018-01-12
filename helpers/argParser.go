package helpers

import "flag"

func ParseArgs(logger Logger, flagset *flag.FlagSet, args []string) error {
	err := flagset.Parse(args[1:])
	if err != nil {
		return err
	}

	return nil
}
