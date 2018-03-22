package main

import (
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

type cmdFlags struct {
	fs               *flag.FlagSet
	kubeConfig       string
	brigadeNamespace string
}

func newCmdFlags() *cmdFlags {
	fls := &cmdFlags{
		fs: flag.NewFlagSet(os.Args[0], flag.ExitOnError),
	}
	fls.init()
	return fls
}
func (c *cmdFlags) init() {
	kubehome := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// register flags
	c.fs.StringVar(&c.kubeConfig, "kubeconfig", kubehome, "kubernetes configuration path, only used when development mode enabled")
	c.fs.StringVar(&c.brigadeNamespace, "namespace", "", "kubernetes namespace where brigade is running")

	// Parse flags
	c.fs.Parse(os.Args[1:])
}
