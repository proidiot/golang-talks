package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	goexpect "github.com/google/goexpect"
)

func main() {
	// START OMIT
	timeout := 1 * time.Second
	conn, _, err := goexpect.Spawn("/bin/sh", timeout)
	if err != nil {
		log.Panic(fmt.Errorf("unable to spawn: %e", err))
	}
	defer conn.Close()
	log.Println("process spawned")

	prompt := regexp.MustCompile(`\$ `)
	a, b, e := conn.Expect(prompt, timeout)
	log.Printf("prompt seen: %v %v %e", a, b, e)
	e = conn.Send("echo Hello, Expect!\n")
	log.Printf("command sent: %e", e)
	a, b, e = conn.Expect(regexp.MustCompile(`Hello, Expect!`), timeout)
	log.Printf("response seen: %v %v %e", a, b, e)

	a, b, e = conn.Expect(prompt, timeout)
	log.Printf("prompt seen: %v %v %e", a, b, e)
	e = conn.Send("exit\n")
	log.Printf("command sent: %e", e)
	//TODO: eof?
	// END OMIT
}
