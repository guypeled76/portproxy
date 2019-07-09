package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/guypeled76/portproxy/proxy"
	"github.com/magiconair/properties"
	"log"
	"sync"
)

func main() {

	var args struct {
		Local  string `arg:"positional"`
		Remote string `arg:"positional"`
	}

	arg.MustParse(&args)

	log.Printf("Initializing port proxy based on %s to %s. \n", args.Local, args.Remote)

	local := properties.MustLoadFile(args.Local, properties.UTF8).Map()
	remote := properties.MustLoadFile(args.Remote, properties.UTF8).Map()

	for name, fromPort := range local {
		toHost, hasHost := remote[name]
		if hasHost {
			go proxy.CreatePortProxy(fmt.Sprintf("localhost:%s", fromPort), toHost)
		} else {
			log.Printf("Could not find mapping for %s", name)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}
