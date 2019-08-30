package kafkactl

type partitionerStrategy int

const (
	random  partitionerStrategy = 0
	mannual partitionerStrategy = 1
	hash	partitionerStrategy = 2
)

// KafkaTopic represents a topic
type KafkaTopic struct {
	name 	   string
	partitions int
	strategy   partitionerStrategy
}

var (
	// TopicPlayer topic
	TopicPlayer = KafkaTopic{name: "player", partitions: 2, strategy: hash}
)
