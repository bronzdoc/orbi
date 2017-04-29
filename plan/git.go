package plan

import (
	"fmt"
	"os"
	"regexp"

	"gopkg.in/src-d/go-git.v4"
)

func Clone(repo_url string) error {
	reg, err := regexp.Compile(`(.*/)|(\.git)`)
	repo_name := reg.ReplaceAllString(repo_url, "")
	path_to_clone := fmt.Sprintf("/home/bronzdoc/.orbi/plans/%s", repo_name)

	_, err = git.PlainClone(path_to_clone, false, &git.CloneOptions{
		URL:      repo_url,
		Progress: os.Stdout,
	})

	return err
}
