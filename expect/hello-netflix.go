package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	expect "github.com/Netflix/go-expect"
)

func must(err error) {
	if err != nil {
		log.Panic(fmt.Errorf("error found: %e", err))
	}
}

func main() {
	// START OMIT
	conn, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	must(err)
	defer conn.Close()

	cmd := exec.Command("/bin/sh")
	cmd.Stdin = conn.Tty()
	cmd.Stdout = conn.Tty()
	cmd.Stderr = conn.Tty()

	prompt := expect.String(`\$ `)
	_, err = conn.Expect(prompt)
	must(err)
	must(conn.Send("echo Hello, Expect!\n"))
	_, err = conn.Expect(expect.String(`Hello, Expect!`))
	must(err)

	_, err = conn.Expect(prompt)
	must(err)
	must(conn.Send("exit\n"))
	must(conn.ExpectEOF())
	// END OMIT
}
