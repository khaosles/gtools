package g

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/hashicorp/go-getter"
	"github.com/khaosles/gtools/gpath"
)

/*
   @File: Download.go
   @Author: khaosles
   @Time: 2023/5/16 21:46
   @Desc:
*/

type Http struct {
}

func (h Http) DownloadOne(src, dst string) {
	// 创建目标文件夹
	gpath.MkParentDir(dst)
	// http
	httpGetter := getter.HttpGetter{
		Header: http.Header{
			"Authorization": []string{fmt.Sprintf("Bearer %s", "aHRoajphSFJvYW5oQWFHOTBiV0ZwYkM1amIyMD06MTYyMzkxNDI0Njo4ZTgyZGNiNmRlMjVhODJjYWExYjA3NTc5Mzc2ZjMzY2FkOTk1NjU0")},
		},
		ReadTimeout: 5 * time.Second,
	}
	// 设置下载参数
	client := &getter.Client{
		Src:  src, // 文件地址
		Dst:  dst, // 本地存储路径
		Mode: getter.ClientModeFile,
		Options: []getter.ClientOption{
			getter.WithGetters(map[string]getter.Getter{
				"httpGetter": &httpGetter,
			}),
		},
	}
	if err := client.Get(); err != nil {
		log.Fatal(err)
	}
}

func Down(src, dst string) {
	client := grab.NewClient()
	//client.HTTPClient.Transport.DisableCompression = true
	req, _ := grab.NewRequest(dst, src)
	req.Filename = "a.hdf"
	// ...
	req.NoResume = true
	req.HTTPRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", "aHRoajphSFJvYW5oQWFHOTBiV0ZwYkM1amIyMD06MTYyMzkxNDI0Njo4ZTgyZGNiNmRlMjVhODJjYWExYjA3NTc5Mzc2ZjMzY2FkOTk1NjU0"))

	resp := client.Do(req)
	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			fmt.Printf("%.02f%% complete\n", resp.Progress())

		case <-resp.Done:
			if err := resp.Err(); err != nil {
				// ...
			}

			// ...
			return
		}
	}
}
