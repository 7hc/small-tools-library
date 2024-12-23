package main

import (
	"fmt"
	"image"
	_ "image/gif"  // 注册 GIF 格式识别器
	_ "image/jpeg" // 注册 JPEG 格式识别器
	_ "image/png"  // 注册 PNG 格式识别器
	"os"
	"path/filepath"

	_ "golang.org/x/image/bmp"  // 注册 BMP 格式识别器
	_ "golang.org/x/image/tiff" // 注册 TIFF 格式识别器
	_ "golang.org/x/image/webp" // 注册 WEBP 格式识别器
)

// GetImageSize 获取图片的宽度和高度
func GetImageSize(imgPath string) (width, height int, err error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return 0, 0, fmt.Errorf("Err_Open_Image: %w", err)
	}
	defer file.Close()

	src, _, err := image.Decode(file)
	if err != nil {
		return 0, 0, fmt.Errorf("Err_Decode_Image: %w", err)
	}

	bounds := src.Bounds()
	return bounds.Dx(), bounds.Dy(), nil
}

// GetFilesOnly 获取指定目录下的所有文件，不包括子目录
func GetFilesOnly(dir string) ([]string, error) {
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}
	return files, nil
}

func main() {
	dir := "./image" // 替换为你的目录路径(Image Folder Path)
	files, err := GetFilesOnly(dir)
	if err != nil {
		fmt.Printf("Err_Getting_Files: %v\n", err)
		return
	}
	for _, file := range files {
		imgPath := file // 图片路径
		width, height, err := GetImageSize(imgPath)
		if err != nil {
			fmt.Printf("Err_Getting_Image_Dimensions: %v\n", err)
			return
		}

		fmt.Printf("%s -> Width: %d, Height: %d\n", file, width, height)
	}
}
