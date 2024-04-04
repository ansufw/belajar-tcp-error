package main

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Error2 - option1
	// err := conn.SetReadDeadline(time.Now().Add(3 * time.Second))

	// Error2 - option2 - put deadline less than the time need by server reading the request
	err := conn.SetReadDeadline(time.Now().Add(100 * time.Microsecond))

	if err != nil {
		fmt.Println("error setting read deadline:", err)
		return
	}

	buffer := make([]byte, 1024)
	start := time.Now() // Waktu sebelum eksekusi

	n, err := conn.Read(buffer)

	elapsed := time.Since(start)                   // Hitung durasi
	fmt.Println("Duration of conn.Read:", elapsed) // Cetak durasi

	if err != nil {
		fmt.Println("error reading from connection:", err)
		return
	}

	msg := string(buffer[:n])
	fmt.Println("message from client:", msg)

	// error 1
	// err = conn.SetWriteDeadline(time.Now().Add(200 * time.Microsecond))
	err = conn.SetWriteDeadline(time.Now().Add(200 * time.Microsecond))
	if err != nil {
		fmt.Println("error setting write deadline:", err)
		return
	}

	// error 1
	// time.Sleep(110 * time.Microsecond)

	start = time.Now() // Waktu sebelum eksekusi

	_, err = conn.Write([]byte(response)) // Eksekusi baris kode

	elapsed = time.Since(start)                     // Hitung durasi
	fmt.Println("Duration of conn.Write:", elapsed) // Cetak durasi

	if err != nil {
		fmt.Println("error writing to connection:", err)
		return
	}

}

func main() {
	ln, err := net.Listen("tcp", ":9292")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("server running on", ln.Addr().String())
	for {
		fmt.Println("running loop")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
