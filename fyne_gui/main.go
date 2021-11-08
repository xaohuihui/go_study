package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"math/rand"
	"reflect"
	"time"
)

// author: songyanhui
// datetime: 2021/11/5 11:23:02
// software: GoLand

func Shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func LoadImages() (fyne.CanvasObject, []fyne.CanvasObject, []int, int) {

	Images := make([]fyne.CanvasObject, 9)
	// 位置数组
	PositionArray := []int{1, 2, 3, 4, 5, 6, 7, 8}
	Shuffle(PositionArray)

	PositionArray = append(PositionArray, 0)
	for i, v := range PositionArray {
		img := canvas.NewImageFromFile(fmt.Sprintf("./images/%d.jpg", v))
		img.FillMode = canvas.ImageFillContain
		Images[i] = img
	}

	img9 := canvas.NewImageFromFile("./images/9.jpg")
	img9.FillMode = canvas.ImageFillContain
	return img9, Images, PositionArray, 8
}

func Show(w fyne.Window, images []fyne.CanvasObject) {
	container := fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithColumns(3),
		images...)
	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewGridWrapLayout(fyne.NewSize(800, 460)),
		container,
	),
	)
}

func ExchangePosition(NilIndex int, keyName string, Images []fyne.CanvasObject, PositionArray []int) int {
	i, j := NilIndex/3, NilIndex%3
	if keyName == "Up" {
		// 交换位置
		Images[i*3+j], Images[(i+1)*3+j] = Images[(i+1)*3+j], Images[i*3+j]
		PositionArray[i*3+j], PositionArray[(i+1)*3+j] = PositionArray[(i+1)*3+j], PositionArray[i*3+j]
		NilIndex = (i+1)*3 + j
	} else if keyName == "Down" {
		// 交换位置
		Images[i*3+j], Images[(i-1)*3+j] = Images[(i-1)*3+j], Images[i*3+j]
		PositionArray[i*3+j], PositionArray[(i-1)*3+j] = PositionArray[(i-1)*3+j], PositionArray[i*3+j]
		NilIndex = (i-1)*3 + j
	} else if keyName == "Left" {
		// 交换位置
		Images[i*3+j], Images[i*3+(j+1)] = Images[i*3+(j+1)], Images[i*3+j]
		PositionArray[i*3+j], PositionArray[i*3+(j+1)] = PositionArray[i*3+(j+1)], PositionArray[i*3+j]
		NilIndex = i*3 + j + 1
	} else if keyName == "Right" {
		// 交换位置
		Images[i*3+j], Images[i*3+(j-1)] = Images[i*3+(j-1)], Images[i*3+j]
		PositionArray[i*3+j], PositionArray[i*3+(j-1)] = PositionArray[i*3+(j-1)], PositionArray[i*3+j]
		NilIndex = i*3 + j - 1
	}
	return NilIndex
}



func main() {
	a := app.New()
	w := a.NewWindow("拼图")
	success := []int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	img9, Images, PositionArray, NilIndex := LoadImages()
	f := func(bool) {w.Close()}
	if reflect.DeepEqual(success, PositionArray) {
		Images[len(Images)-1] = img9
		Show(w, Images)
		dialog.ShowConfirm("success", "You're very smart!", f, w)
	} else {
		Show(w, Images)
		w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
			if k.Name == "Up" {
				i := NilIndex / 3
				// 移动空白图片下面的图片，若下放存在图片
				if i < 2 {
					// TODO 添加空白图片与上方图片交换逻辑
					NilIndex = ExchangePosition(NilIndex, string(k.Name), Images, PositionArray)
					Show(w, Images)
				}
			} else if k.Name == "Down" {
				i := NilIndex / 3
				// 移动空白图片上面的图片, 若上方存在图片
				if i > 0 {
					// TODO 添加空白图片与上方图片交换逻辑
					NilIndex = ExchangePosition(NilIndex, string(k.Name), Images, PositionArray)
					Show(w, Images)
				}
			} else if k.Name == "Left" {
				j := NilIndex % 3
				// 移动空白图右面的图片， 若右面存在图片
				if j < 2 {
					// TODO 添加空白图片与上方图片交换逻辑
					NilIndex = ExchangePosition(NilIndex, string(k.Name), Images, PositionArray)
					Show(w, Images)
				}
			} else if k.Name == "Right" {
				j := NilIndex % 3
				// 移动空白图左面的图片，若左面存在图片
				if j > 0 {
					// TODO 添加空白图片与上方图片交换逻辑
					NilIndex = ExchangePosition(NilIndex, string(k.Name), Images, PositionArray)
					Show(w, Images)
				}
			}

			if reflect.DeepEqual(success, PositionArray) {
				Images[len(Images)-1] = img9
				Show(w, Images)
				dialog.ShowConfirm("success", "You're very smart!", f, w)
			}
		})
	}

	w.ShowAndRun()
}
