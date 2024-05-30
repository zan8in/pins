package snowflake

import (
	"fmt"
	"sync/atomic"
	"time"
)

const (
	nodeBits        = 10                        // 节点ID所占的位数
	sequenceBits    = 12                        // 序列所占的位数
	maxNodeID       = -1 ^ (-1 << nodeBits)     // 节点ID的最大值
	maxSequence     = -1 ^ (-1 << sequenceBits) // 序列的最大值
	timeShift       = nodeBits + sequenceBits   // 时间戳左移位数
	nodeShift       = sequenceBits              // 节点ID左移位数
	epoch           = int64(1288834974657)      // 起始时间戳（可以根据需要修改）
	maxAttempts     = 5                         // 最大尝试次数
	initialWaitTime = 100 * time.Millisecond    // 初始等待时间
)

// Snowflake结构定义
type Snowflake struct {
	timestamp int64
	nodeID    int64
	sequence  int64
}

// NewSnowflake 创建新的Snowflake实例
func NewSnowflake(nodeID int64) (*Snowflake, error) {
	if nodeID < 0 || nodeID > maxNodeID {
		return nil, fmt.Errorf("node ID must be between 0 and %d", maxNodeID)
	}
	return &Snowflake{
		timestamp: 0,
		nodeID:    nodeID,
		sequence:  0,
	}, nil
}

// currentTimeMillis 获取当前时间的毫秒数
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// waitUntilNextMillis 等待直到下一个毫秒
func waitUntilNextMillis(lastTimestamp int64) int64 {
	timestamp := currentTimeMillis()
	for timestamp <= lastTimestamp {
		timestamp = currentTimeMillis()
	}
	return timestamp
}

// NextID 生成下一个唯一ID
func (s *Snowflake) NextID() int64 {
	now := currentTimeMillis()

	// 获取当前时间戳并进行原子操作
	lastTimestamp := atomic.LoadInt64(&s.timestamp)

	if now < lastTimestamp {
		// 系统时间回退，尝试等待恢复
		for attempt := 0; attempt < maxAttempts; attempt++ {
			waitTime := initialWaitTime << attempt // 指数级增加等待时间
			time.Sleep(waitTime)
			now = currentTimeMillis()
			if now >= lastTimestamp {
				break
			}
		}
		if now < lastTimestamp {
			return generateFallbackID(s.nodeID)
		}
	}

	if lastTimestamp == now {
		sequence := atomic.AddInt64(&s.sequence, 1) & maxSequence
		if sequence == 0 {
			now = waitUntilNextMillis(lastTimestamp)
		}
		atomic.StoreInt64(&s.sequence, sequence)
	} else {
		atomic.StoreInt64(&s.sequence, 0)
	}

	atomic.StoreInt64(&s.timestamp, now)
	id := (now-epoch)<<timeShift | (s.nodeID << nodeShift) | atomic.LoadInt64(&s.sequence)
	return id
}

// generateFallbackID 生成备用ID
func generateFallbackID(nodeID int64) int64 {
	now := currentTimeMillis()
	sequence := atomic.AddInt64(&nodeID, 1) & maxSequence
	return (now-epoch)<<timeShift | (nodeID << nodeShift) | sequence
}
