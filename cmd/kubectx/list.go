package main

import (
	"facette.io/natsort"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type kubeconfig struct {
	APIVersion     string    `yaml:"apiVersion"`
	CurrentContext string    `yaml:"current-context"`
	Contexts       []Context `yaml:"contexts"`
}

type Context struct {
	Name    string `yaml:"name"`
}

func printListContexts(writer io.Writer) error {
	// TODO extract printing and sorting into a function that is testable
	// parse kubeconfig --> []context_names, current_context_name
	cfgPath, err := kubeconfigPath()
	if err != nil {
		return errors.Wrap(err, "failed to determine kubeconfig path")
	}
	cfg, err := parseKubeConfig(cfgPath)
	if err != nil {
		return errors.Wrap(err, "failed to read kubeconfig file")
	}

	ctxs := make([]string, 0, len(cfg.Contexts))
	for _, c := range cfg.Contexts {
		ctxs = append(ctxs, c.Name)
	}

	natsort.Sort(ctxs)

	for _, c := range ctxs {
		out := c
		if out == cfg.CurrentContext {
			out = color.New(color.Bold, color.FgHiYellow).Sprintf(c)
		}
		fmt.Println(out)
	}

	return nil
}

func parseKubeConfig(path string) (kubeconfig, error) {
	// TODO refactor to accept io.Reader instead of file
	var v kubeconfig

	f, err := os.Open(path)
	if err != nil {
		return v, errors.Wrap(err, "file open error")
	}

	err = yaml.NewDecoder(f).Decode(&v)

	// errors.Wrap returns nil if err is nil
	return v, errors.Wrap(err, "yaml parse error")
}
