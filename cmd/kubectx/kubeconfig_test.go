package main

import (
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
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


func Test_parseKubeConfig_openError(t *testing.T) {
	_, err := parseKubeConfig("/non/existing/path")
	if err == nil {
		t.Fatalf("expected error" )
	}
	msg := err.Error()
	expectedErrorMessage := `file open error`
	if !strings.Contains(msg, expectedErrorMessage) {
		t.Fatalf("expected error=%q, got=%q", expectedErrorMessage, msg)
	}
}

func Test_parseKubeConfig_yamlFormatError(t *testing.T) {
	file, cleanup := testFile(t, `a: [1, 2`)
	defer cleanup()

	_, err := parseKubeConfig(file)
	if err == nil {
		t.Fatalf("expected error" )
	}
	msg := err.Error()
	expectedErrorMessage := `yaml parse error`
	if !strings.Contains(msg, expectedErrorMessage) {
		t.Fatalf("expected error=%q, got=%q", expectedErrorMessage, msg)
	}
}

func Test_parseKubeConfig_valid_yaml(t *testing.T) {
	file, cleanup := testFile(t, `
apiVersion: v1
current-context: foo
contexts:
- name: c1
- name: c2
- name: c3
`)
	defer cleanup()

	got, err := parseKubeConfig(file)
	if err != nil {
		t.Fatal(err)
	}

	expected := kubeconfig{
		APIVersion:     "v1",
		CurrentContext: "foo",
		Contexts: []Context{
			{Name: "c1"},
			{Name: "c2"},
			{Name: "c3"},
		},
	}

	diff := cmp.Diff(expected, got)
	if diff != "" {
		t.Fatalf("got wrong object:\n%s", diff)
	}

}

func testFile(t *testing.T, contents string) (path string, cleanup func()) {
	t.Helper()

	file, err := ioutil.TempFile(os.TempDir(), "test-file")
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	if _, err = file.Write([]byte(contents)); err != nil {
		t.Fatalf("failed to write to test file: %v", err)
	}

	return file.Name(), func () {
		file.Close()
		os.Remove(file.Name())
	}

}