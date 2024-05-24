package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func formatFunctionName(fileName string) string {
	// Remove the file extension
	name := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	// Split by hyphen and capitalize each part
	parts := strings.Split(name, "-")
	caser := cases.Title(language.AmericanEnglish)
	for i, part := range parts {
		parts[i] = caser.String(part)
	}
	return strings.Join(parts, "")
}

func extractSVGContent(svgFile string) (string, error) {
	data, err := os.ReadFile(svgFile)
	if err != nil {
		return "", err
	}
	// Convert to string
	content := string(data)
	// Use regex to extract content between <svg> tags
	re := regexp.MustCompile(`(?s)<svg[^>]*>(.*?)</svg>`)
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1]), nil
	}
	return "", fmt.Errorf("no svg content found in %s", svgFile)
}

func createTemplFunction(fileName, svgContent string) string {
	functionName := formatFunctionName(fileName)
	templFunction := fmt.Sprintf("templ %s() {\n@SVGWrapper() {\n  %s\n}\n}", functionName, svgContent)
	return templFunction
}

func main() {
	folderPath := "./ui/static/lucide"       // Change this to your folder path
	outputFile := "./ui/lucide/lucide.templ" // Change this to your desired output file

	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var sb strings.Builder

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".svg" {
			svgFile := filepath.Join(folderPath, file.Name())
			svgContent, err := extractSVGContent(svgFile)
			if err != nil {
				fmt.Println("Error extracting SVG content:", err)
				continue
			}
			templFunction := createTemplFunction(file.Name(), svgContent)
			sb.WriteString(templFunction + "\n\n")
		}
	}

	err = appendToFile(outputFile, sb.String())
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	// Create the command to run `go fmt` on the specified file
	cmd := exec.Command("templ", "fmt", outputFile)

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running go fmt: %s\n", err)
		return
	}

	// If there's no error, the file has been formatted and saved
	fmt.Println("File formatted successfully!")
}

func appendToFile(fileName, text string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}
