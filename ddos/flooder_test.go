package ddos

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type FlooderTestSuite struct {
	suite.Suite
	flooder Flooder
}

type MockClient struct {
	mock.Mock
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (suite *FlooderTestSuite) SetupTest() {
	suite.flooder = NewFlooder("http://example.com", 1, 1)
}

func (suite *FlooderTestSuite) TestSetWorkerAmount() {
	assert := assert.New(suite.T())

	amountField := reflect.ValueOf(suite.flooder).Elem().FieldByName("workerAmount")
	assert.False(amountField.IsZero(), "Unable to get workerAmount field")

	assert.Equal(uint64(1), amountField.Uint())

	suite.flooder.SetWorkerAmount(5)

	assert.Equal(uint64(5), amountField.Uint())
}

func (suite *FlooderTestSuite) TestSetDuration() {
	start := time.Now()

	suite.flooder.SetDuration(2)

	suite.flooder.Start()

	elapsed := time.Since(start)

	assert.InDelta(suite.T(), 2*time.Second, elapsed, float64(1000*time.Millisecond))
}

func (suite *FlooderTestSuite) TestSendRequest() {
	mockClient := new(MockClient)
	flooder := suite.flooder.(*flooder)
	flooder.client = mockClient

	mockClient.On("Do", mock.Anything).Maybe().Return(&http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("OK")),
	}, nil)

	mockClient.AssertExpectations(suite.T())
}

func TestFlooderTestSuite(t *testing.T) {
	suite.Run(t, new(FlooderTestSuite))
}
