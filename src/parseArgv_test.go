package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"testing"
)

func getExitStatus(err error) int {
	exitErr, ok := err.(*exec.ExitError)

	if !ok {
		return -1
	}
	waitStatus, ok2 := exitErr.Sys().(syscall.WaitStatus)
	if !ok2 {
		return -1
	}
	return waitStatus.ExitStatus()
}

func getCaller() string {
	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])

	startName := strings.LastIndex(f.Name(), ".") + 1

	return f.Name()[startName:]
}

func testParseArgvFailure(t *testing.T, ndx int) int {
	cmd := exec.Command(os.Args[0], "-test.run="+getCaller())
	cmd.Env = append(os.Environ(), "TEST_ARGV="+strconv.Itoa(ndx))
	err := cmd.Run()
	return getExitStatus(err)
}

func reviewResults(t *testing.T, argv []string, status int) bool {
	if status == FAILURE {
		return true
	}

	fmt.Printf("ParseArgv(%v) -> ", argv)
	if status == -1 {
		fmt.Println("The process didn't exit as expected")
	} else if status != FAILURE {
		fmt.Printf("Expected exit status %d, got %d\n", FAILURE, status)
	}
	return false
}

// TestInvalidNbArgv tests if ParseArgv exits FAILURE when it as an invalid
// number of arguments
func TestInvalidNbArgv(t *testing.T) {
	testArgv := [][]string{
		[]string{},
		[]string{"100"},
		[]string{"100", "15", "90", "95", "5"},
	}
	res := true

	if testNdx := os.Getenv("TEST_ARGV"); testNdx == "" {
		for i := range testArgv {
			status := testParseArgvFailure(t, i)
			res = reviewResults(t, testArgv[i], status) && res
		}
	} else {
		n, _ := strconv.Atoi(testNdx)
		ParseArgv(testArgv[n])
	}

	if !res {
		t.Fatalf("One or more test failed")
	}
}

// TestInvalidValuesArgv tests if ParseArgv exits FAILURE when some of is
// arguments have bad values
func TestInvalidValuesArgv(t *testing.T) {
	testArgv := [][]string{
		[]string{"abc", "15"},
		[]string{"-1", "15"},
		[]string{"0", "15"},
		[]string{"100", "abc"},
		[]string{"100", "-42"},
		[]string{"100", "0"},
		[]string{"100", "15", "abd"},
		[]string{"100", "15", "-123"},
		[]string{"100", "15", "0"},
		[]string{"100", "15", "90", "def"},
		[]string{"100", "15", "90", "-123"},
		[]string{"100", "15", "90", "0"},
		[]string{"100", "15", "90", "85"},
		[]string{"100", "15", "90", "90"},
		[]string{"100", "15", "201"},
		[]string{"100", "15", "201", "210"},
		[]string{"100", "15", "201", "200"},
		[]string{"100", "15", "100", "201"},
	}
	res := true

	if testNdx := os.Getenv("TEST_ARGV"); testNdx == "" {
		for i := range testArgv {
			status := testParseArgvFailure(t, i)
			res = reviewResults(t, testArgv[i], status) && res
		}
	} else {
		n, _ := strconv.Atoi(testNdx)
		ParseArgv(testArgv[n])
	}

	if !res {
		t.Fatalf("One or more test failed")
	}
}

// TestValidValidArgv tests if ParseArgv returns the right params with good
// arguments
func TestValidValidArgv(t *testing.T) {
	argvs := [][]string{
		[]string{"100", "15"},
		[]string{"100", "15", "90"},
		[]string{"100", "15", "90", "95"},
	}
	expect := []*Params{
		&Params{u: 100, s: 15, iq1: DefaultIQ1, iq2: DefaultIQ2, study: Density},
		&Params{u: 100, s: 15, iq1: 90, iq2: DefaultIQ2, study: PercentageBelow},
		&Params{u: 100, s: 15, iq1: 90, iq2: 95, study: PercentageBetween},
	}

	for i := range argvs {
		params := ParseArgv(argvs[i])
		if !reflect.DeepEqual(params, expect[i]) {
			t.Fatalf("ParseArgv(%v) ->\n\texpected: %+v\n\tgot: %+v\n",
				argvs[i],
				params,
				expect[i])
		}
	}
}
