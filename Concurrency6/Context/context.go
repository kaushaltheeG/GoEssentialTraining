package main

import (
	"context"
	"fmt"
	"time"
)

//The context package is a built-in package which is used for deadlines and cancellations



type Bid struct {
	AdURL string
	Price float64
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)

	return Bid{
		AdURL: "https://ads.com/19",
		Price: 0.05,
	}
}

var defaultBid = Bid{
	AdURL: "https://default.com",
	Price: 0.02,
}

func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) //buffered to avoide goroutine leak 

	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch: 
		return bid 
	case <- ctx.Done(): //hits is timeout reaches before bestBids thread sleep()
		return defaultBid
	}
}

func main()  {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	url := "https:/cats.com"
	bid := findBid(ctx, url)
	fmt.Println(bid)
}