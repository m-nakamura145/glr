package main

import (
	"fmt"
)

type Sync struct{}

func (s *Sync) Help() string {
	return "glr sync Help"
}

func (s *Sync) Run(args []string) int {
	localId, err := LocalHeadCommitId()
	if err != nil {
		fmt.Println(err)
	}

	remoteId, err := RemoteHeadCommitId()
	if err != nil {
		fmt.Println(err)
	}

	if IsLatest(localId, remoteId) {
		return 0
	}

	GitPull()
	return 0
}

func (s *Sync) Synopsis() string {
	return "Synchronize local git repository with remote at once"
}
