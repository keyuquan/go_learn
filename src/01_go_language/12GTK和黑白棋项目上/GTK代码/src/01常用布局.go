package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main0101() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("D:/GoCode/src/UI.glade")

	win := gtk.WindowFromObject(builder.GetObject("window1"))
	win.SetSizeRequest(480, 320)

	//获取水平布局
	hbox := gtk.HBoxFromObject(builder.GetObject("hbox1"))
	//创建按钮
	btn := gtk.NewButtonWithLabel("点击进入")
	btn.Clicked(func() { fmt.Println("点击") })
	//将按钮添加到布局中
	hbox.Add(btn)

	win.ShowAll()

	//点击关闭按钮退出程序
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	gtk.Main()

}

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("D:/GoCode/src/UI1.glade")

	win := gtk.WindowFromObject(builder.GetObject("window1"))

	win.SetSizeRequest(480, 320)

	//垂直布局
	vbox := gtk.VBoxFromObject(builder.GetObject("vbox1"))

	//表格布局
	//gtk.TableFromObject(builder.GetObject("table1"))
	lab := gtk.NewLabel("XX社区")
	vbox.Add(lab)

	win.ShowAll()
	//点击关闭按钮退出程序
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	gtk.Main()
}
