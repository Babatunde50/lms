package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Babatunde50/lms/internal/models"
	setting "github.com/Babatunde50/lms/pkg/settings"
	"github.com/Babatunde50/lms/router"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {

	// router setup
	routersInit := router.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

}
