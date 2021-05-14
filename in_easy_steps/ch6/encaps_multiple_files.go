//https://play.golang.org/p/vNWu9fM8Kos

package main

import (
	"fmt"
	"play.ground/cube"
)

func main() {
	box := cube.Dims{}
	//var box cube.Dims 
	
	box.SetSize(2,4,6)
	fmt.Println("Footprint:", box.GetArea())
	fmt.Println("Volume:", box.GetVolume())
  fmt.Println()
//	Lower case not exported
//	fmt.Println("Width:", box.width) // does not work
//	fmt.Println("Area:", box.getArea()) //does not work
//	fmt.Println("Area:", box.GetArea()) // does work
	
	
}
-- go.mod --
module play.ground
-- cube/cube.go --
package cube

type Dims struct {
	width, length, height int
}

func (d *Dims) area() int {
	return d.width * d.length
}

//SetSize exportable
func (d *Dims) SetSize( w, l, h int) {
	d.width = w
	d.length = l
	d.height = h
}

//GetVolume exportable
func (d *Dims) GetVolume() int {
	return d.width * d.length * d.height
}

//Get Area exportable.
func (d *Dims) GetArea() int {
	return d.area()
}

/*
output prior to uncommenting lines 15 & 16
Footprint: 8
Volume: 48

uncomment lines 15 & 16
get errors 
*/
