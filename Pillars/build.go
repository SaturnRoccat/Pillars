package Pillars

import (
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
		PC.WriteEmptyJson(pathToRoot + "/Pillars/BPFileEditTimes.json")
		PC.WriteEmptyJson(pathToRoot + "/Pillars/RPFileEditTimes.json")
	}
	err = buildAddon(config.BPPath, fmt.Sprintf("%s/development_behavior_packs/%s_BP", MCBEPath, config.ProjectName), bypassCache, pathToRoot+"/Pillars/Temp", pathToRoot+"/Pillars/BPFileEditTimes.json")
	if err != nil {
		return err
	}

	return buildAddon(config.RPPath, fmt.Sprintf("%s/development_resource_packs/%s_RP", MCBEPath, config.ProjectName), bypassCache, pathToRoot+"/Pillars/Temp", pathToRoot+"/Pillars/RPFileEditTimes.json")
}

func buildAddon(source, target string, bypassCache bool, pathToTemp string, FETP string) error {

	// get the files to ignore
	var filesToIgnore = getFilesToIgnore(source, FETP)

	// Copy all files to the temp dir
	var err = PC.BulkFileCopy(source, pathToTemp, filesToIgnore)
	if err != nil {
		return err
	}

	/*
		TODO:
			- Call any external plugins that request to be called on build (e.g. a plugin that minifies files)
			- Call all stages that are in the build pipeline (e.g. a stage that compiles typescript to javascript) or any custom stages that are in the build pipeline
	*/

	println("Done building. For " + target + "!...")
	// Create the target dir
	err = os.MkdirAll(target, 0777)
	if err != nil {
		return err
	}

	// Copy all files from the temp dir to the build dir
	err = PC.BulkFileCopy(pathToTemp, target, filesToIgnore)
	if err != nil {
		return err
	}

	// clear the temp dir

	err = PC.DeleteContents(pathToTemp)
	if err != nil {
		return err
	}

	return nil
}

// D:\\Projects\\Go\\Pillars\\CompilerTests/Pillars/Temp
