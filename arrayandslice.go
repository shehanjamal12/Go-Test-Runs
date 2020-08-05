
import(
	"fmt"
)
func main(){
	 var names [3]string //method 1 to create
	games:= make([]string, 3)// method 2 to create

	games[0]="jason"
	games[1]="nadun"
	games[2]="leo"

	fmt.Printf("%v",names)
	fmt.Printf("%v",games)


	num:=[...]int {1,2,3} //the 3 dots mean  u have created an array just large enough to hold the data
	fmt.Printf("%v",num)
	
	//difference when intializing slice and array

	numbers:=[3]int{1,2,3} //specified size
	numbers2:=[]int{1,2,3} //size not spcified	
	fmt.Printf("%v",numbers)
	fmt.Printf("%v",numbers2)
	
	//slice elemnts
	a:=[]int{1,2,3,4,5,6,7,8,9}
	
	b:=a[:]//slice all
	c:=a[3:]//from after 4th to end
	d:=a[:4]//from start to 3rd
	e:=a[3:6]//from 4th to 6th
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
