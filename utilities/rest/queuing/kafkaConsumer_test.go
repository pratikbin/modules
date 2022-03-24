package queuing

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema"
)

func TestKafkaTopicConsumer(t *testing.T) {
	testConsumers := []string{"testConsumers"}

	var Codec = codec.New()

	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)

	require.Panics(t, func() {
		testKafkaState := NewKafkaState(testConsumers)
		partitionConsumer := testKafkaState.Consumers["Topic"]

		var kafkaStore kafkaMsg
		if len(partitionConsumer.Messages()) == 0 {
			kafkaStore = kafkaMsg{Msg: nil}
		}

		kafkaMsg := <-partitionConsumer.Messages()

		err := Codec.UnmarshalJSON(kafkaMsg.Value, &kafkaStore)
		if err != nil {
			panic(err)
		}

		require.Equal(t, kafkaTopicConsumer("Topic", testKafkaState.Consumers, Codec), kafkaStore)
	})
}

func TestNewConsumer(t *testing.T) {
	consumers := []string{"testConsumers"}
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer(consumers, config)

	require.Nilf(t, err, "should not happened. err %v", err)

	require.Panics(t, func() {
		require.Equal(t, newConsumer(consumers), consumer)
	})
}
