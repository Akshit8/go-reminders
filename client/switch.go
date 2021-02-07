package client

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// idsFlag represent []string values passed from CLI
type idsFlag []string

func (list *idsFlag) String() string {
	return strings.Join(*list, ",")
}

func (list *idsFlag) Set(v string) error {
	*list = append(*list, v)
	return nil
}

// Client represents the client for communicating with backend API (here HTTP).
type Client interface {
	Create(title, message string, duration time.Duration) ([]byte, error)
	Edit(id string, title, message string, duration time.Duration) ([]byte, error)
	Fetch(ids []string) ([]byte, error)
	Delete(ids []string) error
	Healthy(host string) bool
}

// Switch represents CLI command switch
type Switch struct {
	client        Client
	backendAPIURI string
	commands      map[string]func() func(string) error
}

// NewSwitch creates new instance of Switch
func NewSwitch(backendAPIURI string) *Switch {
	httpClient := NewHTTPClient(backendAPIURI)
	s := &Switch{
		client:        httpClient,
		backendAPIURI: backendAPIURI,
	}
	s.commands = map[string]func() func(string) error{
		"create": s.create,
		"edit":   s.edit,
		"fetch":  s.fetch,
		"delete": s.delete,
		"health": s.health,
	}
	return s
}

// Switch analyses the CLI args end executes the given command
func (s Switch) Switch() error {
	cmdName := os.Args[1]
	cmd, ok := s.commands[cmdName]
	if !ok {
		return fmt.Errorf("invalid command '%s'", cmdName)
	}
	return cmd()(cmdName)
}

// Help prints a useful message about command usage
func (s Switch) Help() {
	var help string
	for name := range s.commands {
		help += name + "\t --help\n"
	}
	fmt.Printf("usage of %s:\n<command> [<args>]\n%s", os.Args[0], help)
}

func (s Switch) create() func(string) error {
	return nil
}

func (s Switch) edit() func(string) error {
	return nil
}

func (s Switch) fetch() func(string) error {
	return nil
}

func (s Switch) delete() func(string) error {
	return nil
}

func (s Switch) health() func(string) error {
	return nil
}
