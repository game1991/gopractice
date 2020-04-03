package main

import (
	"fmt"
	"image"
	"os"
)

//在很多场景下我们需要确保高并发的场景下只执行一次，例如只加载一次配置文件，只关闭一次通道等

var icons map[string]image.Image

func loadIcon(name string) image.Image {
	name = fmt.Sprintf("./%s.png", name)
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return nil
}

func loadIcons() { //map初始化
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}
