package osu

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func testDeserializing(filename string) func(*testing.T) {
	return func(t *testing.T) {
		f, err := os.Open("./test/" + filename)
		if err != nil {
			t.Errorf("failed to locate file '%s'", filename)
		}
		_, err = ParseBeatmap(f)
		if err != nil {
			t.Errorf("failed to parse file: %+v", err)
		}
	}
}

func TestParsing(t *testing.T) {
	files, err := ioutil.ReadDir("./test")
	if err != nil {
		log.Fatal(err)
	}
	if testing.Short() {
		// shuffle to get random beatmaps
		rand.Shuffle(len(files), func(i, j int) {
			files[i], files[j] = files[j], files[i]
		})
	}

	for i, file := range files {
		if !strings.HasSuffix(file.Name(), ".osu") {
			continue
		}
		if i > 2 && testing.Short() {
			break
		}

		t.Run(fmt.Sprintf("test%d", i), testDeserializing(file.Name()))
	}
}
