package main

import (
	"github.com/duckos-Mods/Pillars/main/PC"
)

func build(sources, targets []string, bypassCache bool, pathToTemp string) error {

	for _, source := range sources {
		for _, target := range targets {
			// get the files to ignore
			var filesToIgnore = getFilesToIgnore(source)

			// Copy all files to the temp dir
			var err = PC.BulkFileCopy(source, pathToTemp+"/Pillars/Temp", filesToIgnore)
			if err != nil {
				return err
			}

			/*
				TODO:
					- Call any external plugins that request to be called on build (e.g. a plugin that minifies files)
			*/

			// Copy all files from the temp dir to the build dir
			err = PC.BulkFileCopy(source, target, nil)
		}
	}

	return nil
}
