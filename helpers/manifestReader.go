package helpers

import (

	"code.cloudfoundry.org/cli/util/manifest"
)

type ManifestReader interface {
	Read(string) (manifest.Application, error)
}

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
