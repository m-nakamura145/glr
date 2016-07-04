package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RemoteHeadCommitId() (string, error) {
	cmd := exec.Command("git", "ls-remote", "origin", "HEAD")
	cmd.Stderr = os.Stderr

	id, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.Split(string(id), "\t")[0], nil
}

func LocalHeadCommitId() (string, error) {
	cmd := exec.Command("git", "show", "HEAD", "-s", "--format=%H")
	cmd.Stderr = os.Stderr

	id, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(id)), nil
}

func IsLatest(localId, remoteId string) bool {
	if localId == "" || remoteId == "" {
		return false
	}
	return localId == remoteId
}

func GitPull() string {
	cmd := exec.Command("git", "pull")
	cmd.Stderr = os.Stderr

	merged, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(merged), nil
}
