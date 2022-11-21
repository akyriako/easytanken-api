package main

import (
	"fmt"
	"github.com/akyriako/easytanken-api/api/v1/handlers"
	"github.com/akyriako/easytanken-api/internal"
	"k8s.io/klog/v2"
	"net/http"
)

var (
	Configuration *internal.Config
)

func main() {
	defer exit()

	handlers.ApiKey = Configuration.ApiKey

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/stations", handlers.GetStationsInProximity)
	http.HandleFunc("/stations/", handlers.GetStationById)

	klog.Infof("start listening at:" + Configuration.Port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", Configuration.Port), nil)
	if err != nil {
		klog.Errorf("start listening at:%s failed!", Configuration.Port)
		klog.Fatalln(err)
	}
}

func init() {
	klog.InitFlags(nil)

	configuration, err := internal.GetConfig()
	if err != nil {
		klog.Fatalln(err)
	}

	Configuration = configuration
}

func exit() {
	exitCode := 10
	klog.Infoln("stopped api listeners")
	klog.FlushAndExit(klog.ExitFlushTimeout, exitCode)
}
