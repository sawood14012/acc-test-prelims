package helper_cli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Get_abs_path(filepath string)(string, error){
	if filepath == "" {
        return "", errors.New("empty file path supplied")
    }
	cmd := exec.Command("pwd")
	stdout, err := cmd.Output()
	if err != nil {
		return "", errors.New("cannot run PWD")
	}
	pwd := string(stdout)
	pwd = strings.TrimSpace(pwd)
    result := pwd + filepath;
	return result, nil

}

func Cleanup_npm(filepath string)(error){
	fmt.Println(filepath)
	err := os.RemoveAll(filepath+ "/node_modules/")
	err1 := os.Remove(filepath + "/package-lock.json")
    if err != nil || err1 != nil {
        return err
    }
	return nil
}

