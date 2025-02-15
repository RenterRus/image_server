package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var port *int
var path *string

func init() {
	port = flag.Int("port", 8111, "set listen port. Default: 8111. Example: 8765")
	path = flag.String("path", "./images", "source directory to img")

	flag.Parse()

}

func main() {
	time.Sleep(time.Second * 10)
	http.Handle("/img", http.HandlerFunc(imgUploader))

	go func() {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("listen port:", *port)
		fmt.Println("path to source directory:", *path)
	}()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal(err)
	}

}

func imgUploader(w http.ResponseWriter, r *http.Request) {
	fmt.Println("query")
	fmt.Println(r.URL)

	files, err := os.ReadDir(*path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if len(files) > 0 {
		images := make([]string, 0, len(files))

		for _, v := range files {
			images = append(images, v.Name())
		}

		img := "./images/" + images[rand.Intn(len(images))]
		fileBytes, err := os.ReadFile(img)
		if err != nil {
			panic(err)
		}
		fmt.Println(img)
		fmt.Println()

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
	}
}
