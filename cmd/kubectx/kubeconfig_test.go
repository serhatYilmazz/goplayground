package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_kubeconfigPath_homePath(t *testing.T) {
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", "/foo/bar")
	defer os.Setenv("HOME", originalHome)

	got, err := kubeconfigPath()
	if err != nil {
		t.Fatal(err)
	}

	expected := filepath.Join(filepath.FromSlash("/foo/bar"), ".kube", "config")

	if got != expected {
		t.Fatalf("wrong value: expected=%s got=%s", expected, got)
	}
}

func Test_kubeconfigPath_userProfile(t *testing.T) {
	originalHome := os.Getenv("HOME")
	os.Unsetenv("HOME")
	os.Setenv("USERPROFILE", "/foo/bar")
	defer os.Setenv("HOME", originalHome)

	got, err := kubeconfigPath()
	if err != nil {
		t.Fatal(err)
	}

	expected := filepath.Join(filepath.FromSlash("/foo/bar"), ".kube", "config")

	if got != expected {
		t.Fatalf("wrong value: expected=%s got=%s", expected, got)
	}
}

func Test_kubeconfigPath_noEnvSet(t *testing.T) {
	originalHome := os.Getenv("HOME")
	originalHomeUserProfile := os.Getenv("USERPROFILE")

	os.Unsetenv("HOME")
	os.Unsetenv("USERPROFILE")

	defer os.Setenv("HOME", originalHome)
	defer os.Setenv("USERPROFILE", originalHomeUserProfile)

	_, err := kubeconfigPath()
	if err == nil {
		t.Fatalf("Expected error")
	}
}

func Test_kubeconfigPath_envOverride(t *testing.T) {
	originalKubeConfig := os.Getenv("KUBECONFIG")
	os.Setenv("KUBECONFIG", "/foo")

	defer os.Setenv("KUBECONFIG", originalKubeConfig)

	v, err := kubeconfigPath()
	if err != nil {
		t.Fatal(err)
	}
	if expected := "foo"; v != "/foo" {
		t.Fatalf("expected=%s got=%s", expected, v)
	}
}

func Test_kubeconfigPath_doesNotSupportPathSeparator(t *testing.T) {
	originalKubeConfig := os.Getenv("KUBECONFIG")
	path := strings.Join([]string{"file1", "file2"}, string(os.PathListSeparator))
	os.Setenv("KUBECONFIG", path)

	defer os.Setenv("KUBECONFIG", originalKubeConfig)
	_, err := kubeconfigPath()
	if err == nil {
		t.Fatalf("error is expected")
	}
}