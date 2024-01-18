package main

import (
	_ "github.com/bhaswanth88/coredns-dblocklist"
	_ "github.com/bhaswanth88/coredns/core/plugin"

	"github.com/bhaswanth88/coredns/core/dnsserver"
	"github.com/bhaswanth88/coredns/coremain"
)

func init() {
	dnsserver.Directives = append(
		[]string{"log", "blocklist"},
		dnsserver.Directives...,
	)
}

func main() {
	coremain.Run()
}
