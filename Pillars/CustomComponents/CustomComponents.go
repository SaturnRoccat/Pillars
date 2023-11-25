package customcomponents

type SectionData struct {
	// The json data for the section
	JsonData map[string]interface{} // We use interface{} because we don't know what the data will be fully. And we will have to at run time parse
}

type Component struct {
	// The name of the component
	Name string

	// Name of the lua functino to execute when the component is created
	OnCreate string

	// Atable will store all variables that are used in the component
	Atable map[string][]map[string]string

	// // ITable will store all included json files
	// ITable map[string]string

	// Component Groups
	ComponentGroupData map[string]SectionData

	// Permutation data
	PermutationData map[string]SectionData

	// Event data
	EventData map[string]SectionData

	// Description Data
	DescriptionData map[string]SectionData

	// Description
	Description string

	// Intented use places
	IntentedUsePlaces []string
}

func (C *Component) Build() ([]map[string]interface{}, error) {
	var returnJsons []map[string]interface{} = make([]map[string]interface{}, 0)
	//TODO: write the code to build the component
	return returnJsons, nil
}
