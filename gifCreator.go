package turtle

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"image/png"
	"log"
	"os"
	"sync"
	"time"
)

func TakeScreenshot(window Window, outputPNGPath string) {
	screenshot := window.GetCanvas().GetScreenshot()
	out, err := os.Create(outputPNGPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = png.Encode(out, screenshot)
	if err != nil {
		panic(err)
	}
}

// Start this as a go routine to create a GIF of your creation.
func CreateGif(
	window Window,
	delayBetweenScreenshots time.Duration,
	delayBetweenGifFrames time.Duration,
	outputGifFilePath string,
	frameCount int,
) {
	//
	hundredthOfSecondDelay := int(delayBetweenGifFrames.Seconds() * 100)
	canvas := window.GetCanvas()

	// Collect the images
	fmt.Printf("GIF: %s: Collecting images\n", outputGifFilePath)
	images := make([]image.Image, 0, frameCount)
	nextTime := time.Now()
	for frameIndex := 0; frameIndex < frameCount; frameIndex++ {
		screenShot := canvas.GetScreenshot()
		images = append(images, screenShot)

		nextTime = nextTime.Add(delayBetweenScreenshots)
		time.Sleep(time.Until(nextTime))
	}

	// Process the images.
	fmt.Printf("GIF: %s: Processing images\n", outputGifFilePath)
	startTime := time.Now()
	imagesPal := make([]*image.Paletted, 0, len(images))
	delays := make([]int, 0, frameCount)

	// Create 10 workers
	requestChan := make(chan palettedImageProcessorRequest, 100)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			go gifPalettedImageProcessor(wg, requestChan)
		}
	}()

	// Fill the request channel with images to convert
	for frameIndex := range images {
		screenShot := images[frameIndex]
		bounds := screenShot.Bounds()
		ssPaletted := image.NewPaletted(bounds, palette.WebSafe)
		imagesPal = append(imagesPal, ssPaletted)
		delays = append(delays, hundredthOfSecondDelay)

		// // All this additional logic to speed up the following commented lines. which takes a couple seconds per frame
		// for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		// 	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		// 		ssPaletted.Set(x, y, screenShot.At(x, y))
		// 	}
		// }

		// // calling convertToPalettedWithCacheRGBA takes 2.1s for 95 images. vs 684ms with 10 workers. 3.12 times slower
		// srcRGBA, _ := screenShot.(*image.RGBA)
		// convertToPalettedWithCacheRGBA(palettedCacheRGBA, srcRGBA, ssPaletted)

		newRequest := palettedImageProcessorRequest{src: screenShot, dest: ssPaletted}
		requestChan <- newRequest
	}
	// Close the channel and wait for all workers to finish.
	close(requestChan)
	wg.Wait()

	// Write the file
	f, _ := os.OpenFile(outputGifFilePath, os.O_WRONLY|os.O_CREATE, 0o600)
	defer f.Close()
	err := gif.EncodeAll(f, &gif.GIF{
		Image: imagesPal,
		Delay: delays,
	})
	if err != nil {
		log.Printf("Error while writing gif to file: %v\n", err)
	}

	deltaTime := time.Since(startTime)
	fmt.Printf("GIF: %s: Done: %s\n", outputGifFilePath, deltaTime)
}

// takes 0.12s on average.
func convertToPalettedWithCache(palettedCache map[color.Color]uint8, src image.Image, dest *image.Paletted) {
	// startTime := time.Now()

	if src.Bounds() != dest.Bounds() {
		log.Println("src and dest do not have the same rectangle")
		return
	}

	bounds := src.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := src.At(x, y)

			// Get the palette color index for this RGBA color
			palettedIndex, ok := palettedCache[c]
			if !ok {
				palettedIndex = uint8(dest.Palette.Index(c))
				palettedCache[c] = palettedIndex
			}
			// dest.Set(x, y, c)
			i := dest.PixOffset(x, y)
			dest.Pix[i] = palettedIndex
		}
	}

	// deltaTime := time.Since(startTime)
	// fmt.Println(deltaTime.Seconds())
}

// takes 0.065s on average.
func convertToPalettedWithCacheRGBA(palettedCache map[color.RGBA]uint8, src *image.RGBA, dest *image.Paletted) {
	// startTime := time.Now()

	if src.Bounds() != dest.Bounds() {
		log.Println("src and dest do not have the same rectangle")
		return
	}

	bounds := src.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Unroll the operations of: dest.Set(x, y, src.At(x, y))
			// first, get the src color: srcImage.At(x, y)
			i := (y-src.Rect.Min.Y)*src.Stride + (x-src.Rect.Min.X)*4
			// This appears to be called "Full slice expressions". a[low : high : max]. It sets the new slice capacity to max-low
			s := src.Pix[i : i+4 : i+4]
			c := color.RGBA{s[0], s[1], s[2], s[3]}

			// Get the palette color index for this RGBA color
			palettedIndex, ok := palettedCache[c]
			if !ok {
				palettedIndex = uint8(dest.Palette.Index(c))
				palettedCache[c] = palettedIndex
			}
			// dest.Set(x, y, c)
			i = (y-dest.Rect.Min.Y)*dest.Stride + (x-dest.Rect.Min.X)*1
			dest.Pix[i] = palettedIndex
		}
	}

	// deltaTime := time.Since(startTime)
	// fmt.Println(deltaTime.Seconds())
}

type palettedImageProcessorRequest struct {
	src  image.Image
	dest *image.Paletted
}

// Has a 2x speed improvement when the src image is an image.RGBA
// The chosen palette color for a given src color is saved in a cache.
func gifPalettedImageProcessor(
	wg *sync.WaitGroup,
	requestChan chan palettedImageProcessorRequest,
) {
	palettedCacheColor := make(map[color.Color]uint8, 100)
	palettedCacheRGBA := make(map[color.RGBA]uint8, 100)
	for {
		request, ok := <-requestChan
		if !ok {
			break
		}

		srcRGBA, ok := request.src.(*image.RGBA)
		if ok {
			convertToPalettedWithCacheRGBA(palettedCacheRGBA, srcRGBA, request.dest)
		} else {
			convertToPalettedWithCache(palettedCacheColor, request.src, request.dest)
		}
	}

	wg.Done()
}
