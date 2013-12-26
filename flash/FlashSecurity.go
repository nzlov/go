package flash

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type FlashSecurity struct {
}

func (f *FlashSecurity) start() {
	go func() {
		xml := "<cross-domain-policy> <allow-access-from domain=\"*\" to-ports=\"*\"/></cross-domain-policy> \n"

		lis, err := net.Listen("tcp", ":843")

		if err != nil {
			fmt.Println("Error when listen: ", ":843")
			return
		}
		defer lis.Close()
		for {
			conn, err := lis.Accept()
			if err != nil {
				fmt.Println("Error accepting client: ", err.Error())
				os.Exit(0)
			}

			go func(con net.Conn) {
				defer con.Close()
				fmt.Println("New connection: ", con.RemoteAddr())
				r := bufio.NewReader(con)

				data, err := r.ReadBytes('>')
				if err != nil {
					fmt.Printf("Client %v quit.\n", con.RemoteAddr())
					con.Close()
					return
				}

				fmt.Printf("%s said: %s\n", con.RemoteAddr(), data)
				con.Write([]byte(xml))
			}(conn)
		}

	}()
}
