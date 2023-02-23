/* Calculate total download size for NYC taxi data for 2020

for each month, we have two files: green and yellow
	Ex:
	https://s3.aws.com/green_tripdata.csv
	https://s3.aws.com/yellow_tripdata.csv

*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)



var ( //global vars 
	urlTemplate = "https://s3.aws.com/%s_tripdata.csv"
	colors = []string{"green", "yellow"}
)

func downloadSize(url string) (int, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)

	if err != nil {
		return 0, err 
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err 
	}

	if rsp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(rsp.Status) 
	}

	return strconv.Atoi(rsp.Header.Get("Content-Type")) //converts content type header length to a number 
}

//added to improve runtime by using goroutines 
type result struct {
	url string
	size int 
	err error
}


func sizeWorker(url string, ch chan result)  {
	fmt.Println(url)
	res := result{url: url}
	res.size, res.err = downloadSize(url)
	ch <- res 
}
func main()  {
	start := time.Now()

	size := 0 

	ch := make(chan result)
	for month := 1; month < 13; month++ {
		for _, color := range colors {
			url := fmt.Sprintf(urlTemplate,color,month) 
			// fmt.Println(url)
			// n,err := downloadSize(url)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// size += n 

			go sizeWorker(url, ch)
		}
	}

	//receiver of send 
	for i := 0; i < len(colors)*12; i++ {
		r := <-ch 
		if r.err != nil {
			log.Fatal(r.err)
		}

		size += r.size 
	}



	duration := time.Since(start)
	fmt.Println(size, duration)

}