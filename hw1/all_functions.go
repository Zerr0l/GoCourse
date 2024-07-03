package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Task 1
func helloWorld() {
	fmt.Println("Hello world!")
}

// Task 2
func sumOfTwo(a, b int) int {
	return a + b
}

// Task 3
func oddOrEven() {
	fmt.Print("Введите число для проверки на чётность: ")
	var x int
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println(err)
		return
	}
	if x%2 == 0 {
		fmt.Printf("Число %d чётное\n", x)
	} else {
		fmt.Printf("Числа %d нечётное\n", x)
	}
}

// Task 4
func maxOfThree(a, b, c int) int {
	switch {
	case a > b && a > c:
		return a
	case b > c:
		return b
	default:
		return c
	}
}

// Task 5
func getFactorial() {
	fmt.Print("Введите число, для которого необходимо посчитать факториал: ")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	fmt.Printf("%d! = %d\n", n, res)
}

// Task 6
func isVowel() {
	fmt.Print("Введите символ: ")
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println(err)
		return
	}
	s = strings.ToLower(s)
	c := s[0]
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'y' {
		fmt.Println("Это гласная")
	} else {
		fmt.Println("Это согласная")
	}
}

// Task 7
func writeAllPrime() {
	fmt.Print("Введите число, до которого будет произведён поиск простых: ")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			if i != 2 {
				fmt.Print(", ")
			}
			fmt.Print(i)
			for j := 2 * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	fmt.Println()
}

