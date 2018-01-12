package helpers

import (
	"code.cloudfoundry.org/cli/util/manifest"
)

type ManifestReader interface {
	Read(string) (manifest.Application, error)
}

// Impl
//
type manifestReader struct{}

func (manifestReader) Read(path string) (manifest.Application, error) {
	applications, err := manifest.ReadAndMergeManifests(path)
	if err != nil {
		return manifest.Application{}, err
	}

	return applications[0], err
}

func NewManifestReader() manifestReader {
	return manifestReader{}
}

// Mock
//
type mockManifestReader struct {
	application manifest.Application
	error       error
}

func (m *mockManifestReader) SetApplication(application manifest.Application, error error) {
	m.application = application
	m.error = error
}

func (m mockManifestReader) Read(path string) (manifest.Application, error) {
	return m.application, m.error
}

func NewMockManifestReader() mockManifestReader {
	return mockManifestReader{
		application: manifest.Application{},
		error:       nil,
	}
}
