package pxcanvas

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"zerotomastery.io/pixl/apptype"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer    *PxCanvasRenderer
	PixelData   image.Image
	mouseState  PxCanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.X)
	y0 := int(pxCanvas.CanvasOffset.Y)
	x1 := int(pxCanvas.PxCols*pxCanvas.PxSize + x0)
	y1 := int(pxCanvas.PxRows*pxCanvas.PxSize + y0)
	return image.Rect(x0, y0, x1, y1)
}

func InBounds(position fyne.Position, bounds image.Rectangle) bool {
	if position.X >= float32(bounds.Min.X) && position.X < float32(bounds.Max.X) &&
		position.Y >= float32(bounds.Min.Y) && position.Y < float32(bounds.Max.Y) {
		return true
	}
	return false
}

func NewBlankImage(cols, rows int, color color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, color)
		}
	}
	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       state,
	}
	pxCanvas.PixelData = NewBlankImage(pxCanvas.PxCols, pxCanvas.PxRows, color.NRGBA{128, 128, 128, 128})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}

	renderer := &PxCanvasRenderer{
		pxCanvas:     pxCanvas,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}
	pxCanvas.renderer = renderer
	return renderer
}

func (pxCanvas *PxCanvas) TryPan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	if previousCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		pxCanvas.Pan(*previousCoord, ev.PointEvent)
	}
}

// Brushable interface implementation
func (pxCanvas *PxCanvas) SetColor(c color.Color, x, y int) {
	if nrgba, ok := pxCanvas.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, c)
	}
	if rgba, ok := pxCanvas.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, c)
	}
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := pxCanvas.Bounds()
	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}
	pxSize := float32(pxCanvas.PxSize)
	xOffset := pxCanvas.CanvasOffset.X
	yOffset := pxCanvas.CanvasOffset.Y
	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y
}

func (pxCanvas *PxCanvas) LoadImage(img image.Image) {
	dimensions := img.Bounds()
	pxCanvas.PxCanvasConfig.PxCols = dimensions.Dx()
	pxCanvas.PxCanvasConfig.PxRows = dimensions.Dy()

	pxCanvas.PixelData = img
	pxCanvas.reloadImage = true
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) NewDrawing(cols, rows int) {
	pxCanvas.appState.SetFilePath("")
	pxCanvas.PxCols = cols
	pxCanvas.PxRows = rows
	pixelData := NewBlankImage(cols, rows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.LoadImage(pixelData)
}
