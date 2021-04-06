package acceptance_cli_test

import (
	"testing"
	"github.com/fabric8-analytics/acceptance_tests/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


func TestAcceptanceTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AcceptanceTests Suite")
}

var _ = BeforeSuite(func() {
	acc_log.Init_log()
})