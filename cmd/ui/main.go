package main

import (
	"image"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetImageFromFile(filename string) image.Image {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func main() {

	imageFile := "./output/chapter13.png"

	a := app.New()
	w := a.NewWindow("Hello")

	img := canvas.NewImageFromImage(GetImageFromFile(imageFile))
	img.SetMinSize(fyne.Size{Width: 1000, Height: 1000})
	// img.SetMinSize(fyne.Size{img.Size().Width, img.Size().Height})

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	main := container.NewVBox(
		img,
		widget.NewButton("Render", func() {
			log.Print("Trigger Render...")
		}),
	)

	content := container.NewBorder(toolbar, nil, nil, nil, main)
	w.SetContent(content)
	w.Show()
	a.Run()

}
