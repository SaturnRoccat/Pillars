package customcomponents

/*
The ComponentInjester Handles the loading of custom components and parsing of the data
*/

import (
	"encoding/json"
	"os"

	"github.com/duckos-Mods/Pillars/Pillars/PC"
)

//"encoding/json"

func ifComponentContainsKey[CompType interface{}](jsonOp *map[string]interface{}, key string, ifContainsFunc func(CompType)) {
	if jsonOp == nil {
		return
	}
	if data, ok := (*jsonOp)[key].(CompType); ok {
		ifContainsFunc(data)
	}
	return
}

func loadCustomComponent(path string) (Component, error) {
	var comp Component
	// open the file json file
	jsonFile, err := os.Open(path)
	if err != nil {
		return comp, err
	}
	// Read the json file into an interface
	var jsonOp map[string]interface{}
	err = json.NewDecoder(jsonFile).Decode(&jsonOp)
	if err != nil {
		return comp, err
	}
	// Read the data we know exists
	comp.Name = jsonOp["name"].(string)
	comp.IntentedUsePlaces = jsonOp["intentedUsePlaces"].([]string)
	comp.Description = jsonOp["description"].(string)

	// Read the data we don't know exists
	ifComponentContainsKey[[]map[string]string](&jsonOp, "ATable", func(data []map[string]string) {
		comp.Atable = make(map[string][]map[string]string)
		for _, v := range data {
			comp.Atable[v["name"]] = append(comp.Atable[v["name"]], v)
		}
	})
	ifComponentContainsKey[map[string]string](&jsonOp, "ITable", func(data map[string]string) {
		comp.ITable = data
	})
	ifComponentContainsKey[map[string]interface{}](&jsonOp, "componentGroupData", func(data map[string]interface{}) {
		comp.ComponentGroupData = make(map[string]SectionData)
		for k, v := range data {
			comp.ComponentGroupData[k] = SectionData{JsonData: v.(map[string]interface{})}
		}
	})
	ifComponentContainsKey[map[string]interface{}](&jsonOp, "permutationData", func(data map[string]interface{}) {
		comp.PermutationData = make(map[string]SectionData)
		for k, v := range data {
			comp.PermutationData[k] = SectionData{JsonData: v.(map[string]interface{})}
		}
	})
	ifComponentContainsKey[map[string]interface{}](&jsonOp, "eventData", func(data map[string]interface{}) {
		comp.EventData = make(map[string]SectionData)
		for k, v := range data {
			comp.EventData[k] = SectionData{JsonData: v.(map[string]interface{})}
		}
	})
	ifComponentContainsKey[map[string]interface{}](&jsonOp, "descriptionData", func(data map[string]interface{}) {
		comp.DescriptionData = make(map[string]SectionData)
		for k, v := range data {
			comp.DescriptionData[k] = SectionData{JsonData: v.(map[string]interface{})}
		}
	})

	// Close the file
	err = jsonFile.Close()
	if err != nil {
		return comp, err
	}
	return comp, nil
}

func LoadCustomComponents(path string) (map[string]Component, error) {
	var components map[string]Component
	// Get all files in the dir
	var files = PC.GetFilesInDirWithExt(path, ".json")
	// Resize the slice to the number of files
	components = make(map[string]Component, len(files))
	// Loop through all files
	for _, file := range files {
		// Load the component
		comp, err := loadCustomComponent(file)
		if err != nil {
			return components, err
		}
		// Add the component to the slice
		components[comp.Name] = comp
	}
	return components, nil
}
