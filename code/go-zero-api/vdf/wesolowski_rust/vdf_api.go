package wesolowski_rust

import (
	"os/exec"
	"strconv"
)

func checkVDFExists() error {
	_, err := exec.LookPath("./vdf-cli")
	return err
}

func execCmd(cmd string, ch chan []byte) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		close(ch)
		return
	}
	ch <- out
	close(ch)
}

func Execute(challenge int64, iterations int) []byte {

	// channel for receving result of vdf command
	ch1 := make(chan []byte)
	// channel for receving result of taskset command
	ch2 := make(chan []byte)

	if checkVDFExists() != nil {
		return nil
	}

	cmd1 := "./vdf-cli " + strconv.FormatInt(challenge, 16) + " " + strconv.Itoa(iterations)
	cmd2 := "taskset -pc 0 $(pidof vdf-cli)"

	go execCmd(cmd1, ch1)
	go execCmd(cmd2, ch2)

	out1 := <-ch1

	return out1
}

func Verify(challenge int64, iterations int, proof string) bool {

	// channel for receving result of vdf command
	ch1 := make(chan []byte)
	// channel for receving result of taskset command
	ch2 := make(chan []byte)

	if checkVDFExists() != nil {
		return false
	}

	cmd1 := "./vdf-cli " + strconv.FormatInt(challenge, 16) + " " + strconv.Itoa(iterations) + " " + proof
	cmd2 := "taskset -pc 0 $(pidof vdf-cli)"

	go execCmd(cmd1, ch1)
	go execCmd(cmd2, ch2)

	out1 := <-ch1

	if string(out1) == "Proof is valid\n" {
		return true
	}
	return false
}
