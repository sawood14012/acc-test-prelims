package tests

import (
	. "github.com/onsi/ginkgo"
)
var (
	Path string = "/data"
	pwd string
	err error
	manifests string = "/manifests"
	file string = "/package.json"
	target string = "/package.json"
)


func TestCRDA_analyse_npm() {
	It("Get Absolute Path", GetAbsPath)
	It("Copy Manifest with vulnerablities", Copy_manifests)
	It("Install Dependencies to run Stack analyses", Install_npm_deps)
	It("Validate analyse for npm ecosystem", Validate_analse_npm)
	It("Cleanup", Cleanup_npm)
}

func TestCRDA_analyse_npm_json() {
	It("Get Absolute Path", GetAbsPath)
	It("Copy Manifest with vulnerablities", Copy_manifests)
	It("Install Dependencies to run Stack analyses", Install_npm_deps)
	It("Validate analyse for npm ecosystem with vulns", Validate_analse_npm_json_vulns)
	It("Cleanup", Cleanup_npm)
}

func TestCRDA_analyse_npm_json_no_vulns() {
	BeforeEach(func(){
		file = "/vulns.json"
	})
	It("Get Absolute Path", GetAbsPath)
	It("Copy Manifest without vulnerablities", Copy_manifests)
	It("Install Dependencies to run Stack analyses", Install_npm_deps)
	It("Validate analyse for npm ecosystem without vulns", Validate_analse_npm_json_no_vulns)
	It("Cleanup", Cleanup_npm)
}

func TestCRDA_analyse_npm_json_verbose() {
	BeforeEach(func(){
		file = "/package.json"
	})
	It("Get Absolute Path", GetAbsPath)
	It("Copy Manifest with vulnerablities", Copy_manifests)
	It("Install Dependencies to run Stack analyses", Install_npm_deps)
	It("Validate analyse for npm ecosystem with vulns json and verbose", Validate_analse_npm_json__vuln_verbose)
	It("Cleanup", Cleanup_npm)
}

func TestCRDA_analyse_npm_debug() {
	It("Get Absolute Path", GetAbsPath)
	It("Copy Manifest with vulnerablities", Copy_manifests)
	It("Install Dependencies to run Stack analyses", Install_npm_deps)
	It("Validate analyse for npm ecosystem with vulns with debug", Validate_analse_npm_vuln_debug)
	It("Cleanup", Cleanup_npm)
}
func TestCRDA_analyse_npm_all_flags() {
	It("Get Absolute Path", GetAbsPath)
	It("Copy Manifest with vulnerablities", Copy_manifests)
	It("Install Dependencies to run Stack analyses", Install_npm_deps)
	It("Validate analyse for npm ecosystem with vulns and all flags set true", Validate_analse_npm_all_flags)
	It("Cleanup", Cleanup_npm)
}


