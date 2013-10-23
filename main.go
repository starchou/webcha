package main

import (
	"github.com/astaxie/beego"
	"github.com/starchou/webcha/controllers"
)

func main() {
	beego.Info("come in")
	beego.Router("/", &controllers.ImageController{})
	beego.Router("/webchat", &controllers.MainController{})
	beego.HttpPort = 8099
	beego.Run()
}

/*
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/wechat", handlerWeixin)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func handlerWeixin(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {

	}
	if req.Method == "POST" {

	}
}
func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	file, _ := os.Open("head.jpg")
	defer file.Close()
	state, _ := file.Stat()
	b := make([]byte, state.Size())
	file.Read(b)
	w.Write(b)
}
*/
