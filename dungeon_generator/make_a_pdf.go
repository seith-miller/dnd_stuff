package main
import (
	"fmt"
	"log"
	"github.com/signintech/gopdf"
	"math/rand"
	"time"
)

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	roomType := []string{"clear ", "key ", "passive ", "danger ", "alt ", "block "}
	fmt.Println(roomType)

	fmt.Println(random(0, 6))

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })  
	pdf.AddPage()
	// err := pdf.AddTTFFont("DroidSansFallbackFull", "/usr/share/fonts/truetype/droid/DroidSansFallbackFull.ttf")
	err := pdf.AddTTFFont("Sixty", "./ttf/Sixty.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	// err = pdf.SetFont("DroidSansFallbackFull", "", 14)
	err = pdf.SetFont("Sixty", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	
	offSetA := 20.0
	// fmt.Println(offSetA)
	offSetB := 140.0
	//fmt.Println(offSetB)
	offSetMid := (offSetB - offSetA)/2 + offSetA
	//fmt.Println(offSetMid)

	pdf.SetY(offSetA)
	pdf.SetX(offSetA)
	pdf.Cell(nil, roomType[random(0, 6)])
	pdf.SetX(offSetMid)
	pdf.Cell(nil, roomType[random(0, 6)])
	pdf.SetX(offSetB)
	pdf.Cell(nil, roomType[random(0, 6)])

	pdf.SetY(offSetMid)
	pdf.SetX(offSetA)
	pdf.Cell(nil, roomType[random(0, 6)])
	pdf.SetX(offSetMid)
	pdf.Cell(nil, roomType[random(0, 6)])
	pdf.SetX(offSetB)
	pdf.Cell(nil, roomType[random(0, 6)])

	pdf.SetY(offSetB)
	pdf.SetX(offSetA)
	pdf.Cell(nil, roomType[random(0, 6)])
	pdf.SetX(offSetMid)
	pdf.Cell(nil, roomType[random(0, 6)])
	pdf.SetX(offSetB)
	pdf.Cell(nil, roomType[random(0, 6)])

	pdf.WritePdf("hello.pdf")
}