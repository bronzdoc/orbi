package plan

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"

	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

func Clone(repo_url string) error {
	reg, err := regexp.Compile(`(.*/)|(\.git)`)
	if err != nil {
		return err
	}

	u, err := url.Parse(repo_url)
	if err != nil {
		return fmt.Errorf("Couldn't parse %s: %s", repo_url, err)
	}

	repo_name := reg.ReplaceAllString(repo_url, "")
	path_to_clone := fmt.Sprintf("%s/.orbi/plans/%s", os.Getenv("HOME"), repo_name)

	switch u.Scheme {
	default:
		return fmt.Errorf("URL scheme \"%s\": for %s is not supported", u.Scheme, repo_url)
	/*
	 Check for empty string since a common url for cloning a git repo is
	 the shorter scp-like syntax url (i.e user@server:project.git).
	 This url is schemeless so u.Scheme will return an empty string for this
	 kind of url.
	*/
	case "":
		// Pattern to match the shorter url scp-like syntax for the SSH protocol
		ssh_pattern := `^(\w*@)?[A-Za-z0-9-._]+:[A-Za-z0-9-._]+/`

		reg, err := regexp.Compile(ssh_pattern)
		if err != nil {
			return err
		}

		is_ssh := reg.MatchString(repo_url)

		if is_ssh {
			if err := sshClone(repo_url, path_to_clone); err != nil {
				return err
			}
		}
	case "ssh":
		if err := sshClone(repo_url, path_to_clone); err != nil {
			return err
		}
	case "https":
		if err := httpsClone(repo_url, path_to_clone); err != nil {
			return err
		}
	}

	return nil
}

func sshClone(repo_url, path_to_clone string) error {
	/*
		TODO when a new release of go-git be available use this code to get an authMethod.

		auth, err := gitssh.NewPublicKeysFromFile("user", "pub_path")
		if err != nil {
			return err
		}
	*/

	key_path := fmt.Sprintf("%s/.ssh/id_rsa", os.Getenv("HOME"))
	key, err := ioutil.ReadFile(key_path)
	if err != nil {
		return fmt.Errorf("Couldn't read private key file: %s", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return fmt.Errorf("Couldn't parse private key: %s", err)
	}

	_, err = git.PlainClone(path_to_clone, false, &git.CloneOptions{
		Auth: &gitssh.PublicKeys{
			User:   "git",
			Signer: signer,
		},
		URL:      repo_url,
		Progress: os.Stdout,
	})

	if err != nil {
		return fmt.Errorf("Couldn't get plan: %s", err)
	}

	return nil
}

func httpsClone(repo_url, path_to_clone string) error {
	_, err := git.PlainClone(path_to_clone, false, &git.CloneOptions{
		URL:      repo_url,
		Progress: os.Stdout,
	})

	if err != nil {
		return fmt.Errorf("Couldn't get plan: %s", err)
	}

	return nil
}
