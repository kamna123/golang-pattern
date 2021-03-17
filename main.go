package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*https://dev.to/nadirbasalamah/golang-tutorial-10-concurrency-with-channel-54co*/

func fibonacciGenerator(n int) chan int {
	ch := make(chan int)
	go func() {
		//fmt.Println("inside fun")
		a := 1
		ch <- a
		//fmt.Println("a :", a)
		b := 1
		ch <- b
		c := 2
		for i := 2; i < n; i++ {
			c = a + b
			a = b
			b = c
			//	fmt.Println("c :", c)
			ch <- c
		}
		close(ch)
	}()
	return ch
}
func Count(start int, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			// Blocks on the operation
			ch <- i
		}

		close(ch)
	}(ch)

	return ch
}
func main() {
	// fmt.Println("No bottles of beer on the wall")

	// for i := range Count(1, 9) {
	// 	fmt.Println("Pass it around, put one up,", i, "bottles of beer on the wall")
	// 	// Pass it around, put one up, 1 bottles of beer on the wall
	// 	// Pass it around, put one up, 2 bottles of beer on the wall
	// 	// ...
	// 	// Pass it around, put one up, 99 bottles of beer on the wall
	// }

	// fmt.Println(100, "bottles of beer on the wall")
	for i := range fibonacciGenerator(20) {
		fmt.Print(" ", i)
	}

	count := make(chan int, 10)

	for i := 0; i < 100; i++ {
		go DownloadFile1(i, count)
	}
	for i := 0; i < 100; i++ {

		//<-(count - 1)
		//	fmt.Println("counter is ", count-1)

		//time.Sleep(1 * time.Second)

	}
	// ctr := int32(0)

	// var wg sync.WaitGroup
	// for i := 0; i < 10000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		atomic.AddInt32(&ctr, 1)
	// 	}()
	// }
	// wg.Wait()

	// time.Sleep(5 * time.Second)

	// fmt.Printf("ctr=%+v\n", ctr)

	// a := [6]string{
	// 	"https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
	// 	"http://www.africau.edu/images/default/sample.pdf",
	// 	"https://file-examples-com.github.io/uploads/2017/10/file-sample_150kB.pdf",
	// 	"https://file-examples-com.github.io/uploads/2017/10/file-example_PDF_500_kB.pdf",
	// 	"https://file-examples-com.github.io/uploads/2017/10/file-example_PDF_1MB.pdf",
	// 	"http://www.pdf995.com/samples/pdf.pdf"}
	//	fmt.Println(a)
	// How it will work with unbuffered?

	// urlCount := make(chan string)
	// for i := 0; i < len(a); i++ {
	// 	fileName := strings.SplitN(a[i], "/", -1)
	// 	fmt.Println(fileName[len(fileName)-1])

	// 	go DownloadFile(i, a[i], fileName[len(fileName)-1], urlCount)
	// 	//fmt.Println("calling fun download file " + fileName[len(fileName)-1])
	// }

	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(<-urlCount)
	// }

	//time.Sleep(1 * time.Second)
}

func DownloadFile1(ctr int, urlCount chan int) error {
	//	fmt.Println("abc")
	//var mutex = &sync.Mutex{}

	if true {
		//	time.Sleep(5*time.Second + time.Duration(ctr)*time.Second)
		//	fmt.Printf("file Nme in Download file %v\n", ctr)

		urlCount <- (ctr + 1) % 10
		//atomic.AddInt32(urlCount, 1)

		return nil
	}
	return nil
}
func DownloadFile(ctr int, url string, fileName string, urlCount chan string) error {
	//	fmt.Println("abc")
	fmt.Println("file Nme in Download file " + fileName)
	if true {
		time.Sleep(5*time.Second + time.Duration(ctr)*time.Second)

		urlCount <- fileName
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("non 200")
	}
	file, errFile := os.Create(fileName)
	if errFile != nil {
		return errFile
	}
	defer file.Close()
	//Write the bytes to the fiel
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	urlCount <- fileName
	return nil
}
