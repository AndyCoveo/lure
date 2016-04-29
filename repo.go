package main

import (
	"log"
	"os/exec"
	"time"

	"github.com/k0kubun/pp"
	"github.com/vsekhar/govtil/guid"
)

// This part interesting
// https://github.com/golang/go/blob/1441f76938bf61a2c8c2ed1a65082ddde0319633/src/cmd/go/vcs.go

func checkForUpdatesJob(projects []Project) {
	for {
		for _, project := range projects {
			pp.Println("updating: ", project.Remote)
			updateProject(project)
		}
		time.Sleep(60 * time.Second)
	}
}

func updateProject(project Project) {
	if project.Token == nil {
		return
	}

	repoGUID, err := guid.V4()
	if err != nil {
		log.Printf("Error: \"Could not generate guid\" %s", err)
		return
	}
	repoPath := "/tmp/" + repoGUID.String()

	if err := hgClone("https://x-token-auth:"+project.Token.AccessToken+"@"+project.Remote, repoPath); err != nil {
		log.Printf("Error: \"Could not clone\" %s", err)
		return
	}

}

func a() {
	repo := "ssh://hg@bitbucket.org/pastjean/dummy"
	repoGUID, err := guid.V4()
	if err != nil {
		log.Fatalf("Error: \"Could not generate guid\" %s", err)
	}

	repoPath := "/tmp/" + repoGUID.String()

	if err := hgClone(repo, repoPath); err != nil {
		log.Fatalf("Error: \"Could not clone\" %s", err)
	}

	if err := hgUpdate(repoPath, "default"); err != nil {
		log.Fatalf("Error: \"Could not update\" %s", err)
	}

	// TODO: verifier les dépendances

	// TODO: for each dependency to update
	// for () {
	if err := hgUpdate(repoPath, "default"); err != nil {
		log.Fatalf("Error: \"Could not update\" %s", err)
	}

	if err := hgBranch(repoPath, "lure-yournewbranchname"); err != nil {
		log.Fatalf("Error: \"Could not update\" %s", err)
	}
	// TODO: update dependency

	if err := hgCommit(repoPath, "MOTHERFUKING NEW DEPENDENCY"); err != nil {
		log.Fatalf("Error: \"Could not update\" %s", err)
	}

	// TODO: make pull request
	// }
}

func execute(pwd string, command string, params ...string) error {
	cmd := exec.Command(command, params...)
	cmd.Dir = pwd

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
