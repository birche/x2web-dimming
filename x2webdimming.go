package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"net/http"
	"strconv"
)

const lcdBrightness = "/sys/class/backlight/backlight_bempc/brightness"
const lcdActualBrightness = "/sys/class/backlight/backlight_bempc/actual_brightness"
const maxBrightness = 100

func setBrightness(value string) error {
	brightness, err := strconv.Atoi(value)
	if brightness > maxBrightness {
		brightness = maxBrightness
	}
	if brightness < 0 {
		brightness = 0
	}
	err = ioutil.WriteFile(lcdBrightness, []byte(strconv.Itoa(brightness)), os.ModeExclusive)
	if err != nil {
		fmt.Println("Failed to set brightness", err.Error())
		return err
	}
	fmt.Println("Did set new brightness, value ", brightness)
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s", r.URL.Path[1:])
}

func brightnessHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Backlight brightness handler")
	fmt.Println("Request URL", r.URL.Query())
	brightness := r.URL.Query().Get("brightness")
	if brightness != "" {
		fmt.Println("New brightness level: %s", brightness)
		setBrightness(brightness)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/backlightsvc", brightnessHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

