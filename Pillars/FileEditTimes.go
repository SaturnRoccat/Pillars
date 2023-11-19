package Pillars

import (
	"encoding/json"
	"os"

	"github.com/duckos-Mods/Pillars/Pillars/PC"
)

func updateFileEditTimes(FETTpUpdate map[string]int64, PFJ PC.ProjectFileJson) {
	// loop through the files to update
	for file, time := range FETTpUpdate {
		// update the file edit time
		PFJ.FileArray[file] = time

	}

	// marshal the json
	json, _ := json.Marshal(PFJ)

	var FEDP = PTPR + "/Pillars/FileEditTimes.json"
	// write the json to the file
	os.WriteFile(FEDP, []byte(json), 0777)
}

func getFilesToIgnore(pathToRoot string) map[string]bool {
	println("Getting files to ignore...")
	var filesToIgnore = make(map[string]bool)
	var filesToUpdate = make(map[string]int64)

	// read in the file edit times
	var FET, _ = os.ReadFile(FETP)

	// unmarshal the json
	var FETJson PC.ProjectFileJson
	json.Unmarshal(FET, &FETJson)
	if FETJson.FileArray == nil {
		FETJson.FileArray = make(map[string]int64)
	}

	// get the files in the root dir
	var files = PC.GetFileInfoInDir(pathToRoot)

	// loop through all files in the root dir
	for file, info := range files {
		// check if the file is in the FETJson
		if _, ok := FETJson.FileArray[file]; ok {
			// check if the edit time is the same
			if FETJson.FileArray[file] == info.ModTime().Unix() {
				// if it is, add it to the filesToIgnore map
				filesToIgnore[file] = true
			} else {
				// if it isn't, add it to the filesToUpdate map
				filesToUpdate[file] = info.ModTime().Unix()
			}
		} else {
			// if it isn't, add it to the filesToUpdate map
			filesToUpdate[file] = info.ModTime().Unix()
		}
	}
	println("Done getting files to ignore.")
	println("Updating file edit times...")
	// update the file edit times
	updateFileEditTimes(filesToUpdate, FETJson)
	println("Done updating file edit times.")
	return filesToIgnore
}
