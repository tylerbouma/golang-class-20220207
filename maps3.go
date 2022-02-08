package main

import "fmt"

func main() {
	var fileEx = map[string]string{
		"Golang":     ".go",
		"Python":     ".py",
		"Java":       ".java",
		"C++":        ".cpp",
		"C":          ".c",
		"Ansible":    ".yml",
		"Terraform":  ".tf",
		"Javascript": ".js",
		"C#":         ".cs",
	}

	fmt.Println(fileEx)

	delete(fileEx, "Javascript")

	fileEx["Bash"] = ".sh"

	fmt.Println(fileEx)
}
