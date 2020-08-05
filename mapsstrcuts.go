import(
	"fmt"
)
type person struct{
	name string
	age int
}
func main(){
	grades := map[string]int{
		"chemistry" :96,
		"biology":98,
		"physics":93,
	}
	fmt.Println(grades)
	grades["ict"] = 66
	fmt.Println(grades)
	
	aperson:=person{
		name: "kin",
		age:52,
	}
	fmt.Println(aperson)
}	
