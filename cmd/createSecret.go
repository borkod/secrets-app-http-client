package cmd

import (
	"errors"
	"flag"
	"log"
)

func NewCreateSecretCommand() *CreateSecretCommand {
	gc := &CreateSecretCommand{
		fs: flag.NewFlagSet("create", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.req.data, "data", "", "secret data")
	gc.fs.StringVar(&gc.req.url, "url", "", "api url")
	gc.req.action = "POST"
	return gc
}

type CreateSecretCommand struct {
	fs *flag.FlagSet

	req request
}

func (g *CreateSecretCommand) Name() string {
	return g.fs.Name()
}

func (g *CreateSecretCommand) Init(args []string) error {
	if err := g.fs.Parse(args); err != nil {
		return err
	}

	if len(g.req.data) == 0 {
		return errors.New("-data not set")
	}
	if len(g.req.url) == 0 {
		return errors.New("-url not set")
	}
	if len(g.req.id) > 0 {
		return errors.New("-id cannot be set")
	}

	return nil
}

func (g *CreateSecretCommand) Run() error {
	resp, err := g.req.processRequest()
	if err != nil {
		return err
	}
	log.SetFlags(0)
	log.Println(resp)
	return nil
}
