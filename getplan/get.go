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
		ZapiszPlik(GetFile(ul[i]), ul[i])
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

func ZapiszPlik(cozap []byte, url string) {
	nazwa := DajNazwePliku(url)
	_, err := os.Stat(`./` + nazwa)
	if err == nil {
		fileopen, openerr := os.Open(`./` + nazwa)
		if openerr != nil {
			log.Fatal(string(openerr.Error()))
		}
		fileread, readallerr := ioutil.ReadAll(fileopen)
		if readallerr != nil {
			log.Fatal(string(readallerr.Error()))
		}
		if bytes.Equal(cozap, fileread) {
			log.Println("same", url)
			return
		}
	} else {
		os.Create(`./` + nazwa)
		log.Println("create ", nazwa)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(`./`+nazwa, cozap, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func FullService() {
	SaveFiles(GetLinks())
}
