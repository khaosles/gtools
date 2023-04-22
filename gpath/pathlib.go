package gpath

import (
	"os"
	"path/filepath"
	"strings"
)

// Exist judge whether exists filepath
func Exist(path string) bool {
	path = Format(path)
	// path stat
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// Format the path
func Format(path string) string {
	// delete the space at the both ends
	path = strings.TrimSpace(path)
	// simplified path
	path = filepath.Clean(path)
	// \\ to /
	path = filepath.ToSlash(path)
	// / to \ or /
	path = filepath.FromSlash(path)
	return path
}

// IsFile judge whether is a file
func IsFile(path string) bool {
	path = Format(path)
	// path stat
	fileStat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileStat.IsDir()
}

// IsDir judge whether is a dir
func IsDir(path string) bool {
	path = Format(path)
	// path stat
	fileStat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileStat.IsDir()
}

// FileSize obtain file size
func FileSize(path string) int64 {
	path = Format(path)
	if !IsFile(path) {
		panic("FileNotExist: '" + path + "' not exist")
	}
	fileStat, err := os.Stat(path)
	if err != nil {
		panic("FileOpenError: cannot open the file" + path)
	}
	return fileStat.Size()
}

// Basename get filename
func Basename(path string) string {
	path = Format(path)
	return filepath.Base(path)
}

// Join the path
func Join(elem ...string) string {
	return Format(filepath.Join(elem...))
}

// Dirname get file dir name
func Dirname(path string) string {
	path = Format(path)
	return filepath.Dir(path)
}

// Split get file dir name
func Split(path string) (string, string) {
	path = Format(path)
	return filepath.Split(path)
}

// Suffix get file suffix
func Suffix(path string) string {
	path = Format(path)
	return filepath.Ext(path)
}

// Mkdir create a folder
func Mkdir(path string) {
	path = Format(path)
	// if path is a file, raise an error
	if IsFile(path) {
		panic("FolderCreateError: path is a file, " + path)
	}
	if !IsDir(path) {
		// create the folder
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic("FolderCreateError: " + err.Error())
		}
	}
}

// MkParentDir create a path's parent dir
func MkParentDir(path string) {
	Mkdir(Dirname(path))
}

// Abs get file absolute path
func Abs(path string) string {
	path = Format(path)
	if filepath.IsAbs(path) {
		return path
	} else {
		absPath, err := filepath.Abs(path)
		if err != nil {
			panic("AbsPathGetError" + err.Error())
		}
		return absPath
	}
}

// Remove file or folder
func Remove(path string) {
	if !Exist(path) {
		return
	}
	err := os.RemoveAll(path)
	if err != nil {
		panic("RemoveFileError: " + err.Error())
	}
}

// RemoveFile remove a file
func RemoveFile(path string) {
	if !IsFile(path) {
		return
	}
	err := os.Remove(path)
	if err != nil {
		panic("RemoveFileError: " + err.Error())
	}
}

func RootPath() string {
	rootPath, _ := os.Getwd()
	return rootPath
}
