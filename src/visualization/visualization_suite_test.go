package visualization_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVisualization(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Visualization Suite")
}
