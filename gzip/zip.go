package gzip

import (
	"archive/zip"
	"compress/flate"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/khaosles/gtools/gpath"
)

/*
   @File: zip.go
   @Author: khaosles
   @Time: 2023/5/30 00:12
   @Desc:
*/

// Compress 压缩文件
func Compress(zipFile string, files []string) {

	gpath.MkParentDir(zipFile)
	newZipFile, err := os.Create(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// 设置压缩级别为"zip.Deflate"
	zipWriter.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		info, err := f.Stat()
		if err != nil {
			log.Fatal(err)
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Fatal(err)
		}
		header.Name = gpath.Basename(file)

		// 使用压缩级别"zip.Deflate"进行压缩
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(writer, f)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// Decompress 解压文件
func Decompress(zipFile string, dest string) {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), os.ModePerm)
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			_, err = io.Copy(f, rc)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
