package gen

import "os"

func RecreateDir(dirname string) {
	if dirExists(dirname) {
		deleteDir(dirname)
	}
	createDir(dirname)
}

func dirExists(dirname string) bool {
	_, err := os.Stat(dirname)
	return !os.IsNotExist(err)
}

func deleteDir(dirname string) {
	if err := os.RemoveAll(dirname); err != nil {
		panic(err.Error())
	}
}

func createDir(dirname string) {
	if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
		panic(err.Error())
	}
}
