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
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: sampleRate,
			LanguageCode:    language,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})

	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
	}

	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
}
