// @Time : 2018/6/1 19:27
// @Author : minigeek
package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerCount; i++ {
		c.createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	for _, request := range seeds {
		if isDuplicate(request.Url) {
			// log.Printf("duplicate request:%s", request.Url)
			continue
		}
		c.Scheduler.Submit(request)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				c.ItemChan <- item
			}()
		}
		// deduplicate
		for _, r := range result.Requests {
			if isDuplicate(r.Url) {
				// log.Printf("duplicate request:%s", r.Url)
				continue
			}
			c.Scheduler.Submit(r)
		}
	}
}

func (c *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			parseResult, err := c.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
