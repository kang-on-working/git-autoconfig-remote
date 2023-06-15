package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecCmdlnStr(Cmdln ...string) error {
	var CmdlnStr string = ""
	for _, str := range Cmdln {
		CmdlnStr += str
	}
	// fmt.Println("Execute this command: ", CmdlnStr)
	var Split []string = strings.Split(CmdlnStr, " ")
	var Prog string = Split[0]
	var Params []string = Split[1:]

	cmd := exec.Command(Prog, Params...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		// fmt.Println(err)
		return err
	}
	return nil
}

func AddRemote(git GIT) error {
	// process the strings
	split := strings.Split(git.Email, "@")
	id := split[0]
	id = strings.ReplaceAll(id, ".", "%2E")
	domain := split[1]
	Email := id + "%40" + domain
	cmd := fmt.Sprintf("git remote add origin https://%s:%s@github.com/%s/%s", Email, git.Token, git.Username, git.Repo)

	err := ExecCmdlnStr(cmd)
	if err != nil {
		cmd = RemoveRemote()
		err = ExecCmdlnStr(cmd)
		if err != nil {
			fmt.Println("Failed ...")
			return err
		}
	}

	return nil
}

func RemoveRemote() string {
	return "git remote remove origin"
}
