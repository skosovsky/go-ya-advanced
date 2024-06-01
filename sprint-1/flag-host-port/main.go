package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Value interface {
	String() string
	Set(string) error
}

type NetAddress struct {
	Host string
	Port int
}

func (n *NetAddress) String() string {
	return n.Host + ":" + strconv.Itoa(n.Port)
}

func (n *NetAddress) Set(flagValue string) error {
	flagValues := strings.Split(flagValue, ":")

	if len(flagValues) != 2 { //nolint:mnd //example
		return fmt.Errorf("invalid net address format: %s", flagValue) //nolint:err113 // example
	}

	port, err := strconv.Atoi(flagValues[1])
	if err != nil {
		return fmt.Errorf("invalid net address format: %s", flagValue) //nolint:err113 // example
	}

	n.Host = flagValues[0]
	n.Port = port

	return nil
}

func main() {
	addr := new(NetAddress)

	// _ = flag.Value(addr) // check interface

	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()

	log.Println(addr.Host, addr.Port)
}
