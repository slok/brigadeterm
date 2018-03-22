package main

import (
	"fmt"
	"os"

	"github.com/Azure/brigade/pkg/storage/kube"
	"github.com/rivo/tview"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/slok/brigadeterm/pkg/controller"
	"github.com/slok/brigadeterm/pkg/service/brigade"
	"github.com/slok/brigadeterm/pkg/ui"
)

// Main is the main package.
type Main struct {
	flags *cmdFlags
}

// NewMain returns a new main application.
func NewMain() *Main {
	return &Main{
		flags: newCmdFlags(),
	}
}

// Run will run the main application.
func (m *Main) Run() error {
	// Create external dependencies.
	k8scli, err := m.createKubernetesClients()
	if err != nil {
		return err
	}
	brigadek8s := kube.New(k8scli, m.flags.brigadeNamespace)

	// Create services
	brigadeService := brigade.NewService(brigadek8s)

	// Create controller.
	uictrl := controller.NewController(brigadeService)

	// Create the terminal app.
	app := tview.NewApplication()

	index := ui.NewIndex(uictrl, app)

	return index.Render()
}

func (m *Main) createKubernetesClients() (kubernetes.Interface, error) {
	config, err := m.loadKubernetesConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

// loadKubernetesConfig loads kubernetes configuration based on flags.
func (m *Main) loadKubernetesConfig() (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", m.flags.kubeConfig)
}

func main() {
	m := NewMain()
	if err := m.Run(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
