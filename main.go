package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"

	ping "github.com/wsollers/go-posix-utils/cmd/ping"
	web "github.com/wsollers/go-posix-utils/cmd/web"
)

var rootCmd = &cobra.Command{
	Use:   "go-posix-utils [sub]",
	Short: "golang reimplementation of some posix utilities",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd Run with args: %v\n", args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	},
}

func main() {
	rootCmd.AddCommand(ping.PingCmd)
	rootCmd.AddCommand(web.WebServCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	//  ping.PingCmd.Run(nil, nil)
	/*
	  var subCmd = &cobra.Command{
	    Use:   "sub [no options!]",
	    Short: "My subcommand",
	    PreRun: func(cmd *cobra.Command, args []string) {
	      fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
	    },
	    Run: func(cmd *cobra.Command, args []string) {
	      fmt.Printf("Inside subCmd Run with args: %v\n", args)
	    },
	    PostRun: func(cmd *cobra.Command, args []string) {
	      fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
	    },
	    PersistentPostRun: func(cmd *cobra.Command, args []string) {
	      fmt.Printf("Inside subCmd PersistentPostRun with args: %v\n", args)
	    },
	  }

	  rootCmd.AddCommand(subCmd)

	  rootCmd.SetArgs([]string{""})
	  rootCmd.Execute()
	  fmt.Println()
	  rootCmd.SetArgs([]string{"sub", "arg1", "arg2"})
	  rootCmd.Execute()
	*/
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
