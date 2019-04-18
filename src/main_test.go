package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func getOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC
	return out
}

// TestStudyDensity tests if studyDensity displays correctly
func TestStudyDensity(t *testing.T) {
	out := getOutput(func() {
		studyDensity(100, 15)
	})
	expect, err := ioutil.ReadFile("../examples/example_density.txt")

	if err != nil {
		fmt.Println(err)
		t.Fatalf("")
	}

	if out != string(expect) {
		t.Fatalf("diff between output and example")
	}
}

// TestStudyIQBelow tests if studyDensity displays correctly
func TestStudyIQBelow(t *testing.T) {
	out := getOutput(func() {
		studyIQBelow(100, 24, 90)
	})
	expect, err := ioutil.ReadFile("../examples/example_iq_below.txt")

	if err != nil {
		fmt.Println(err)
		t.Fatalf("")
	}

	if out != string(expect) {
		t.Fatalf("diff between output and example:\n\texpected: %s\tgot: %s\t",
			expect,
			out)
	}
}

// TestStudyIQBetween tests if studyDensity displays correctly
func TestStudyIQBetween(t *testing.T) {
	out := getOutput(func() {
		studyIQBetween(100, 24, 90, 95)
	})
	expect, err := ioutil.ReadFile("../examples/example_iq_between.txt")

	if err != nil {
		fmt.Println(err)
		t.Fatalf("")
	}

	if out != string(expect) {
		t.Fatalf("diff between output and example:\n\texpected: %s\tgot: %s\t",
			expect,
			out)
	}
}
