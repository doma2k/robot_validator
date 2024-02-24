package internals

import (
	"log"
	"os"
	"os/exec"
)

var binary = "elys"
var github_repo = "https://github.com/elys-network/elys.git"

func BinaryBuild(version string) {
	log.Println("Building the binary...")
	cmds := []string{
		"cd $HOME && rm -rf " + binary,
		"cd $HOME && git clone " + github_repo,
		"cd $HOME/" + binary + " && git pull && git checkout " + version,
		"cd $HOME/" + binary + " && make install",
		"mkdir -p $HOME/." + binary +"/cosmovisor/upgrades/" + version + "/bin",
		"cp -a ~/go/bin/" + binary + "d ~/." + binary + "/cosmovisor/upgrades/" + version + "/bin/" + binary + "d",
	}

	for _, cmdStr := range cmds {
		cmd := exec.Command("/bin/sh", "-c", cmdStr)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			log.Fatalf("Command failed to start: %v", cmdStr)
		}
		err = cmd.Wait()
		if err != nil {
			log.Fatalf("Command failed to complete: %v", cmdStr)
		}
	}
	log.Println("Binary built successfully")
}
