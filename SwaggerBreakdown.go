package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

// BreakdownSwagger extracts values from the "paths" key and its first-level children
func BreakdownSwagger(data interface{}, filename string, fileDir string) {
	swaggersFilePath := filepath.Join(fileDir, "swaggers.md")
	file, err := os.Create(swaggersFilePath)
	if err != nil {
		fmt.Errorf("Error creating file:", err)
		return
	}
	defer file.Close()

	if rootMap, ok := data.(map[string]interface{}); ok {
		if paths, exists := rootMap["paths"].(map[string]interface{}); exists {
			for path, value := range paths {
				fmt.Printf("Path: %s\n", path)
				if childMap, ok := value.(map[string]interface{}); ok {
					for key := range childMap {
						fmt.Printf("  Child: %s\n", key)
						file.WriteString(fmt.Sprintf("{%% openapi src=\"%s\" path=\"%s\" method=\"%s\" expanded=\"true\" %%}\n", filename, path, key))
						file.WriteString("{% endopenapi %}\n\n")
					}
				}
			}
		} else {
			fmt.Errorf("No 'paths' key found in the file.")
		}
	} else {
		fmt.Errorf("Invalid file structure.")
	}
}

// ProcessFiles traverses the root directory and processes all openapi.yaml or openapi.json files
func ProcessFiles(root string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Errorf("Error accessing file:", err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if info.Name() == "openapi.yaml" || info.Name() == "openapi.json" {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Errorf("Error reading file %s: %v\n", path, err)
				return nil
			}

			var parsedData interface{}
			ext := filepath.Ext(path)
			if ext == ".json" {
				err = json.Unmarshal(data, &parsedData)
			} else if ext == ".yaml" || ext == ".yml" {
				err = yaml.Unmarshal(data, &parsedData)
			} else {
				return nil
			}

			if err != nil {
				fmt.Printf("Error parsing file %s: %v\n", path, err)
				return nil
			}

			fileDir := filepath.Dir(path)
			BreakdownSwagger(parsedData, info.Name(), fileDir)
		}
		return nil
	})
}

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Error getting current directory:", err)
		return
	}

	ProcessFiles(rootDir)
}
