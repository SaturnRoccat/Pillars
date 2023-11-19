package main

import (
	"fmt"

	"github.com/duckos-Mods/Pillars/main/PC"
)

func build(pathToRoot string) error {
	// Get project config
	var config, err = PC.PullJson[PC.ProjectConfig](pathToRoot + "/Pillars/ProjectConfig.json")
	if err != nil {
		return err
	}
	return buildAddon(
		[]string{config.BPPath, config.RPPath},
		[]string{fmt.Sprintf("%s/development_behavior_packs/%s_BP", MCBEPath, config.ProjectName), fmt.Sprintf("%s/development_resource_packs/%s_RP", MCBEPath, config.ProjectName)},
		false,
		pathToRoot+"/Pillars/Temp",
	)
}

func buildAddon(sources, targets []string, bypassCache bool, pathToTemp string) error {

	for i := 0; i < len(sources); i++ {
		source := sources[i]
		target := targets[i]
		// get the files to ignore
		var filesToIgnore = getFilesToIgnore(source)

		// Copy all files to the temp dir
		var err = PC.BulkFileCopy(source, pathToTemp, filesToIgnore)
		if err != nil {
			return err
		}

		/*
			TODO:
				- Call any external plugins that request to be called on build (e.g. a plugin that minifies files)
		*/

		println("Done building. For " + target + "!...")
		// Copy all files from the temp dir to the build dir
		// err = PC.BulkFileCopy(source, target, nil)

		// clear the temp dir

		err = PC.DeleteContents(pathToTemp)
		if err != nil {
			return err
		}
	}

	return nil
}

// D:\\Projects\\Go\\Pillars\\CompilerTests/Pillars/Temp
