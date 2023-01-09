package routes

import (
	"fmt"
	"log"

	"github.com/donvito/hellomod"
	"github.com/tarm/serial"
	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func Test() {
	fmt.Println("Hello world1")
	value := gjson.Get(json, "name.last")
	println(value.String())
	hellomod.SayHello()
	c := &serial.Config{Name: "COM3", Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("%q", buf[:n+1])
}
