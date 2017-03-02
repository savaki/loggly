package loggly_test

import (
	"strings"
	"testing"
	"time"

	"github.com/savaki/loggly"
)

func TestClient(t *testing.T) {
	token := "blah"
	samples := []string{
		"hello world\n",
		"argle bargle\n",
		"glip glop\n",
	}
	var received string
	fn := func(data []byte) error {
		received = string(data)
		return nil
	}
	client := loggly.New(token, loggly.Publish(fn), loggly.Interval(time.Minute))

	// When
	for _, v := range samples {
		sample := v
		client.Write([]byte(sample))
	}
	time.Sleep(time.Millisecond * 50)
	client.Flush()
	client.Close()

	// Then
	if expected := strings.Join(samples, ""); received != expected {
		t.Errorf("expected %v; got %v", expected, received)
	}
}
