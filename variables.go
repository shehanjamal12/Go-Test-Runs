
import(
	"fmt"
	"strconv"
)
//different ways to type variables
var(
	i int //method 1
)
func main(){
	var i int //method 2
	i=0
	
	j:=1	//method 3
	
	fmt.Printf("%v , %T\n",j,j)
	fmt.Printf("%v , %T\n",i,i)
	
	//variable experiment

	a:="shehan"
	b:=2

	fmt.Printf("%v , %T\n",a,a)
	fmt.Printf("%v , %T\n",b,b)

	//variable conversion
	var c int =  42
	var d string 
	d = strconv.Itoa(c)

	fmt.Printf("%v , %T\n",c,c)
	fmt.Printf("%v , %T\n",d,d)

}