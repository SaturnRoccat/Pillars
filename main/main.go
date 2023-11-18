package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	VERSION = []int{0, 0, 1}

	InitMode  = flag.NewFlagSet("init", flag.ExitOnError)
	BuildMode = flag.NewFlagSet("build", flag.ExitOnError)
)

func gwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

func main() {
	// add the init flags
	projectName := InitMode.String("projectName", "", "The name of the project")
	pathToRoot := InitMode.String("pathToRoot", gwd(), "The path to the root of the project")

	// add the build flags
	pathToRootBuild := BuildMode.String("pathToRoot", gwd(), "The path to the root of the project")
	bypassCache := BuildMode.Bool("bypassCache", false, "Bypass the cache and recompile everything. Also regenerates the cache.")

	fmt.Println("Started The Pillar Compiler! For MCBE add-on development! You are using version " + fmt.Sprint(VERSION) + ".")

	if len(os.Args) < 2 {
		fmt.Println("You must specify a mode! (init or build)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		InitMode.Parse(os.Args[2:])
		if *projectName == "" {
			fmt.Println("You must specify a project name!")
			os.Exit(1)
		}

		fmt.Println("Creating default state...")
		err := createDefaultState(*pathToRoot, *projectName)
		if err != nil {
			panic(err)
		}
		fmt.Println("Done! Creating Default Config...")
	case "build":
		BuildMode.Parse(os.Args[2:])
		fmt.Println("Building...")
	}

}
