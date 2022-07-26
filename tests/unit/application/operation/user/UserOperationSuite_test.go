package operation_user

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
)

var ctrl *gomock.Controller

func TestUser(t *testing.T) {
	ctrl = gomock.NewController(t)
	// fetch the current config
	suiteConfig, reporterConfig := GinkgoConfiguration()
	// adjust it
	suiteConfig.SkipStrings = []string{"NEVER-RUN"}
	reporterConfig.FullTrace = true
	RegisterFailHandler(Fail)
	RunSpecs(t, "OPERATION :: TESTS :: USER :: SUITE", suiteConfig, reporterConfig)
}
