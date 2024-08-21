package main

import (
	"fmt"
	"log"
	"regexp"
	"time"

	expect "github.com/google/goexpect"
)

func must(err error) {
	if err != nil {
		log.Panic(fmt.Errorf("error found: %e", err))
	}
}

func main() {
	// START OMIT
	timeout := 1 * time.Second // you'll see why we added this...
	conn, _, err := expect.Spawn("/bin/sh", timeout)
	must(err)
	defer conn.Close()

	prompt := regexp.MustCompile(`\$ `)
	_, _, err = conn.Expect(prompt, timeout)
	must(err)
	must(conn.Send("echo Hello, Expect!\n"))
	_, _, err = conn.Expect(regexp.MustCompile(`Hello, Expect!`), timeout)
	must(err)

	_, _, err = conn.Expect(prompt, timeout)
	must(err)
	must(conn.Send("exit\n"))
	//TODO: eof?
	// END OMIT
}
