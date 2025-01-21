package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"

	// Packages

	image "github.com/docker/docker/api/types/image"
	client "github.com/docker/docker/client"
)

var (
	flagManifest = flag.String("name", "", "Name of the manifest")
	flagImages   = flag.String("images", "", "Comma separated list of images to include in the manifest")
	flagPush     = flag.Bool("push", true, "Push the manifest to the registry")
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Parse command line flags
	flag.Parse()

	// Create a new docker client
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create docker client:", err)
		os.Exit(1)
	}
	defer client.Close()

	for _, ref := range strings.Split(*flagImages, ",") {
		// Pull an image
		if err := ImagePull(ctx, client, ref); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to pull image:", err)
			os.Exit(1)
		}
	}
}

func ImagePull(ctx context.Context, cli *client.Client, ref string) error {
	reader, err := cli.ImagePull(ctx, ref, image.PullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = io.Copy(os.Stdout, reader)
	return err
}
