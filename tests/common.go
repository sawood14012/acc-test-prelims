package tests

import (
	"os/exec"
	"github.com/fabric8-analytics/acceptance_tests/helper"
	"github.com/fabric8-analytics/acceptance_tests/log"
	. "github.com/onsi/gomega"
)

func GetAbsPath() {
	pwd, err = helper_cli.Get_abs_path(Path)
		if err != nil {
			acc_log.ErrorLogger.Println(err)
		}
}

func Install_npm_deps() {
	
	cmd := exec.Command("/bin/sh", "-c", "cd "+pwd + "; npm i;")
	stdout, err:= cmd.Output()
	acc_log.InfoLogger.Println(string(stdout))
	if err != nil {
		acc_log.ErrorLogger.Println(err)
	}
}

func Validate_analse_npm() {
	cmd := exec.Command("crda", "analyse", pwd+ "/package.json")
	stdout, err := cmd.Output()
	acc_log.InfoLogger.Println(string(stdout))
	Expect(err).NotTo(HaveOccurred())

}

func Validate_analse_npm_json_vulns() {
	cmd := exec.Command("crda", "analyse", pwd+ "/package.json", "-j")
	stdout, err := cmd.Output()
	acc_log.InfoLogger.Println(string(stdout))
	e := err.(*exec.ExitError);
	acc_log.InfoLogger.Println(e.ExitCode())
	Expect(e.ExitCode()).To(Equal(2))

}

func Validate_analse_npm_json_no_vulns() {
	cmd := exec.Command("crda", "analyse", pwd+ "/package.json", "-j")
	stdout, err := cmd.Output()
	acc_log.InfoLogger.Println(string(stdout))
	Expect(err).NotTo(HaveOccurred())

}

func Validate_analse_npm_json__vuln_verbose() {
	cmd := exec.Command("crda", "analyse", pwd+ "/package.json", "-j", "-v")
	stdout, err := cmd.Output()
	acc_log.InfoLogger.Println(string(stdout))
	e := err.(*exec.ExitError);
	acc_log.InfoLogger.Println(e.ExitCode())
	Expect(e.ExitCode()).To(Equal(2))

}

func Validate_analse_npm_vuln_debug() {
	cmd := exec.Command("crda", "analyse", pwd+ "/package.json", "-d")
	stdout, err := cmd.Output()
	acc_log.InfoLogger.Println(string(stdout))
	e := err.(*exec.ExitError);
	acc_log.InfoLogger.Println(e.ExitCode())
	Expect(e.ExitCode()).To(Equal(2))
}

func Validate_analse_npm_all_flags() {
	cmd := exec.Command("crda", "analyse", pwd+ "/package.json", "-d", "-v", "-j")
	stdout, err := cmd.Output()
	acc_log.InfoLogger.Println(string(stdout))
	e := err.(*exec.ExitError);
	acc_log.InfoLogger.Println(e.ExitCode())
	Expect(e.ExitCode()).To(Equal(2))

}

func Cleanup_npm(){
	err := helper_cli.Cleanup_npm(pwd);
		if err != nil{
			acc_log.ErrorLogger.Println(err)
		}
}

func Copy_manifests(){
	dir1, err := helper_cli.Get_abs_path(manifests)
		if err != nil {
			acc_log.ErrorLogger.Println(err)
		}
	dir2, err := helper_cli.Get_abs_path(Path)
		if err != nil {
			acc_log.ErrorLogger.Println(err)
		}
		acc_log.InfoLogger.Println(dir1+file)
		acc_log.InfoLogger.Println(dir2+target)
	err = helper_cli.Copy_contents_to_target(dir1+file, dir2+target)
	Expect(err).NotTo(HaveOccurred())
}
