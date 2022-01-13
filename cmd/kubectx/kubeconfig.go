package main

import (
	"errors"
	"os"
	"path/filepath"
)

func  kubeconfigPath() (string, error) {
	// KUBECONFIG env variable
	if v := os.Getenv("KUBECONFIG"); v != "" {
		// TODO KUBECONFIG=file1:file2 currently not supported.
		list := filepath.SplitList(v)
		if len(list) > 1 {
			return "", errors.New("multiple files in KUBECONFIG is not supported")
		}
		return v, nil
	}

	// return default path
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE") // windows
	}
	if home == "" {
		return "", errors.New("HOME or USERPROFILE environment variable not set")
	}
	return filepath.Join(home, ".kube", "config"), nil
}
