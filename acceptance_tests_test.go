package acceptance_cli_test

import (
	"fmt"
	"os/exec"
	"github.com/fabric8-analytics/acceptance_tests/helper"
	"github.com/fabric8-analytics/acceptance_tests/tests"
	"github.com/fabric8-analytics/acceptance_tests/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	Path string
	pwd string
	err error
)

var _ = Describe("AcceptanceTests", func() {
    Path = "/data"
	Describe("Run Crda version", tests.TestCRDA_version)
	Describe("Run Crda help", tests.TestCRDA_help)

	Describe("Run Crda analyse", func() {
        It("Validate analyse without flags throws error", func(){
			cmd := exec.Command("crda", "analyse")
			stdout, err := cmd.Output()
			fmt.Println(string(stdout))
			Expect(err).To(HaveOccurred())
			
		})
	})
	Describe("Run Crda analyse", func() {
		It("Get Absolute Path", func(){
			pwd, err = helper_cli.Get_abs_path(Path)
			if err != nil {
				acc_log.ErrorLogger.Println(err)
			}
		})
		It("Install Dependencies to run Stack analyses", func(){
			cmd := exec.Command("/bin/sh", "-c", "cd "+pwd + "; npm i;")
			stdout, err:= cmd.Output()
			acc_log.InfoLogger.Println(string(stdout))
	        if err != nil {
				acc_log.ErrorLogger.Println(err)
			}
		})
        It("Validate analyse for npm ecosystem", func(){
			cmd := exec.Command("crda", "analyse", pwd+ "/package.json")
			stdout, err := cmd.Output()
			acc_log.InfoLogger.Println(string(stdout))
			Expect(err).NotTo(HaveOccurred())
			
		})
		It("Cleanup", func(){
			err := helper_cli.Cleanup_npm(pwd);
			if err != nil{
				acc_log.ErrorLogger.Println(err)
			}
		})
	})
})

