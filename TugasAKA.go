package main

import (
	"fmt"
	"path/filepath"
	"io/ioutil"
	"time"
)

// Recursive folder traversal
func traverseFolderRecursive(path string, indent string) (int, error) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("Gagal membuka folder: %v", err)
	}

	count := 0
	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			count++ // Hitung folder
			fmt.Printf("%s[DIR]  %s\n", indent, entry.Name())
			subCount, err := traverseFolderRecursive(fullPath, indent+"    ")
			if err != nil {
				return 0, err
			}
			count += subCount
		} else {
			fmt.Printf("%s[FILE] %s\n", indent, entry.Name())
		}
	}
	return count, nil
}

// Iterative folder traversal
func traverseFolderIterative(path string) (int, error) {
	stack := []struct {
		path  string
		indent string
	}{ {path, ""} }

	count := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		entries, err := ioutil.ReadDir(current.path)
		if err != nil {
			return 0, fmt.Errorf("Gagal membuka folder: %v", err)
		}

		for _, entry := range entries {
			fullPath := filepath.Join(current.path, entry.Name())
			if entry.IsDir() {
				count++ // Hitung folder
				fmt.Printf("%s[DIR]  %s\n", current.indent, entry.Name())
				stack = append(stack, struct {
					path  string
					indent string
				}{fullPath, current.indent + "    "})
			} else {
				fmt.Printf("%s[FILE] %s\n", current.indent, entry.Name())
			}
		}
	}
	return count, nil
}

func main() {
	rootPath := "."

	// Measure recursive traversal time
	startRecursive := time.Now()
	fmt.Println("=== Isi Folder dari Path (Recursive) ===")
	fmt.Println("-----------------------------------------")
	folderCountRecursive, err := traverseFolderRecursive(rootPath, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	elapsedRecursive := time.Since(startRecursive)

	// Measure iterative traversal time
	startIterative := time.Now()
	fmt.Println("\n=== Isi Folder dari Path (Iterative) ===")
	fmt.Println("-----------------------------------------")
	folderCountIterative, err := traverseFolderIterative(rootPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	elapsedIterative := time.Since(startIterative)

	// Print performance results in a table format
	fmt.Println("\n+----+-------------------------+-------------------------+")
	fmt.Println("| n  |    Recursive Time (s)   |    Iterative Time (s)   |")
	fmt.Println("+----+-------------------------+-------------------------+")
	fmt.Printf("| %2d | %23.15e | %23.15e |\n", folderCountRecursive, elapsedRecursive.Seconds(), elapsedIterative.Seconds())
	fmt.Printf("| %2d | %23.15e | %23.15e |\n", folderCountIterative, elapsedRecursive.Seconds(), elapsedIterative.Seconds())
	fmt.Println("+----+-------------------------+-------------------------+")
}