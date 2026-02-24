package main

import "fmt"

type Income1 interface {
	calculate() float64
	source() string
}

type FixedBilling1 struct {
	projectName1  string
	biddedAmount1 float64
}

type TimeAndMaterial1 struct {
	projectName1 string
	timeinHrs    int
	ratePerHr    int
}

func (fi FixedBilling1) calculate() float64 {
	return fi.biddedAmount1
}
func (fi FixedBilling1) source() string {
	return fi.projectName1
}
func (fi TimeAndMaterial1) source() string {
	return fi.projectName1
}
func (fi TimeAndMaterial1) calculate() float64 {
	return float64(fi.timeinHrs * fi.ratePerHr)
}


func calculateSum(sa []Income1){
    netincome := 0.0

    for _ ,val := range sa{
        
       	fmt.Printf("Income From %s = $%f\n", val.source(), val.calculate())
		netincome += val.calculate()
        
    }
    fmt.Printf("Net income of organization = $%f", netincome)
}
func mainPract() {

	p1 := FixedBilling1{projectName1: "P1", biddedAmount1: 100}
	p4 := FixedBilling1{projectName1: "P4", biddedAmount1: 1090}
	p2 := TimeAndMaterial1{projectName1: "P2", timeinHrs: 10, ratePerHr: 50}
	p3 := TimeAndMaterial1{projectName1: "P3", timeinHrs: 102, ratePerHr: 4}

	va :=[]Income1{p1,p2,p3 ,p4}
    calculateSum(va)
}
