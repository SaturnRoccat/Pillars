package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ProjectConfig struct {
	Version     []int  `json:"version"`
	ProjectName string `json:"projectName"`
	BPPath      string `json:"bpPath"`
	RPPath      string `json:"rpPath"`
	TempPath    string `json:"tempPath"`
}

type ProjectFileNameAndEditTime struct {
	FileName string `json:"fileName"`
	EditTime int64  `json:"editTime"`
}

type ProjectFileJson struct {
	// We need to store an array of file names and their edit times
	FileArray []ProjectFileNameAndEditTime `json:"FileEditTimes"`
}

func writeDefaultConfig(pathToConfig string, config ProjectConfig) {
	json, _ := json.Marshal(config)
	os.WriteFile(pathToConfig, []byte(json), 0777)
}

func writeCJson(pathToCJson string, fjson ProjectFileJson) {
	json, _ := json.Marshal(fjson)
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

	defaultConfig := ProjectConfig{
		Version:     VERSION,
		ProjectName: projectName,
		BPPath:      fmt.Sprintf("%s/%s_Bp", pathToRoot, projectName),
		RPPath:      fmt.Sprintf("%s/%s_Rp", pathToRoot, projectName),
		TempPath:    fmt.Sprintf("%s/Pillars/Temp", pathToRoot),
	}
	// Write the default config to the ProjectConfig.json file
	writeDefaultConfig(fmt.Sprintf("%s/Pillars/ProjectConfig.json", pathToRoot), defaultConfig)

	// Write the default file edit times to the FileEditTimes.json file
	defaultFileEditTimes := ProjectFileJson{
		FileArray: []ProjectFileNameAndEditTime{},
	}

	writeCJson(fmt.Sprintf("%s/Pillars/FileEditTimes.json", pathToRoot), defaultFileEditTimes)

	return nil
}
