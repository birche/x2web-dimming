package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

const lcdBrightness = "/sys/class/backlight/backlight_bempc/brightness"

//const lcdActualBrightness = "/sys/class/backlight/backlight_bempc/actual_brightness"
const configFile = "/etc/beijer/misc.conf"
const maxBrightness = 100

func changeIniFileBrightness(value int) error {
	cfg, err := ini.Load(configFile)
	if err != nil {
		fmt.Println("Failed to read config file.")
		return err
	}
	fmt.Println("Ini file Brightness: ", cfg.Section("General").Key("brightness"))
	fmt.Println("Storing new value ", value)
	cfg.Section("General").Key("brightness").SetValue(strconv.Itoa(value))
	cfg.SaveTo(configFile)
	return nil
}

func setBrightness(value string) error {
	brightness, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Failed to convert to int")
		return err
	}
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
	err = changeIniFileBrightness(brightness)
	if err != nil {
		fmt.Println("Failed to store to ini file")
		return err
	}
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
		fmt.Printf("New brightness level: %s\n", brightness)
		setBrightness(brightness)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/backlightsvc", brightnessHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
