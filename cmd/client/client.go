package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9292")
	if err != nil {
		fmt.Println("error client dial:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter message: ")

	// Error2
	// time.Sleep(5 * time.Second)

	message, _ := reader.ReadString('\n')

	// Error2
	// time.Sleep(5 * time.Second)

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("error write conn:", err)
		return
	}

	// err = conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	// if err != nil {
	// 	fmt.Println("error set deadline:", err)
	// 	return
	// }

	buf := make([]byte, 1024)
	// time.Sleep(7 * time.Second)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Pesan dari server:", string(buf[:n]))
}
