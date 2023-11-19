package PC

type ProjectConfig struct {
	Version     []int  `json:"version"`
	ProjectName string `json:"projectName"`
	BPPath      string `json:"bpPath"`
	RPPath      string `json:"rpPath"`
	TempPath    string `json:"tempPath"`
}

type ProjectFileJson struct {
	// We need to store an array of file names and their edit times
	FileArray map[string]int64 `json:"fileArray"`
}
