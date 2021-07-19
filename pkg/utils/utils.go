package utils

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/bricef/ray-tracer/pkg/camera"
	"github.com/bricef/ray-tracer/pkg/canvas"
	"github.com/bricef/ray-tracer/pkg/scene"
)

const Epsilon = 1e-5

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= Epsilon
}

func EnsureDir(dirName string) error {
	err := os.MkdirAll(dirName, 0755)
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

func DegressToRadians(d float64) float64 {
	return (d / 360.0) * 2 * math.Pi
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func SaveFrame(frame *canvas.ImageCanvas, c *camera.Camera, s *scene.Scene, filepath string) {
	// progress := make(chan float64)
	c.Render(s, frame)
	frame.WritePNG(filepath)
	fmt.Printf("Wrote output to %v\n", filepath)
}
