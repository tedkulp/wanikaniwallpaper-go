package main

import "flag"

var (
	outputFile    = flag.String("output", "out.png", "path to ouput file")
	orderFilename = flag.String("orderfile", "data/order", "path to order file")
	key           = flag.String("key", "", "API key")
	xmargin       = flag.Int("xmargin", 0, "margin on left and right edges")
	ymargin       = flag.Int("ymargin", 0, "margin on top and bottom edges")
)

func main() {
	flag.Parse()

	kanjiList := GetKanjiForApiKey(*key)
	order := NewOrder(*orderFilename)
	order.Update(kanjiList)

	renderer := NewRenderer()

	Draw(order, renderer, *width, *height, *xmargin, *ymargin)

	renderer.SaveImage(*outputFile)
}
