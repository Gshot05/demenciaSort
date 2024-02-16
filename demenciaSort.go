package main 
 
import ( 
 "fmt" 
 "math/rand" 
 "sort" 
 "time" 
) 
 
func demenciaSort(arr []int) []int { 
 // Создаем копию массива для работы 
 arrCopy := make([]int, len(arr)) 
 copy(arrCopy, arr) 
 
 // Итерируемся до тех пор, пока копия массива не станет отсортированной 
 for !sort.IntsAreSorted(arrCopy) { 
  // Рандомно выбираем индекс элемента для удаления 
  indexToRemove := rand.Intn(len(arrCopy)) 
 
  // Удаляем выбранный элемент из массива 
  arrCopy = append(arrCopy[:indexToRemove], arrCopy[indexToRemove+1:]...) 
 
  // Перемешиваем массив 
  rand.Seed(time.Now().UnixNano()) 
  rand.Shuffle(len(arrCopy), func(i, j int) { 
   arrCopy[i], arrCopy[j] = arrCopy[j], arrCopy[i] 
  }) 
 } 
 
 return arrCopy 
} 
 
func main() { 
 // Пример использования функции 
 array := []int{4, 2, 7, 1, 9, 5} 
 
 fmt.Println("Исходный массив:", array) 
 
 newArray := demenciaSort(array) 
 
 fmt.Println("Отсортированный массив:", newArray) 
}
