package acceptance_cli_test

import (
	"github.com/fabric8-analytics/acceptance_tests/tests"
	. "github.com/onsi/ginkgo"
)


var _ = Describe("AcceptanceTests", func() {
	Describe("Run Crda version", tests.TestCRDA_version)
	Describe("Run Crda help", tests.TestCRDA_help)
	Describe("Validate there is a help page for all commands", tests.TestCRDA_all_commands_help)
	Describe("Run Crda analyse without any file", tests.TestCRDA_analyse_without_file)
	Context("Run CRDA analyse for npm", tests.TestCRDA_analyse_npm)
	Context("Run CRDA analyses for npm with json" , tests.TestCRDA_analyse_npm_json)
	Context("Run CRDA analyse for npm with no vulnerablities", tests.TestCRDA_analyse_npm_json_no_vulns)
	Context("Run CRDA aanalyse with verbose for npm", tests.TestCRDA_analyse_npm_json_verbose)
	Context("Run CRDA completion command based on os", tests.TestCRDA_completion)
})

