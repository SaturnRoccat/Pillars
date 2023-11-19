package Pillars

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duckos-Mods/Pillars/Pillars/PC"
)

func build(pathToRoot string, bypassCache bool) error {
	// Get project config
	var config, err = PC.PullJson[PC.ProjectConfig](pathToRoot + "/Pillars/ProjectConfig.json")
	if err != nil {
		return err
	}
	// if bypassCache is true, set cache data to {"fileArray":{}}
	if bypassCache {
		var FEDP = pathToRoot + "/Pillars/FileEditTimes.json"
		var FETJson = PC.ProjectFileJson{FileArray: make(map[string]int64)}
		var json, _ = json.Marshal(FETJson)
		os.WriteFile(FEDP, json, 0777)
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
		// Create the target dir
		err = os.MkdirAll(target, 0777)
		if err != nil {
			return err
		}

		// Copy all files from the temp dir to the build dir
		err = PC.BulkFileCopy(source, target, nil)
		if err != nil {
			return err
		}

		// clear the temp dir

		err = PC.DeleteContents(pathToTemp)
		if err != nil {
			return err
		}
	}

	return nil
}

// D:\\Projects\\Go\\Pillars\\CompilerTests/Pillars/Temp
