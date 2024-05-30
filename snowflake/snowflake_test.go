package snowflake_test

import (
	"sync"
	"testing"

	"github.com/zan8in/pins/snowflake"
)

func TestSnowflake(t *testing.T) {
	// TODO
	nodeID := int64(1)
	sf, err := snowflake.NewSnowflake(nodeID)
	if err != nil {
		t.Logf("Error: %v\n", err)
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Logf("ID: %d\n", sf.NextID())
		}()
	}

	wg.Wait()

}
