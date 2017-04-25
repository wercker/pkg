package kubernetes

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/wercker/pkg/log"
)

// Metadata holds all the data exposed by the Downward API
type Metadata struct {
	Root      string // root dir, defaults to /metadata
	name      string
	namespace string
	labels    map[string]string
}

// LoadAll loads all of the metadata we can find
func (m *Metadata) LoadAll() {
	if labels, err := m.Labels(); err == nil {
		m.labels = labels
	}
	if name, err := m.Name(); err == nil {
		m.name = name
	}
	if namespace, err := m.Namespace(); err == nil {
		m.namespace = namespace
	}
}

// Fields gives back our metadata as log.Fields
func (m *Metadata) Fields() log.Fields {
	fields := log.Fields{}
	for k, v := range m.labels {
		fields[fmt.Sprintf("kubernetes_labels_%s", k)] = v
	}
	if m.name != "" {
		fields["kubernetes_pod_name"] = m.name
	}

	if m.namespace != "" {
		fields["kubernetes_namespace_name"] = m.namespace
	}
	return fields
}

func (m *Metadata) readString(path string) (string, error) {
	if m.Root == "" {
		m.Root = "/metadata"
	}
	f, err := os.Open(filepath.Join(m.Root, path))
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Labels looks like
/*
app="kiddie-pool"
branch="logging"
commit="64f78e45dbe3eb886f0f24da5447f0dedaaaf665"
pod-template-hash="1222321528"
*/
func (m *Metadata) Labels() (map[string]string, error) {
	labelMatcher := regexp.MustCompile(`(\S+)="(\S+)"`)
	labels := map[string]string{}

	s, err := m.readString("labels")
	if err != nil {
		return labels, err
	}

	for _, match := range labelMatcher.FindAllStringSubmatch(s, -1) {
		labels[match[1]] = match[2]
	}
	return labels, nil
}

func (m *Metadata) Name() (string, error) {
	return m.readString("name")
}

func (m *Metadata) Namespace() (string, error) {
	return m.readString("namespace")
}
