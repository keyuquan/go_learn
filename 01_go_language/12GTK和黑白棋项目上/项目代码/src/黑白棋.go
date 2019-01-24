package main

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

type ChessWidget struct {
	Win      *gtk.Window //窗体
	BlackImg *gtk.Image  //黑棋图片
	WhiteImg *gtk.Image  //白棋图片

	MinBtn   *gtk.Button //最小化按钮
	CloseBtn *gtk.Button //关闭按钮

	BlackScore *gtk.Label //黑棋分数
	WhiteScore *gtk.Label //白棋分数

	TimerLab *gtk.Label //计时器

	w, h int //窗体大小

	startx, starty int //棋盘起始坐标
	endx, endy     int //棋盘结束坐标

	gridw, gridh int //棋盘格子宽高

	board   [8][8]int //存储棋盘棋子
	curRole int       //执棋角色
}

//枚举
const (
	Empty = iota //0
	Black        //1
	White        //2

)

//棋盘初始化方法
func (obj *ChessWidget) InitChess() {
	//初始化数据
	obj.startx = 210
	obj.starty = 70
	obj.endx = 620
	obj.endy = 480
	//棋盘格子大小
	obj.gridw = 52
	obj.gridh = 51
	//初始化棋盘棋子
	obj.board[3][3] = Black
	obj.board[4][4] = Black

	obj.board[3][4] = White
	obj.board[4][3] = White

	//黑方执棋
	obj.curRole = Black

}

//为棋盘对象绑定方法
func (obj *ChessWidget) CreateWidGet() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("D:/GoCode/src/UI.glade")

	//加载窗体信息
	obj.Win = gtk.WindowFromObject(builder.GetObject("Win"))
	//设置窗体大小不可变
	obj.Win.SetResizable(false)
	//设置窗体无边框
	obj.Win.SetDecorated(false)
	//获取窗体大小
	obj.Win.GetSize(&obj.w, &obj.h)

	//绘图事件
	obj.Win.SetAppPaintable(true)
	obj.Win.Connect("expose-event", DrawWindowImagefromFile, obj)

	//允许鼠标事件
	obj.Win.SetEvents(int(gdk.BUTTON_PRESS_MASK))
	obj.Win.Connect("button-press-event", MouseClickEvent, obj)

	//最小化按钮
	obj.MinBtn = gtk.ButtonFromObject(builder.GetObject("BtnMin"))
	//设置按钮图片
	DrawButtonImageFromFile(obj.MinBtn, "D:/GoCode/src/image/min.png")
	//设置按钮点击事件
	obj.MinBtn.Connect("clicked", BtnMinClick, obj)

	//关闭按钮
	obj.CloseBtn = gtk.ButtonFromObject(builder.GetObject("BtnClose"))
	//设置按钮图片
	DrawButtonImageFromFile(obj.CloseBtn, "D:/GoCode/src/image/close.png")
	//设置按钮点击事件
	obj.CloseBtn.Connect("clicked", BtnCloseClick)

	//设置图片
	obj.BlackImg = gtk.ImageFromObject(builder.GetObject("Img1"))
	DrawImageFromFile(obj.BlackImg, "D:/GoCode/src/image/black.png")

	//设置图片
	obj.WhiteImg = gtk.ImageFromObject(builder.GetObject("Img2"))
	DrawImageFromFile(obj.WhiteImg, "D:/GoCode/src/image/white.png")

	//设置标签
	obj.TimerLab = gtk.LabelFromObject(builder.GetObject("Timer"))
	obj.TimerLab.ModifyFontSize(16)

	obj.BlackScore = gtk.LabelFromObject(builder.GetObject("BlackScore"))
	obj.BlackScore.ModifyFontSize(40)

	obj.WhiteScore = gtk.LabelFromObject(builder.GetObject("WhiteScore"))
	obj.WhiteScore.ModifyFontSize(40)

	//显示窗体信息
	obj.Win.ShowAll()
	gtk.Main()
}

func main() {

	var obj ChessWidget
	//初始化方法
	obj.InitChess()
	//创建棋盘
	obj.CreateWidGet()
}
