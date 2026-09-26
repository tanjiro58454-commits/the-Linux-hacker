package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	green  = "\033[92m"
	red    = "\033[91m"
	cyan   = "\033[96m"
	yellow = "\033[93m"
	reset  = "\033[0m"
	bold   = "\033[1m"
)

var packages = []string{
	"root-repo",
	"unstable-repo",
	"x11-repo",

	"termux-tools",
	"coreutils",
	"findutils",
	"diffutils",
	"which",
	"man",
	"procps",
	"psmisc",

	"python",
	"python2",
	"nodejs",
	"php",
	"ruby",
	"perl",
	"golang",
	"openjdk-17",

	"git",
	"make",
	"cmake",
	"autoconf",
	"automake",
	"pkg-config",

	"openssl",
	"curl",
	"wget",
	"openssh",

	"netcat-openbsd",
	"inetutils",
	"dnsutils",
	"nmap",
	"traceroute",
	"tcpdump",
	"whois",
	"w3m",

	"zip",
	"unzip",
	"tar",
	"unrar",
	"tree",
	"lf",

	"vim",
	"nano",
	"neovim",
	"fish",
	"tmux",
	"htop",
	"neofetch",
	"figlet",
	"cowsay",
	"cmatrix",
	"jp2a",
	"mpv",

	"termux-api",
	"proot",
	"proot-distro",
	"inxi",
	"tor",
	"findomain",
}

func run(command string, args ...string) bool {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	return cmd.Run() == nil
}

func banner() {
	fmt.Print("\033[2J\033[H")

	fmt.Println(green + bold + `
╔══════════════════════════════════════════╗
║                                          ║
║          THE LINUX HACKER               ║
║                                          ║
║       TERMUX SETUP ASSISTANT            ║
║                                          ║
╚══════════════════════════════════════════╝
` + reset)

	fmt.Println(cyan + "        Learn • Build • Explore • Secure" + reset)
	fmt.Println()
}

func wait() {
	fmt.Print("\nPress ENTER to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func updateSystem() {
	fmt.Println(cyan + "\n[+] Running apt update..." + reset)

	if !run("apt", "update") {
		fmt.Println(red + "✗ apt update failed." + reset)
		return
	}

	fmt.Println(green + "\n✓ apt update completed." + reset)

	fmt.Println(cyan + "\n[+] Running apt upgrade..." + reset)

	if run("apt", "upgrade", "-y") {
		fmt.Println(green + "\n✓ System upgrade completed." + reset)
	} else {
		fmt.Println(red + "\n✗ apt upgrade failed." + reset)
	}
}

func masterSetup() {
	fmt.Println(yellow + `
⚠ MASTER BASIC SETUP

This will install a collection of
Termux, Linux, development, networking
and terminal utilities.

C = Continue
E = Exit
` + reset)

	fmt.Print("Enter C/E: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	if input != "c" {
		fmt.Println(yellow + "\nReturning to menu..." + reset)
		return
	}

	fmt.Println(cyan + "\n[+] Updating package information..." + reset)

	if !run("apt", "update") {
		fmt.Println(red + "\n✗ apt update failed. Setup stopped." + reset)
		return
	}

	total := len(packages)
	installed := 0
	failed := 0

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println(bold + "          MASTER BASIC SETUP" + reset)
	fmt.Println(strings.Repeat("=", 50))

	for i, pkg := range packages {
		fmt.Printf("\n%s[%d/%d] Installing %s%s\n",
			cyan, i+1, total, pkg, reset)

		if run("apt", "install", "-y", pkg) {
			fmt.Println(green + "✓ Installed: " + pkg + reset)
			installed++
		} else {
			fmt.Println(red + "✗ Failed/unavailable: " + pkg + reset)
			failed++
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println(bold + "             SETUP SUMMARY" + reset)
	fmt.Println(strings.Repeat("=", 50))

	fmt.Printf(green+"✓ Installed: %d\n"+reset, installed)
	fmt.Printf(red+"✗ Failed:    %d\n"+reset, failed)

	fmt.Println(green + "\n✓ Master Basic Setup finished." + reset)
}

func main() {
	// Check whether apt exists.
	if _, err := exec.LookPath("apt"); err != nil {
		fmt.Println(red + "✗ apt was not found." + reset)
		fmt.Println("This program is designed for Termux.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		banner()

		fmt.Println("1. Check apt update && upgrade")
		fmt.Println("2. Master Basic Setup")
		fmt.Println("3. Exit")

		fmt.Print("\nSelect an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		case "1":
			updateSystem()
			wait()

		case "2":
			masterSetup()
			wait()

		case "3":
			fmt.Println(green + "\nGoodbye! 👋" + reset)
			return

		default:
			fmt.Println(red + "\nInvalid option." + reset)
			wait()
		}
	}
}
