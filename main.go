package main

import (
	"github.com/formych/dota/config"
	"github.com/formych/dota/router"
)

func main() {
	router.Run(config.C.Port)
	// 粗糙的关闭连接，后续优化
	config.Close()
	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router.Mu,
	// }

	// go func() {
	// 	err := server.ListenAndServe()
	// 	if err == http.ErrServerClosed {
	// 		logrus.Info("Closed server under request")
	// 	} else {
	// 		logrus.Info("Server unexpected err:[%s]", err.Error())
	// 	}
	// }()
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// s := <-c
	// logrus.Infof("Recive signal: %s", s)
	// if err := server.Shutdown(context.Background()); err != nil {
	// 	logrus.Fatalln("Closed seever failed: ", err)
	// }
	// logrus.Infof("Closed server sucessfully")
}
