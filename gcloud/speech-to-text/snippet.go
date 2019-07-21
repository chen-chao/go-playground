// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2019-06-20 13:56:18

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	"github.com/cryptix/wav"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// example from google cloud

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("Usage: go run snippet .wav language-code")
	}
	filename := flag.Arg(0)
	language := flag.Arg(1)

	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data length:", len(data))

	file, _ := os.Open(filepath.Clean(filename))
	defer file.Close()

	stat, _ := os.Stat(filename)
	reader, err := wav.NewReader(file, stat.Size())

	sampleRate := int32(reader.GetSampleRate())
	fmt.Println(reader.String())

	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:                   speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz:            sampleRate,
			LanguageCode:               language,
			EnableWordTimeOffsets:      true,
			EnableAutomaticPunctuation: true,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})

	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
	}

	out, err := os.Create("results.time.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	for i, result := range resp.Results {
		fmt.Fprintf(out, "%s\n", strings.Repeat("-", 20))
		fmt.Fprintf(out, "Result %d\n", i+1)
		for j, alt := range result.Alternatives {
			fmt.Fprintf(out, "Alternative %d: \"%v\" (confidence=%3f)\n", j+1, alt.Transcript, alt.Confidence)
			for _, w := range alt.Words {
				fmt.Fprintf(out,
					"Word: \"%v\" (startTime=%3f, endTime=%3f)\n",
					w.Word,
					float64(w.StartTime.Seconds)+float64(w.StartTime.Nanos)*1e-9,
					float64(w.EndTime.Seconds)+float64(w.EndTime.Nanos)*1e-9)
			}
		}
	}
}
