package utils

import (
	"errors"
	"log"
	"math"
	"os"
	"time"
)

const Epsilon = 1e-5

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= Epsilon
}

func EnsureDir(dirName string) error {
	err := os.Mkdir(dirName, 0755)
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
