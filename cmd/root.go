package cmd

import (
	"errors"
	"log"
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func Root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		NewCreateSecretCommand(),
		NewGetSecretCommand(),
	}

	subcommand := args[0]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			if err := cmd.Init(args[1:]); err != nil {
				log.SetFlags(0)
				log.Println(err.Error())
				return err
			}
			err := cmd.Run()
			if err != nil {
				log.SetFlags(0)
				log.Println(err.Error())
			}
			return err
		}
	}

	err := errors.New("Unknown subcommand: " + subcommand)
	log.SetFlags(0)
	log.Println(err.Error())
	return err
}
