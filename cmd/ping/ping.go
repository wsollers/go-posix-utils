package cmd

import (
	"fmt"
  "strconv"
  "log"
	"os"
  "net"
  //"net/netip"
	//"runtime"
	"golang.org/x/net/icmp"
  ///"net"
  //"golang.org/x/net/icmp.Echo"
	"golang.org/x/net/ipv4"
	//"golang.org/x/net/ipv6"
  

//"syscall"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
  //"golang.org/x/sys/unix"
)

// https://datatracker.ietf.org/doc/html/rfc792
// /usr/sbin/ping host [timeout]
// /usr/sbin/ping  -s [-l | -U]  [-adlLnrRv] [-A addr_family] 
//                 [-c traffic_class] [-g gateway  [ -g	gateway...]]
//                 [-F flow_label] [-I interval]

var host string
var timeout int

var PingCmd = &cobra.Command{
  Use:   "ping",
  Short: "Ping the server",
  Long:  "Ping the server",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("ping called")

    if len(args) < 1 {
      log.Fatal("usage: ping host [timeout]")
    }
    host = args[0]
    if len(host) < 1 {
      log.Fatal("No host provided")
    }
    err := error(nil)
    timeout, err = strconv.Atoi(args[1])
    if err != nil {
      log.Fatal("Invalid timeout")
    }
    if timeout < 1 {
      log.Fatal("Invalid timeout")
    }
    fmt.Println("host:", host)
    if host == "" {
      log.Fatal("No host provided")
    }
  
    ping()
  },
}

func ping() {
  fmt.Println("Pinging server...")
  c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
  if err != nil {
    log.Fatalf("listen err, %s", err)
  }
  defer c.Close()

  ip, err := net.ResolveIPAddr("ip", host)
  if err != nil {
    log.Fatal("ResolveIPAddr: ",err)
  }
  wm := icmp.Message{
    Type: ipv4.ICMPTypeEcho, Code: 0,
    Body: &icmp.Echo{
     ID: os.Getpid() & 0xffff, Seq: 1,
     Data: []byte("HELLO-R-U-THERE"),
    },
  }
  wb, err := wm.Marshal(nil)
  if err != nil {
    log.Fatal(err)
  }
  if _, err := c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(ip.IP.String())}); err != nil {
    log.Fatalf("WriteTo err, %s", err)
  }

  rb := make([]byte, 1500)
  n, peer, err := c.ReadFrom(rb)
  if err != nil {
    log.Fatal(err)
    }
    rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
    if err != nil {
        log.Fatal(err)
    }
    switch rm.Type {
    case ipv4.ICMPTypeEchoReply:
        log.Printf("got reflection from %v", peer)
      fmt.Println();
      fmt.Println("Packet received from: ", rm)
    default:
        log.Printf("got %+v; want echo reply", rm)
    }
/*
  c, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

 	hostname, err := os.Hostname()
  if err != nil {
    log.Fatal("Hostname: ", err)
  }
  ip, err := net.ResolveIPAddr("ip", host)
  if err != nil {
    fmt.Println("ResolveIPAddr: ",err)
  }
	fmt.Println("Host: ", host)
  fmt.Println("ip: ", ip)
	fmt.Println("Hostname: ", hostname)

  serverFD, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, unix.IPPROTO_ICMP)
	if err != nil {
		log.Fatal("Socket: ", err)
	}
  defer unix.Close(serverFD)
*/
  /*
  */
  /*
  addrRaw1 := net.ParseIP(host)
  addrRaw2 := addrRaw1.To4()
  var byteArray [4]byte
  copy (byteArray[:], addrRaw2.To4())
  serverAddr := &unix.SockaddrInet4{
		Port: 0,
		Addr: byteArray,
	}
  /*
  err = unix.Bind(serverFD, serverAddr)
	if err != nil {
		log.Fatal("Bind: ", err)
	}
	fmt.Printf("Server: Bound to addr: %d, port: %d\n", serverAddr.Addr, serverAddr.Port)
  */
  /*
  err = unix.Connect(serverFD, serverAddr)
	if err != nil {
		if err == unix.ECONNREFUSED {
      log.Fatal("* Connection failed: ", err)
		}
	}
  fmt.Printf("Server: Connected to addr: %d, port: %d\n", serverAddr.Addr, serverAddr.Port)
  /*
  server, err := net.ResolveUDPAddr("udp4", ip.IP.String())
  if err != nil {
    log.Fatal("ResolveUDPAddr: ", err)
  }
  conn, err := net.DialUDP("udp4:icmp", nil, server)
  if err != nil {
    fmt.Println("Dial Error: ", err)
    return
  }
  defer conn.Close()
	wm := icmp.Message {
		Type: ipv4.ICMPTypeEcho,
    Code: 0,
		Body: &icmp.Echo {
		  ID: os.Getpid() & 0xffff, 
      Seq: 1,
		  Data: []byte("HELLO-R-U-THERE"),
		},
	}

	wb, err := wm.Marshal(nil)
	if err != nil {
    log.Fatal("Marshal: ", err)
	}
  if _, err := conn.Write(wb); err != nil {
    log.Fatal("Write Error: ", err)
	}

	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	if err != nil {
    log.Fatal("ReadFrom Error: ", err)
	}
	rm, err := icmp.ParseMessage(58, rb[:n])
	if err != nil {
    log.Fatal("Parse Message: ", err)
	}
	switch rm.Type {
	  case ipv6.ICMPTypeEchoReply:
      log.Printf("IPV6: got reflection from %v", peer)
    case ipv4.ICMPTypeEchoReply:
      log.Printf("IPV4: got reflection from %v", peer)
	  default:
	  	log.Printf("got %+v; want echo reply", rm)
	}
*/















  /*
  pkt := icmp.Echo {
    ID: 1,
    Seq: 1,
    Data: []byte("Hello"),
  }
  fmt.Println("pkt:", pkt)

  for {
    rawpkt, err := pkt.Marshal(3)
    if err != nil {
      Log.Fatal("Error marshalling packet")
    }

    // send packet
    // wait for response
    // print response
  }
  // we have a host and a timeout
  // build a raw icmp ping packet and send it to the specified host
  */
}


