package main

import (
	"fmt"
	"pkg"
)

func main() {
	key := pkg.GetStrKey()

	jsonFile := "./git.json"
	jsonData, err := pkg.JsonParse(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	git, err := pkg.DecryptJsonValue(jsonData, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	pkg.ExecCmdlnStr("git config user.email \"", git.Email, "\"")
	pkg.ExecCmdlnStr("git config user.name \"", git.Username, "\"")
	pkg.AddRemote(git)
}
