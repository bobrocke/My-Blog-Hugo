package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
    title := os.Args[1]
    if title == "" {
        fmt.Println("Please provide a title for the blog post.")
        return
    }

    filename := generateFilename(title)
    content := generateContent(title)

    err := createFile(filename, content)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }

    fmt.Println("Blog post created successfully:", filename)
}

func generateFilename(title string) string {
    now := time.Now()
    date := now.Format("2006-01-02")
    slug := strings.ReplaceAll(strings.ToLower(title), " ", "-")
    return fmt.Sprintf("%s-%s.md", date, slug)
}

func generateContent(title string) string {
	return fmt.Sprintf(`---
title: %s
date: %s
---

Write your blog content here.
`, title, time.Now().Format("2006-01-02"))
}

func createFile(filename string, content string) error {
    dir := "posts"
    err := os.MkdirAll(dir, 0755)
    if err != nil {
        return err
    }

    filepath := filepath.Join(dir, filename)
    file, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(content)
    return err
}

go run new_post.go "Your Blog Post Title"
