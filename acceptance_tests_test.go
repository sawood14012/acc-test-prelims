package acceptance_cli_test

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	acc_helper_cli "github.com/fabric8-analytics/acceptance_tests/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
	Path string
	pwd string
)

var _ = Describe("AcceptanceTests", func() {
    Path = "/data"
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)


	Describe("Run Crda version", func() {
        It("Runs and Validate CLI version", func(){
			cmd := exec.Command("crda", "version")
			_, err := cmd.Output()
			//fmt.Println(string(stdout))
			Expect(err).NotTo(HaveOccurred())
			
		})
	})
	Describe("Run Crda help", func() {
        It("Runs and Validate Help command", func(){
			cmd := exec.Command("crda", "help")
			_, err := cmd.Output()
			//fmt.Println(string(stdout))
			Expect(err).NotTo(HaveOccurred())
			
		})
	})
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
			pwd, err = acc_helper_cli.Get_abs_path(Path)
			if err != nil {
				ErrorLogger.Println(err)
			}
		})
		It("Install Dependencies to run Stack analyses", func(){
			cmd := exec.Command("/bin/sh", "-c", "cd "+pwd + "; npm i;")
			stdout, err:= cmd.Output()
			InfoLogger.Println(string(stdout))
	        if err != nil {
				ErrorLogger.Println(err)
			}
		})
        It("Validate analyse for npm ecosystem", func(){
			cmd := exec.Command("crda", "analyse", pwd+ "/package.json")
			stdout, err := cmd.Output()
			InfoLogger.Println(string(stdout))
			Expect(err).NotTo(HaveOccurred())
			
		})
		It("Cleanup", func(){
			err := acc_helper_cli.Cleanup_npm(pwd);
			if err != nil{
				ErrorLogger.Println(err)
			}
		})
	})
})

