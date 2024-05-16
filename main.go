package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"github.com/hxzhouh/medium-to-md.git/convert"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var filesize int
var file string
var templateFile string

func init() {
	flag.IntVar(&filesize, "size", 5000, "file size")
	flag.StringVar(&file, "file", "", "file")
	flag.StringVar(&templateFile, "template", "template.md", "template file")
	flag.Parse()
}

func main() {
	println("hello, welcome to use medium-to-md")
	println("clean old post")
	err := os.RemoveAll("posts-md")
	log.Println(err)
	if len(file) == 0 {
		pattern := "medium-export*.zip"
		file, err = FindZipFile(pattern)
		if err != nil {
			log.Fatalf("find zip file error: %v", err)
		}
	}
	fmt.Printf("unzip file: %s\n", file)
	files, err := UnzipToMemory(file)
	if err != nil {
		log.Fatalf("unzip error: %v", err)
	}
	posts := make([]*convert.Post, 0)
	for name, content := range files {
		if strings.HasPrefix(name, "posts/") && len(content) > filesize {
			post, err := convert.Convert(name, content)
			if err != nil {
				log.Fatalf("convert error: %v", err)
			}
			posts = append(posts, post)
		}
	}
	tmpl, err := template.New("template.md").ParseFiles(templateFile)
	for _, v := range posts {
		err = os.MkdirAll(filepath.Dir(v.FileName), os.ModePerm)
		if err != nil {
			log.Fatalf("create file folder error: %v", err)
		}
		f, err := os.Create(v.FileName)
		defer func() {
			if f != nil {
				_ = f.Close()
			}
		}()
		if err != nil {
			log.Println("create file: ", v.FileName)
		}
		err = tmpl.ExecuteTemplate(f, templateFile, v)
		if err != nil {
			log.Fatalf("wirte file: %v", err)
		}
	}
}

// UnzipToMemory 解压缩 zip 文件到内存
func UnzipToMemory(src string) (map[string][]byte, error) {
	result := make(map[string][]byte)

	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}

		rc, err := f.Open()
		if err != nil {
			return nil, err
		}

		var buf bytes.Buffer
		if _, err := io.Copy(&buf, rc); err != nil {
			rc.Close()
			return nil, err
		}
		rc.Close()

		result[f.Name] = buf.Bytes()
	}
	return result, nil
}

// FindZipFile 查找符合模式的 zip 文件
func FindZipFile(pattern string) (string, error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", fmt.Errorf("未找到匹配的 zip 文件")
	}

	return files[0], nil
}
