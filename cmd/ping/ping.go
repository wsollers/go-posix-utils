package cmd

import (
	"fmt"
	"os"
//    "syscall"

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
        fmt.Println("No host provided")
        os.Exit(1)
    }
}