// Task 8
func reverseString(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

// Task 9
func sumOfElementsInArray(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

// Task 10
type Rectangle struct {
	Height, Width int
}

func (r *Rectangle) area() int {
	return r.Height * r.Width
}

// Task 11
func temperatureConvertor(t float32) float32 {
	return t*9/5 + 32
}

// Task 12
func countdown() {
	fmt.Print("Введите число: ")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := n; i >= 1; i-- {
		if i != n {
			fmt.Print(" ")
		}
		fmt.Print(i)
	}
	fmt.Println()
}

// Task 13
func myLength(s string) int {
	res := 0
	for range s {
		res++
	}
	return res
}

// Task 14
func findInArray() {
	fmt.Print("Введите элементы массива через пробел: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	arr := strings.Split(s, " ")
	fmt.Print("Введите число, которое необходимо найти: ")
	var x int
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, a := range arr {
		aInt, _ := strconv.Atoi(a)
		if x == aInt {
			fmt.Println("Число содержится в массиве")
			return
		}
	}
	fmt.Println("Число не содержится в массиве")
}

// Task 15
func getAverage(arr []int) float64 {
	res := 0.0
	for _, x := range arr {
		res += float64(x) / float64(len(arr))
	}
	return res
}

// Task 16
func writeMultiplicationTable() {
	fmt.Print("Введите число, до которого будет выведена таблица умножения: ")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("\t", i*j)
		}
		fmt.Println()
	}
}

// Task 17
func isPalindrome(s string) bool {
	for i := 0; i < (len(s)-1)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Task 18
func findMinAndMax(arr []int) (int, int) {
	mn, mx := arr[0], arr[0]
	for _, x := range arr {
		if x < mn {
			mn = x
		}
		if x > mx {
			mx = x
		}
	}
	return mn, mx
}

// Task 19
func deleteElement[T any](arr *[]T, ind int) {
	*arr = append((*arr)[:ind], (*arr)[ind+1:]...)
}

// Task 20
func linearSearch(arr []int, x int) int {
	for i, a := range arr {
		if a == x {
			return i
		}
	}
	return -1
}

// Task 21
func deleteDuplicates(arr *[]int) {
	sort.Ints(*arr)
	res := []int{(*arr)[0]}
	for _, x := range *arr {
		if x != res[len(res)-1] {
			res = append(res, x)
		}
	}
	*arr = res
}

// Task 22
func bubbleSort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

// Task 23
func generateFibonacci(n int) []int {
	res := make([]int, n)
	if n == 1 {
		res[0] = 1
		return res
	}
	res[0] = 1
	res[1] = 1
	for i := 2; i < n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res
}

// Task 24
func countNumberOfOccurrences[T comparable](arr []T, x T) int {
	res := 0
	for _, a := range arr {
		if a == x {
			res++
		}
	}
	return res
}

// Task 25
func intersectionOfArrays(arr1, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			res = append(res, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

// Task 26
func isAnagram(s1, s2 string) bool {
	var cnt1, cnt2 [256]int
	for _, ch := range s1 {
		cnt1[ch]++
	}
	for _, ch := range s2 {
		cnt2[ch]++
	}
	return cnt1 == cnt2
}

// Task 27
func mergeArrays(arr1, arr2 []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			res = append(res, arr1[i])
			i++
		} else if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		res = append(res, arr1[i])
		i++
	}
	for j < len(arr2) {
		res = append(res, arr2[j])
		j++
	}
	return res
}

// Task 28
// Я не понял, как реализовать хеш-таблицу с generic типом, так что напишу её для int-ов
const P = 257

type HashTableElement struct {
	key   string
	value int
}

type HashTable struct {
	table           [][]HashTableElement
	numberOfBuckets int
	size            int
}

func getHashOfString(s string, module int) int {
	res := 0
	for _, ch := range s {
		res = (res*P + int(ch)) % module
	}
	return res
}

func (h *HashTable) add(key string, value int) {
	hash := getHashOfString(key, h.numberOfBuckets)
	h.size++
	h.table[hash] = append(h.table[hash], HashTableElement{key, value})
}

func (h *HashTable) get(key string) int {
	hash := getHashOfString(key, h.numberOfBuckets)
	for _, element := range h.table[hash] {
		if element.key == key {
			return element.value
		}
	}
	return -1
}

func (h *HashTable) delete(key string) {
	hash := getHashOfString(key, h.numberOfBuckets)
	for i, element := range h.table[hash] {
		if element.key == key {
			deleteElement(&h.table[hash], i)
		}
	}
}

func (h *HashTable) getSize() int {
	return h.size
}

func (h *HashTable) toDefault() {
	h.size = 0
	h.numberOfBuckets = 1
	h.table = make([][]HashTableElement, h.numberOfBuckets)
}

func (h *HashTable) setNumberOfBuckets(n int) {
	h.size = 0
	h.numberOfBuckets = n
	h.table = make([][]HashTableElement, h.numberOfBuckets)
}

// Task 29
func binarySearch(arr []int, x int) int {
	l, r := 0, len(arr)
	for r-l > 1 {
		m := (l + r) / 2
		if arr[m] > x {
			r = m
		} else {
			l = m
		}
	}
	if arr[l] == x {
		return l
	}
	return -1
}

// Task 30
type Queue[T any] struct {
	begin, end []T
}

func (q *Queue[T]) push(element T) {
	q.begin = append(q.begin, element)
}

type ReturnValue struct {
	value interface{}
	err   error
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func (q *Queue[T]) back() ReturnValue {
	if len(q.end) != 0 {
		return ReturnValue{q.end[len(q.end)-1], nil}
	}
	if len(q.begin) == 0 {
		var t T
		return ReturnValue{t, &MyError{
			time.Now(),
			"cannot return an element, queue is empty",
		}}
	}
	var tmp []T
	for len(q.begin) > 0 {
		tmp = append(tmp, q.begin[len(q.begin)-1])
		q.begin = q.begin[:len(q.begin)-1]
	}
	for len(tmp) > 0 {
		q.end = append(q.end, tmp[len(tmp)-1])
		tmp = tmp[:len(tmp)-1]
	}
	return ReturnValue{q.end[len(q.end)-1], nil}
}

func (q *Queue[T]) pop() {
	if len(q.end) != 0 {
		q.end = q.end[:len(q.end)-1]
		return
	}
	if len(q.begin) == 0 {
		return
	}
	var tmp []T
	for len(q.begin) > 0 {
		tmp = append(tmp, q.begin[len(q.begin)-1])
		q.begin = q.begin[:len(q.begin)-1]
	}
	for len(tmp) > 0 {
		q.end = append(q.end, tmp[len(tmp)-1])
		tmp = tmp[:len(tmp)-1]
	}
	q.end = q.end[:len(q.end)-1]
}

func main() {
	//// Checks for tasks

	// Task 1
	helloWorld()

	// Task 2
	fmt.Println(sumOfTwo(3, 2))

	// Task 3
	oddOrEven()

	// Task 4
	fmt.Println(maxOfThree(1, 2, 3))

	// Task 5
	getFactorial()

	// Task 6
	isVowel()

	// Task 7
	writeAllPrime()

	// Task 8
	fmt.Println(reverseString("hello world"))

	// Task 9
	fmt.Println(sumOfElementsInArray([]int{1, 2, 3}))

	// Task 10
	rectangle := Rectangle{3, 6}
	fmt.Println(rectangle.area())

	// Task 11
	fmt.Println(temperatureConvertor(12))

	// Task 12
	countdown()

	// Task 13
	fmt.Println(myLength("Hello world!"))

	// Task 14
	findInArray()

	// Task 15
	fmt.Println(getAverage([]int{2, 2, 3}))

	// Task 16
	writeMultiplicationTable()

	// Task 17
	fmt.Println(isPalindrome("abacaba"))

	// Task 18
	fmt.Println(findMinAndMax([]int{2, 1, 9, 4, -1}))

	// Task 19
	arr19 := []int{1, 2, 3, 4, 5}
	deleteElement(&arr19, 1)
	fmt.Println(arr19)

	// Task 20
	fmt.Println(linearSearch([]int{1, 2, 5, 8, 9}, 7))

	// Task 21
	arr21 := []int{1, 1, 3, 3, 2, 2}
	deleteDuplicates(&arr21)
	fmt.Println(arr21)

	// Task 22
	arr22 := []int{1, 5, 3, 2, 2, 3}
	bubbleSort(&arr22)
	fmt.Println(arr22)

	// Task 23
	fmt.Println(generateFibonacci(8))

	// Task 24
	fmt.Println(countNumberOfOccurrences([]int{1, 2, 1, 2, 2, 2}, 2))

	// Task 25
	fmt.Println(intersectionOfArrays([]int{1, 3, 5, 5}, []int{2, 3, 5}))

	// Task 26
	fmt.Println(isAnagram("hello world", "world hello"))

	// Task 27
	fmt.Println(mergeArrays([]int{1, 3, 5, 5}, []int{2, 3, 5}))

	// Task 28
	ht := HashTable{}
	ht.setNumberOfBuckets(3)
	ht.add("e", 4)
	ht.add("e'", 5)
	ht.add("kf", 3)
	ht.add("kc", 6)
	fmt.Println(ht.size)
	fmt.Println(ht.get("e"))
	ht.delete("e")
	fmt.Println(ht.size)
	fmt.Println(ht.get("e"))

	// Task 29
	fmt.Println(binarySearch([]int{1, 1, 2, 2, 2, 4}, 2))

	// Task 30
	q := Queue[int]{}
	q.push(1)
	q.push(2)
	q.push(3)
	fmt.Println(q.back())
	q.pop()
	q.push(4)
	q.pop()
	q.pop()
	q.pop()
	fmt.Println(q.back().err)
}
