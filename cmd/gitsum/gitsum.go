/**
 * Copyright(c) 2024 Michael Estrin.  All rights reserved.
 */

package main

import (
	"flag"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func checkIfError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	author := flag.String("author", "", "author for commits to extract")
	in := flag.String("in", "", "path to input project")
	out := flag.String("out", "", "path to output project")
	flag.Parse()

	if *author == "" || *in == "" || *out == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	r, err := git.Open(filesystem.NewStorage(osfs.New(filepath.Join(*in, ".git")), cache.NewObjectLRUDefault()),
		osfs.New(*in),
	)
	checkIfError(err)

	ref, err := r.Head()
	checkIfError(err)

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	checkIfError(err)

	commits := make([]*object.Commit, 0)
	err = cIter.ForEach(func(c *object.Commit) error {
		if strings.Contains(c.Author.Name, *author) || strings.Contains(c.Author.Email, *author) {
			commits = append(commits, c)
		}
		return nil
	})
	checkIfError(err)

	r, err = git.Init(
		filesystem.NewStorage(osfs.New(filepath.Join(*out, ".git")), cache.NewObjectLRUDefault()),
		osfs.New(*out),
	)
	checkIfError(err)

	w, err := r.Worktree()
	checkIfError(err)

	fn := "hash"

	for l := len(commits); l > 0; l-- {
		c := commits[l-1]

		err = os.WriteFile(filepath.Join(*out, fn), []byte(c.Hash.String()), os.ModePerm)
		checkIfError(err)

		_, err = w.Add(fn)
		checkIfError(err)

		_, err := w.Commit(c.Hash.String(), &git.CommitOptions{
			All:    true,
			Author: &c.Author,
		})
		checkIfError(err)
	}
}
