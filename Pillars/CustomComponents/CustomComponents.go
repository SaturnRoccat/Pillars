package customcomponents

import (
	"fmt"
	"regexp"
)

type SectionData struct {
	// The json data for the section
	JsonData map[string]interface{} // We use interface{} because we don't know what the data will be fully. And we will have to at run time parse
}

type ComponentArguments struct {
	Name string `json:"argName"`
	Type string `json:"argType"`
	Val  interface{}
}

type Component struct {
	// The name of the component
	Name string

	// Name of the lua functino to execute when the component is created
	OnCreate string

	// Atable will store all variables that are used in the component
	Atable []ComponentArguments

	// // ITable will store all included json files
	// ITable map[string]string

	// Component Groups
	ComponentGroupData SectionData

	// Permutation data
	PermutationData SectionData

	// Event data
	EventData SectionData

	// Description Data
	DescriptionData SectionData

	// Description
	Description string

	// Intented use places
	IntentedUsePlaces []string
}

func parseString(str string, regExp string, replaceVal interface{}, typeOfVal string) interface{} {
	compiledRegExp := regexp.MustCompile(regExp)
	if len(str) == len(regExp) {
		// If the length of the string and the regexp are the same we have to do some extra work
		switch typeOfVal {
		case "string":
			return replaceVal.(string)
		case "int":
			return replaceVal.(int)
		case "float":
			return replaceVal.(float64)
		case "bool":
			return replaceVal.(bool)
		case "any":
			return replaceVal
		case "jsonObject":
			return replaceVal.(map[string]interface{})
		case "jsonArray":
			return replaceVal.([]interface{})
		default:
			panic("Unknown type")
		}
	} else {
		str = compiledRegExp.ReplaceAllString(str, replaceVal.(string))
		return str
	}
}

func parseOutArray(arr []interface{}, regExp string, replaceVal interface{}, typeOfVal string) []interface{} {
	var returnArr []interface{} = make([]interface{}, len(arr))
	for i, v := range arr {
		// Check if the type of the value is a string we only can replace strings
		if val, ok := v.(string); ok {
			// Replace val
			var rval = parseString(val, regExp, replaceVal, typeOfVal)
			returnArr[i] = rval
		} else if val, ok := v.(map[string]interface{}); ok {
			// If the value is a map[string]interface{} then we need to recurse
			val = parseOutTable(val, regExp, replaceVal, typeOfVal)
			returnArr[i] = val
		} else if val, ok := v.([]interface{}); ok {
			// If the value is a []interface{} then we need to recurse
			val = parseOutArray(val, regExp, replaceVal, typeOfVal)
			returnArr[i] = val
		}
	}
	return returnArr
}

func parseOutTable(table map[string]interface{}, regExp string, replaceVal interface{}, typeOfVal string) map[string]interface{} {
	var returnTable map[string]interface{} = make(map[string]interface{})
	for k, v := range table {
		// Check if the type of the value is a string we only can replace strings
		if val, ok := v.(string); ok {
			var rval = parseString(val, regExp, replaceVal, typeOfVal)
			returnTable[k] = rval
		} else if val, ok := v.(map[string]interface{}); ok {
			// If the value is a map[string]interface{} then we need to recurse
			val = parseOutTable(val, regExp, replaceVal, typeOfVal)
			returnTable[k] = val
		} else if val, ok := v.([]interface{}); ok {
			// If the value is a []interface{} then we need to recurse
			val = parseOutArray(val, regExp, replaceVal, typeOfVal)
			returnTable[k] = val
		}
	}
	return returnTable

}

func (C *Component) handleArguments(args []ComponentArguments) {
	for _, v := range args {
		// We are gonna use a regex fight me its the simplest way that i know with out looping through every character
		// We are gonna assume that we dont have to modify the keys and only the values

		C.ComponentGroupData.JsonData = parseOutTable(C.ComponentGroupData.JsonData, v.Name, v.Val, v.Type)
		C.PermutationData.JsonData = parseOutTable(C.PermutationData.JsonData, v.Name, v.Val, v.Type)
		C.EventData.JsonData = parseOutTable(C.EventData.JsonData, v.Name, v.Val, v.Type)
		C.DescriptionData.JsonData = parseOutTable(C.DescriptionData.JsonData, v.Name, v.Val, v.Type)
	}
}

// Im lazy and dont wanna do this rn
func (C *Component) DecodeArgs(argsToDecode interface{}) []ComponentArguments {
	var args []ComponentArguments = make([]ComponentArguments, 0)
	return args
}

func (C *Component) handleNestedCustomComponents(otherComponents *map[string]Component) {
	for k, v := range C.ComponentGroupData.JsonData {
		if _, ok := (*otherComponents)[k]; ok {
			var args = C.DecodeArgs(v)
			var otherComp = (*otherComponents)[k]
			var compData, err = otherComp.Build(args, otherComponents)
			if err != nil {
				// panic(err) // We can comment out because we just are gonna have a stripping phase that removes all components that have errors and throws a warning
				return
			}

			// We need to merge the data
			for k, v := range compData {
				switch k {
				case "componentGroupData":
					for k, v := range v {
						C.ComponentGroupData.JsonData[k] = v
					}
				case "permutationData":
					for k, v := range v {
						C.PermutationData.JsonData[k] = v
					}
				case "eventData":
					for k, v := range v {

						C.EventData.JsonData[k] = v
					}
				case "descriptionData":
					for k, v := range v {
						C.DescriptionData.JsonData[k] = v
					}
				default:
					println(fmt.Sprintf("Unknown key %s. This shouldnt ever trigger", k))
				}
			}
		}
	}
}

func (C *Component) Build(args []ComponentArguments, otherComponents *map[string]Component) (map[string]map[string]interface{}, error) {
	var returnJsons map[string]map[string]interface{} = make(map[string]map[string]interface{}, 0)
	// Handle the arguments
	C.handleArguments(args)
	C.handleNestedCustomComponents(otherComponents)
	return returnJsons, nil
}
