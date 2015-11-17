package getplan

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func GetInfo() (string, []string) {
	we, err := http.Get(`http://loiv.torun.pl/index.php/pl/dla-uczniow/organizacja-zajec/plan-lekcji`)
	if err != nil {
		log.Fatal(err)
	}
	replan := regexp.MustCompile(`attachments.*plan_KLAS.pdf`)
	rezast := regexp.MustCompile(`attachments.*ast.*pstwa.*pdf`)
	//ourw := make([]byte, 0, 100)
	ourw, werr := ioutil.ReadAll(we.Body)
	//wn, werr := we.Body.Read(ourp)
	if werr != nil {
		log.Fatal(perr)
	}
	wplan := replan.FindAll(ourw, -1)
	wzast := rezast.FindAll(ourw, -1)
	if len(wplan) > 0 {
		log.Fatal(wplan)
		fmt.Println("Too long!")
	}
	oz := make([]string, 0, 10)
	for i := range wzast {
		oz = append(oz, wzast[i])
	}
	return string(wplan[0]), oz
}

func GetPlan(u string) []byte {
	we, err := http.Get(`http://loiv.torun.pl/` + url.QueryEscape(u))
	if err != nil {
		log.Fatal(err)
	}
	naszplik, err := ioutil.ReadAll(we.Body)
	if err != nil {
		log.Fatal(err)
	}
	return naszplik
}

func SaveZasts(ul []string) {
	//out := make([][]byte, 0, len(ul))
	for i := range ul {
		//out = append(out, GetPlan(ul[i]))
		ZapiszPlik(GetPlan(ul[i]), ul[i])
	}
	//return out
}

//func SavePlan

func FindByte(s string, char byte) int {
	bs := []byte(s)
	last := 0
	for i := range bs {
		if bs[i] == char {
			last = i
		}
	}
	return last
}

func DajNazwePliku(s string) string {
	last := FindByte(s, byte(`/`))
	return string([]byte(s)[:last])
}

func ZapiszPlik(cozap []byte, nazwa string) {
	if _, err := os.Stat(nazwa); err == nil {
		if bytes.Equal(cozap, ioutil.ReadAll(os.Open(nazwa))) {
			return
		}
	}
	ioutil.WriteFile(nazwa, cozap, 0644)
	return
}
