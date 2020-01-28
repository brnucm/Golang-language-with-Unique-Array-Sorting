package main
  
import ("fmt"
"sort"
)
//int dizi tipinde bi unque adında int dizi tipinde parametre alan bi fonksiyon oluşturulmuş
func unique(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{} 
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}
  
func main() {
	//Bir dizi oluşturuldu
    dizi := []int{0,5,5,8,9,49,9,25,2,1,1,61,59,59,77,77,21,49,0,5,8,1,1}
	fmt.Println(dizi) 
	//Unique olan fonksiyon kullanılıyor bu fonksiyon bir dizi parametresi alıyor ve kendiside bir dizi
	uniquedizi := unique(dizi)
	sort.Ints(uniquedizi)
    fmt.Println(uniquedizi)
}