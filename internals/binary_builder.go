package internals

import (
	"log"
	"os"
	"os/exec"
)

func BinaryBuild(version string) {
	log.Println("Building the binary...")
	cmds := []string{
		"cd $HOME && rm -rf elys",
		"cd $HOME && git clone https://github.com/elys-network/elys.git",
		"cd $HOME/elys && git pull && git checkout " + version,
		"cd $HOME/elys && make install",
		"mkdir -p $HOME/.elys/cosmovisor/upgrades/" + version + "/bin",
		"cp -a ~/go/bin/elysd ~/.elys/cosmovisor/upgrades/" + version + "/bin/elysd",
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
