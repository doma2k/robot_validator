package internals

import (
	"log"
	"os/exec"
)

func BinaryBuild(version string) {
	log.Println("Building the binary")
	cmds := []string{
		"cd $HOME && rm -r elys",
		"git clone https://github.com/elys-network/elys.git",
		"cd elys",
		"git pull",
		"git checkout " + version,
		"make install",
		"cp -a ~/go/bin/elysd ~/.elys/cosmovisor/upgrades/" + version + "/bin/elysd",
	}

	for _, cmdStr := range cmds {
		cmd := exec.Command("/bin/sh", "-c", cmdStr)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Command failed with error: %v", err)
		}
	}
	log.Println("Binary built successfully")
}
