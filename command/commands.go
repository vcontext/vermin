package command

import (
	"fmt"
	"github.com/mhewedy/vermin/db"
	"runtime"
)

func VBoxManage(args ...string) *cmd {
	return &cmd{
		command: "vboxmanage",
		args:    args,
	}
}

func Scp(args ...string) *cmd {
	return &cmd{
		command: "scp",
		args:    args,
	}
}

func Arp(args ...string) *cmd {
	return &cmd{
		command: "arp",
		args:    args,
	}
}

func Wget(url string, file string) *cmd {
	if runtime.GOOS == "windows" {
		return &cmd{
			command: fmt.Sprintf("(New-Object System.Net.WebClient).DownloadFile('%s', '%s')", url, file),
		}
	} else {
		return &cmd{
			command: "wget",
			args: []string{
				"-O",
				file,
				url,
			},
		}
	}
}

func Ssh(ipAddr string, extraArgs ...string) *cmd {

	args := []string{"-i", db.GetPrivateKeyPath(), "-o", "StrictHostKeyChecking=no", db.GetUsername() + "@" + ipAddr}
	args = append(args, extraArgs...)

	return &cmd{
		command: "ssh",
		args:    args,
	}
}

func Ping(ip string) *cmd {
	if runtime.GOOS == "windows" {
		return &cmd{
			command: "ping",
			args:    []string{"-n", "1", "-w", "0.1", ip},
		}
	} else {
		return &cmd{
			command: "ping",
			args:    []string{"-c", "1", "-W", "0.1", ip},
		}
	}
}
