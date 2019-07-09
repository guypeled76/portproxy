package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/guypeled76/portproxy/proxy"
	"github.com/magiconair/properties"
	"log"
	"os"
	"sync"
)

func main() {

	var args struct {
		Local  string `arg:"positional"`
		Remote string `arg:"positional"`
	}

	arg.MustParse(&args)

	if _, err := os.Stat(args.Local); os.IsNotExist(err) {
		log.Fatalf("Local properties file '%s' was not found.", args.Local)
	}

	if _, err := os.Stat(args.Remote); os.IsNotExist(err) {
		log.Fatalf("Remote properties file '%s' was not found.", args.Remote)
	}

	log.Printf("Initializing port proxy based on %s to %s. \n", args.Local, args.Remote)

	local := properties.MustLoadFile(args.Local, properties.UTF8).Map()
	remote := properties.MustLoadFile(args.Remote, properties.UTF8).Map()

	for proxyName, fromPort := range local {
		toHost, hasHost := remote[proxyName]
		if hasHost {
			go proxy.CreatePortProxy(proxyName, fmt.Sprintf("localhost:%s", fromPort), toHost)
		} else {
			log.Printf("Could not find mapping for %s", proxyName)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}
