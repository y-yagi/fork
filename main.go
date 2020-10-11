package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/y-yagi/color"
)

func main() {
	root := "/home/y-yagi/src/github.com/"
	user := "y-yagi"
	green := color.New(color.FgGreen, color.Bold).SprintFunc()

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "please specify repository\n")
		return
	}

	repo := strings.Split(os.Args[1], "/")
	if len(repo) != 2 {
		fmt.Fprintf(os.Stderr, "please specify repository as 'USERNAME/REPOSITORY' format\n")
		return
	}

	fmt.Printf("%s\n", green("clone repository"))
	dir := filepath.Join(root, repo[0])
	if err := os.MkdirAll(dir, 0700); err != nil {
		fmt.Fprintf(os.Stderr, "cannot create directory: %v\n", err)
		return
	}

	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(os.Stderr, "cannot change directory: %v\n", err)
		return
	}

	cmd := exec.Command("git", "clone", fmt.Sprintf("git@github.com:%s/%s.git", user, repo[1]))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot clone repository: %v\n", err)
		return
	}

	fmt.Printf("%s\n", green("add remote"))
	if err := os.Chdir(repo[1]); err != nil {
		fmt.Fprintf(os.Stderr, "cannot change directory: %v\n", err)
		return
	}

	cmd = exec.Command("git", "remote", "add", "upstream", fmt.Sprintf("git@github.com:%s/%s.git", repo[0], repo[1]))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot clone repository: %v\n", err)
		return
	}

	fmt.Printf("%s\n", green("complete"))
}
