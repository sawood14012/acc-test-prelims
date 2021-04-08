package acceptance_cli_test

import (
	"github.com/fabric8-analytics/acceptance_tests/tests"
	. "github.com/onsi/ginkgo"
)


var _ = Describe("AcceptanceTests", func() {
	Describe("Test CRDA auth command", tests.TestCRDA_auth)
	Describe("Test CRDA analyse Go", tests.GolangTestSuite)
	Describe("Test for invalid path throws error", tests.TestInvalidPath)
	Describe("Test for invalid command", tests.TestInvalidCommand)
	Describe("Test for invalid flag throws error", tests.TestInvalidFlag)
	Describe("Run Crda version", tests.TestCRDA_version)
	Describe("Run Crda help", tests.TestCRDA_help)
	Describe("Validate CRDA completion command by os", tests.TestCRDA_completion)
	Describe("Validate there is a help page for all commands", tests.TestCRDA_all_commands_help)
	Describe("Run Crda analyse without any file", tests.TestCRDA_analyse_without_file) 
    Describe("Test CRDA analyse Maven", tests.MavenTestSuite)
	Describe("Test CRDA analyse Npm", tests.NpmTestSuite)
	Describe("Test CRDA analyse Pypi", tests.PypiTestSuite)
})

