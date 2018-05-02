package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type cmdFlags struct {
	fs               *flag.FlagSet
	kubeConfig       string
	brigadeNamespace string
	showVersion      bool
}

func newCmdFlags() (*cmdFlags, error) {
	fls := &cmdFlags{
		fs: flag.NewFlagSet(os.Args[0], flag.ExitOnError),
	}
	err := fls.init()

	return fls, err
}
func (c *cmdFlags) init() error {
	var kubehome string

	if kubehome = os.Getenv("KUBECONFIG"); kubehome == "" {
		kubehome = filepath.Join(
			os.Getenv("HOME"), ".kube", "config",
		)
	}

	// register flags
	c.fs.StringVar(&c.kubeConfig, "kubeconfig", kubehome, "kubernetes configuration path, only used when development mode enabled")
	c.fs.StringVar(&c.brigadeNamespace, "namespace", "default", "kubernetes namespace where brigade is running")
	c.fs.BoolVar(&c.showVersion, "version", false, "show app version")

	// Parse flags
	if err := c.fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	if err := c.validate(); err != nil {
		return err
	}

	return nil
}

func (c *cmdFlags) validate() error {
	if c.brigadeNamespace == "" {
		return fmt.Errorf("namespace is required")
	}
	return nil
}
