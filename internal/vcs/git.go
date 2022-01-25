package vcs

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func UseGit(dir string) error {

	r, err := git.PlainInit(dir, false)
	if err != nil {
		fmt.Println("could not open the git repo")
		os.Exit(1)
	}

	h := plumbing.NewSymbolicReference(plumbing.HEAD, plumbing.ReferenceName("refs/heads/develop"))
	err = r.Storer.SetReference(h)
	if err != nil {
		return err
	}
	// this https://github.com/hairyhenderson/gomplate/blob/c14539eca2ef59445c4287f83b2377e3baefb50c/data/datasource_git_test.go
	// also set the default branch, but testing found that the key
	// change was setting the new Storer reference above
	// so did not need to do this?
	// c, _ := r.Config()
	// c.Init.DefaultBranch = "develop"
	// err = r.Storer.SetConfig(c)

	// another example to checkout is
	// https://github.com/hairyhenderson/go-fsimpl/blob/main/gitfs/git.go

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// err = w.Checkout(&git.CheckoutOptions{
	// 	Branch: plumbing.NewBranchReferenceName("develop"),
	// 	Create: true,
	// 	Force:  true,
	// 	Keep:   true,
	// })
	// if err != nil {
	// 	fmt.Println("could not checkout develop")
	// 	os.Exit(1)
	// }

	err = w.AddGlob("*")
	if err != nil {
		return err
	}
	_, err = w.Commit("init: project scaffold", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "mrgfab",
			Email: "no-reply",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}
	return nil
}
