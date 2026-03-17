package main

import (
	"encoding/json"
	"fmt"
)

type Image struct {
	URl    string `json:"url"`
	Width  int    `jsonL:"width"`
	Height int    `json:"height"`
}

func main() {
	imageData := `[
		{
			"url":"abc",
			"width": 1920,	
			"height": 1080	
		},
		{
			"url":"abcd",
			"width": 2250,	
			"height": 1500	
		},
		{
			"url":"abce",
			"width": 3840,	
			"height": 2880	
		},
		{
			"url":"abcf",
			"width": 1000,	
			"height": 1000	
		}
	]`

	var images []Image

	err := json.Unmarshal([]byte(imageData), &images)
	if err != nil {
		fmt.Printf("Error while unmarshaling!")
	}

	// err := json.Unmarshal([]byte(imageData), &images)
	// if err != nil {
	// 	fmt.Printf("Error in unmarshelling!")
	// }

	for _, img := range images {
		ratio := float64(img.Width) / float64(img.Height)
		fmt.Printf("URL: %s, Ratio: %.2f\n", img.URl, ratio)
	}

	for _, img := range images {
		g := gcd(img.Width, img.Height)
		fmt.Printf(
			"URL: %s, Ratio: %d:%d\n",
			img.URl,
			img.Width/g,
			img.Height/g,
		)
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
