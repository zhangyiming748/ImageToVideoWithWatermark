package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		panic("第二个参数需要填写图片根目录")
	}
	root := os.Args[1]
	fmt.Println("root:", root)
	ext := ".jpg"
	files, err := findFilesWithExt(root, ext)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("找到的文件：")
	length := len(files)
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	index1 := seed.Intn(length)
	fmt.Printf("选中 %d 图片\n", index1)
	vname := strings.Replace(files[index1], ext, ".mp4", 1)
	//ffmpeg -loop 1 -framerate 1 -i bd%d.jpg -i watermark.png -filter_complex "[0:v]fps=25,scale=1920:1080,format=yuv420p[bg];[bg][1:v]overlay=(W-w)/2:(H-h)/2" -c:v libx264 -t 30 -y bd2.mp4
	cmd := exec.Command("ffmpeg", "-loop", "1", "-framerate", "1", "-i", files[index1], "-i", "/mnt/c/Users/zen/Github/ImageToVideoWithWatermark/watermark.png", "-filter_complex", "[0:v]fps=25,scale=1920:1080,format=yuv420p[bg];[bg][1:v]overlay=(W-w)/2:(H-h)/2", "-c:v", "libx264", "-t", "30", vname)
	execCommand(cmd)
}
func execCommand(c *exec.Cmd) {
	fmt.Println(c.String())
	output, err := c.CombinedOutput()
	if err != nil {
		return
	} else {
		fmt.Println(string(output))
	}
}
func findFilesWithExt(dir, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
