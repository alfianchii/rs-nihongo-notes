package main

import (
	"context"
	"os"

	"github.com/alfianchii/rs-nihongo-notes/internal/app"
	"github.com/alfianchii/rs-nihongo-notes/internal/cli"
	"github.com/alfianchii/rs-nihongo-notes/internal/utils"
)

func main() {
	flags, err := cli.Parse()
	utils.Must(err, "flag")

	err = app.Run(context.Background(), os.DirFS(flags.DocsRoot), flags)
	utils.Must(err, "app run")
}
