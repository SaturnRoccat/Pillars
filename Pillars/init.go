package Pillars

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duckos-Mods/Pillars/Pillars/PC"
)

var (
	ProjectConfigPath = ""
	FileEditTimesPath = ""

	// TODO: Make this configurable
	MCBEPath = os.Getenv("LOCALAPPDATA") + "/Packages/Microsoft.MinecraftUWP_8wekyb3d8bbwe/LocalState/games/com.mojang"
)

func writeDefaultConfig(pathToConfig string, config PC.ProjectConfig) {
	json, _ := json.Marshal(config)
	ProjectConfigPath = pathToConfig
	os.WriteFile(pathToConfig, []byte(json), 0777)
}

func writeCJson(pathToCJson string, fjson PC.ProjectFileJson) {
	json, _ := json.Marshal(fjson)
	FileEditTimesPath = pathToCJson
	os.WriteFile(pathToCJson, []byte(json), 0777)
}

func createDefaultState(pathToRoot string, projectName string) error {
	// Make the Pillars directory
	err := os.Mkdir(fmt.Sprintf("%s/Pillars", pathToRoot), 0777)
	if err != nil {
		return err
	}

	// Make the Pillars/ProjectConfig.json file
	_, err = os.Create(fmt.Sprintf("%s/Pillars/ProjectConfig.json", pathToRoot))
	if err != nil {
		return err
	}

	// Make the Pillars/FileEditTimes.json
	_, err = os.Create(fmt.Sprintf("%s/Pillars/FileEditTimes.json", pathToRoot))
	if err != nil {
		return err
	}

	// Make the Pillars/ProjectName.txt
	_, err = os.Create(fmt.Sprintf("%s/Pillars/ProjectName.txt", pathToRoot))
	if err != nil {
		return err
	}

	// Make the Pillars/Temp directory
	err = os.Mkdir(fmt.Sprintf("%s/Pillars/Temp", pathToRoot), 0777)
	if err != nil {
		return err
	}

	// Create the RP and BP directories
	err = os.Mkdir(fmt.Sprintf("%s/%s_Bp", pathToRoot, projectName), 0777)
	if err != nil {
		return err
	}
	err = os.Mkdir(fmt.Sprintf("%s/%s_Rp", pathToRoot, projectName), 0777)
	if err != nil {
		return err
	}

	// Write the project name to the ProjectName.txt file
	err = os.WriteFile(fmt.Sprintf("%s/Pillars/ProjectName.txt", pathToRoot), []byte(projectName), 0777)
	if err != nil {
		return err
	}

	defaultConfig := PC.ProjectConfig{
		Version:     VERSION,
		ProjectName: projectName,
		BPPath:      fmt.Sprintf("%s/%s_Bp", pathToRoot, projectName),
		RPPath:      fmt.Sprintf("%s/%s_Rp", pathToRoot, projectName),
		TempPath:    fmt.Sprintf("%s/Pillars/Temp", pathToRoot),
	}
	// Write the default config to the ProjectConfig.json file
	writeDefaultConfig(fmt.Sprintf("%s/Pillars/ProjectConfig.json", pathToRoot), defaultConfig)

	// Write the default file edit times to the FileEditTimes.json file
	defaultFileEditTimes := PC.ProjectFileJson{
		FileArray: map[string]int64{},
	}

	writeCJson(fmt.Sprintf("%s/Pillars/FileEditTimes.json", pathToRoot), defaultFileEditTimes)

	return nil
}
