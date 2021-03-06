package captain // import "github.com/harbur/captain/captain"

import (
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
)

var endpoint = "unix:///var/run/docker.sock"
var client, _ = docker.NewClient(endpoint)

func buildImage(dockerfile string, image string, tag string) error {
	info("Building image %s:%s", image, tag)

	opts := docker.BuildImageOptions{
		Name:                image + ":" + tag,
		Dockerfile:          dockerfile,
		NoCache:             false,
		SuppressOutput:      false,
		RmTmpContainer:      true,
		ForceRmTmpContainer: true,
		OutputStream:        os.Stdout,
		ContextDir:          ".",
	}
	err := client.BuildImage(opts)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return err
}

func tagImage(repo string, origin string, tag string) error {
	info("Tagging image %s:%s as %s:%s", repo, origin, repo, tag)
	// var imageID = getImageID(repo, origin)
	opts := docker.TagImageOptions{Repo: repo, Tag: tag, Force: true}
	err := client.TagImage(repo, opts)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return err
}

func getImageID(repo string, tag string) string {
	images, _ := client.ListImages(docker.ListImagesOptions{})
	for _, image := range images {
		for _, b := range image.RepoTags {
			if b == repo+":"+tag {
				return image.ID
			}
		}
	}
	return ""
}

func imageExist(repo string, tag string) bool {
	images, _ := client.ListImages(docker.ListImagesOptions{})
	for _, image := range images {
		for _, b := range image.RepoTags {
			if b == repo+":"+tag {
				return true
			}
		}
	}
	return false
}
