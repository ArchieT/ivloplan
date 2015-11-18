package getplan

import (
	"bytes"
	"fmt"
	//	"io"
	"io/ioutil"
	"log"
	"net/http"
	//	"net/url"
	"os"
	"regexp"
	//	"strings"
)

var Replan = regexp.MustCompile(`attachments.{6,26}plan_KLAS.pdf`)
var Rezast = regexp.MustCompile(`attachments.{6,26}ast.{0,6}pstwa.{0,15}pdf`)

func mergebytesandgivestring(b [][]byte) string {
	var c bytes.Buffer
	for _, j := range b {
		c.Write(j)
	}
	return c.String()
}

func GetLinks() []string {
	we, err := http.Get(`http://loiv.torun.pl/index.php/pl/dla-uczniow/organizacja-zajec/plan-lekcji`)
	if err != nil {
		log.Fatal(err)
	}
	repdf := regexp.MustCompile(`attachments.{10,50}\.pdf`)
	ourw, werr := ioutil.ReadAll(we.Body)
	if werr != nil {
		log.Fatal(err)
	}
	wpdf := repdf.FindAll(ourw, -1)
	log.Println("Got link: ", mergebytesandgivestring(wpdf))
	if len(wpdf) > 1 {
		log.Println("More than one plan!")
	}
	od := make([]string, 0, 10)
	for i := range wpdf {
		od = append(od, string(wpdf[i]))
	}
	return od
}

func GetFile(u string) []byte {
	we, err := http.Get(`http://loiv.torun.pl/` + u)
	if err != nil {
		log.Fatal(err)
	}
	naszplik, err := ioutil.ReadAll(we.Body)
	if err != nil {
		log.Fatal(err)
	}
	return naszplik
}

func SaveFiles(ul []string) {
	for i := range ul {
		Save(GetFile(ul[i]), ul[i])
	}
}

func findByte(s string, char byte) int {
	bs := []byte(s)
	last := 0
	for i := range bs {
		if bs[i] == char {
			last = i
			fmt.Println(string(bs[i+1:]))
		}
	}
	return last + 1
}

func DajNazwePliku(s string) string {
	last := findByte(s, []byte(`/`)[0])
	a := string([]byte(s)[last:])
	fmt.Println(a)
	return a
}

func Save(cozap []byte, url string) {
	nazwa := DajNazwePliku(url)
	fileopen, err := os.OpenFile(nazwa, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fileread, err := ioutil.ReadAll(fileopen)
	if err != nil {
		log.Fatal(err)
	}
	if bytes.Equal(cozap, fileread) {
		log.Println("same", nazwa)
		return
	}
	if len(fileread) > 100 {
		Save(cozap, url+`_new.pdf`)
		return
	}
	n, err := fileopen.Write(cozap)
	if err != nil {
		log.Fatal(err)
	}
	if n < len(cozap) {
		log.Fatal("Za malo: ", n, " zamiast ", len(cozap))
	}
	return
}

func FullService() {
	SaveFiles(GetLinks())
}
