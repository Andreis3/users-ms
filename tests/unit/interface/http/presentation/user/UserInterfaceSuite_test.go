package user_interfaces

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
)

var ctrl *gomock.Controller

func Test_UserInterface_Suite(t *testing.T) {
	ctrl = gomock.NewController(t)
	// fetch the current config
	suiteConfig, reporterConfig := GinkgoConfiguration()
	// adjust it
	suiteConfig.SkipStrings = []string{"NEVER-RUN"}
	reporterConfig.FullTrace = true
	RegisterFailHandler(Fail)
	RunSpecs(t, "INTERFACE :: HTTP :: PRESENTATION :: USER :: CONTROLLER : ROUTER", suiteConfig, reporterConfig)
	//RunSpecs(t, "INTERFACE :: HTTP :: PRESENTATION :: USER :: ROUTER", suiteConfig, reporterConfig)
}
