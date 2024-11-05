package git

import (
	"strings"

	"github.com/joshmedeski/sesh/shell"
)

type Git interface {
	ShowTopLevel(name string) (bool, string, error)
	GitCommonDir(name string) (bool, string, error)
	GitMainWorktree(name string) (bool, string, error)
	Clone(name string) (string, error)
}

type RealGit struct {
	shell shell.Shell
}

func NewGit(shell shell.Shell) Git {
	return &RealGit{shell}
}

func (g *RealGit) ShowTopLevel(path string) (bool, string, error) {
	out, err := g.shell.Cmd("git", "-C", path, "rev-parse", "--show-toplevel")
	if err != nil {
		return false, "", err
	}
	return true, out, nil
}

func (g *RealGit) GitCommonDir(path string) (bool, string, error) {
	out, err := g.shell.Cmd("git", "-C", path, "rev-parse", "--git-common-dir")
	if err != nil {
		return false, "", err
	}
	return true, out, nil
}

func (g *RealGit) GitMainWorktree(path string) (bool, string, error) {
	out, err := g.shell.Cmd("git", "-C", path, "worktree", "list")
	if err != nil {
		return false, "", err
	}
	main := strings.Fields(out)[0]
	// TODO: does strings.Fields need err handling too?
	return true, main, nil
}

func (g *RealGit) Clone(name string) (string, error) {
	out, err := g.shell.Cmd("git", "clone", name)
	if err != nil {
		return "", err
	}
	return out, nil
}
