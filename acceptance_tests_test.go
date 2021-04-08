package acceptance_cli_test

import (
	"github.com/fabric8-analytics/acceptance_tests/tests"
	. "github.com/onsi/ginkgo"
)


var _ = Describe("AcceptanceTests", func() {

	Describe("PR ACCEPTANCE TESTS", tests.PR_CHECK_SUITE)

	Describe("NIGHTLY SUITE", tests.NIGHTLY_SUITE)
})


