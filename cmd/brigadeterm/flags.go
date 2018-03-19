package main

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

type cmdFlags struct {
	kubeConfig       string
	brigadeNamespace string
}

func newCmdFlags() *cmdFlags {
	fls := &cmdFlags{}
	fls.init()
	return fls
}
func (c *cmdFlags) init() {
	kubehome := filepath.Join(homedir.HomeDir(), ".kube", "config")

	// register flags
	flag.StringVar(&c.kubeConfig, "kubeconfig", kubehome, "kubernetes configuration path, only used when development mode enabled")
	flag.StringVar(&c.brigadeNamespace, "namespace", "", "kubernetes namespace where brigade is running")

	// Parse flags
	flag.Parse()
}
