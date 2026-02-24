package main 
import "fmt" 

// func functionname(parametername datatype) returntype {
//  //function body
// }



func main1212112(){
	var bill  = calculateBill(10 , 23.20)
	fmt.Println("your bill is" , bill)

	var area ,perimeter  = calculateArea(5 , 23.20)

	// area, _ := rectProps(10.8, 5.6) // perimeter is discarded means when i dint want to use multiple return values
	fmt.Println("your area is" , area); 

	fmt.Println("your perimeter is" , perimeter)
}


func calculateBill(item int , price float32 ) float32 {  //single return value
	return  float32(item * int(price)) 
};


func calculateArea(item int , price float32 ) (float32 ,float32) {  //multiple return value
	var areaaa float32= float32(item * int(price)) ;
	var perriii float32 = float32(item) + price ;

	return  areaaa  ,  perriii 
}
