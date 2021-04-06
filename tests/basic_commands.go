package tests

import (
	"os/exec"
	"github.com/fabric8-analytics/acceptance_tests/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)



func TestCRDA_version() {

	It("Runs and Validate CLI version", func(){
		cmd := exec.Command("crda", "version")
		stdout, err := cmd.Output()
		acc_log.InfoLogger.Println(string(stdout))
		Expect(err).NotTo(HaveOccurred())
		
	})

}

func TestCRDA_help() {
	It("Runs and Validate Help command", func(){
		cmd := exec.Command("crda", "help")
		stdout, err := cmd.Output()
		acc_log.InfoLogger.Println(string(stdout))
		Expect(err).NotTo(HaveOccurred())
		
	})

}