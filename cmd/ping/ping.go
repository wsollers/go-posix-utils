package cmd

import (
	"fmt"
	//"os"
  "log"
  //"net"
  //"golang.org/x/net/icmp.Echo"
  //"golang.org/x/net/ipv4"
	"golang.org/x/net/icmp"
	//"golang.org/x/net/ipv4"
	//"golang.org/x/net/ipv6"
  

//"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

)

// https://datatracker.ietf.org/doc/html/rfc792
// /usr/sbin/ping host [timeout]
// /usr/sbin/ping  -s [-l | -U]  [-adlLnrRv] [-A addr_family] 
//                 [-c traffic_class] [-g gateway  [ -g	gateway...]]
//                 [-F flow_label] [-I interval]

var PingCmd = &cobra.Command{
    Use:   "ping",
    Short: "Ping the server",
    Long:  "Ping the server",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("ping called")
        ping()
    },
}

func ping() {
  fmt.Println("Pinging server...")
  host := viper.GetString("host")
  fmt.Println("host:", host)
  if host == "" {
    log.Fatal("No host provided")
  }
  timeout := viper.GetInt("timeout")
  if timeout < 0 {
    log.Fatal("Invalid timeout")
  }
  pkt := icmp.Echo {
    ID: 1,
    Seq: 1,
    Data: []byte("Hello"),
  }

  fmt.Println("pkt:", pkt)
  // we have a host and a timeout
  // build a raw icmp ping packet and send it to the specified host

}


