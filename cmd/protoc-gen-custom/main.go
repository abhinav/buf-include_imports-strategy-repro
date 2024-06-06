package main

import (
	"context"
	"log"
	"slices"
	"sort"
	"strings"

	"github.com/bufbuild/protoplugin"
)

func main() {
	log.SetFlags(0)
	protoplugin.Main(protoplugin.HandlerFunc(handle))
}

func handle(
	_ context.Context,
	_ protoplugin.PluginEnv,
	resw protoplugin.ResponseWriter,
	req protoplugin.Request,
) error {
	files := slices.Clone(req.CodeGeneratorRequest().FileToGenerate)
	sort.Strings(files)
	log.Printf("protoc: %s", strings.Join(files, ", "))
	return nil
}
