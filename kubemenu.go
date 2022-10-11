package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"io/ioutil"
	"moul.io/banner"
	"os"
	"path"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getAllDirs(rootDir string) []string {
	file, err := ioutil.ReadDir(rootDir)
	checkError(err)
	var dirs []string
	for _, f := range file {
		if f.IsDir() {
			dirs = append(dirs, f.Name())
		}
	}
	return dirs
}

func isDirExist(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func choseDir() []string {
	homeDir, homeErr := os.UserHomeDir()
	checkError(homeErr)
	defaultKubePath := path.Join(homeDir, ".kube")
	if isDirExist("kube") {
		dirs := getAllDirs("kube")
		return dirs
	} else if isDirExist(defaultKubePath) {
		dirs := getAllDirs(defaultKubePath)
		return dirs
	} else {
		fmt.Println("Can't find kube config dir")
		return []string{}
	}
}

func removeIfFileExist(path string) {
	if _, err := os.Stat(path); err == nil {
		err := os.Remove(path)
		checkError(err)
	}
	if _, err := os.Lstat(path); err == nil {
		err := os.Remove(path)
		checkError(err)
	}
}

func main() {
	fmt.Println(banner.Inline("kube menu"))
	dirs := choseDir()
	var qs = []*survey.Question{
		{
			Name: "cluster",
			Prompt: &survey.Select{
				Message: "Choose kube cluster:",
				Options: dirs,
				Default: dirs[0],
			},
		},
	}

	answers := struct {
		KubeCluster string `survey:"cluster"`
	}{}

	err := survey.Ask(qs, &answers)
	checkError(err)

	homeDir, homeErr := os.UserHomeDir()
	checkError(homeErr)

	if isDirExist("kube") {
		curPath, err := os.Getwd()
		checkError(err)

		newKubeConfig := path.Join(curPath, "kube", answers.KubeCluster, "config")

		oldKubeConfig := path.Join(homeDir, ".kube", "config")
		removeIfFileExist(oldKubeConfig)

		err = os.Symlink(newKubeConfig, path.Join(homeDir, ".kube", "config"))
		checkError(err)
	} else {
		newKubeConfig := path.Join(homeDir, ".kube", answers.KubeCluster, "config")

		oldKubeConfig := path.Join(homeDir, ".kube", "config")
		removeIfFileExist(oldKubeConfig)

		err = os.Symlink(newKubeConfig, path.Join(homeDir, ".kube", "config"))
		checkError(err)
	}
}
