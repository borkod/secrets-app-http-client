package cmd

import (
	"errors"
	"flag"
	"log"
)

func NewGetSecretCommand() *GetSecretCommand {
	gc := &GetSecretCommand{
		fs: flag.NewFlagSet("view", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.req.id, "id", "", "secret id")
	gc.fs.StringVar(&gc.req.url, "url", "", "api url")
	gc.req.action = "GET"
	return gc
}

type GetSecretCommand struct {
	fs *flag.FlagSet

	req request
}

func (g *GetSecretCommand) Name() string {
	return g.fs.Name()
}

func (g *GetSecretCommand) Init(args []string) error {
	if err := g.fs.Parse(args); err != nil {
		return err
	}

	if len(g.req.id) == 0 {
		return errors.New("-id not set")
	}
	if len(g.req.url) == 0 {
		return errors.New("-url not set")
	}
	if len(g.req.data) > 0 {
		return errors.New("-data cannot be set")
	}

	return nil
}

func (g *GetSecretCommand) Run() error {
	resp, err := g.req.processRequest()
	if err != nil {
		return err
	}
	log.SetFlags(0)
	log.Println(resp)
	return nil
}
