package config

type Conf struct {
	KafkaBroker        []string
	KafkaTopicInbound  string
	KafkaConsumerGroup string
	ConnectionDB       string
}

var con *Conf

func InitConf() *Conf {
	if con == nil {
		con = &Conf{}
	}
	con.KafkaBroker = []string{"kafka-location-test.com:9092"}
	con.KafkaTopicInbound = "wa-test"
	con.KafkaConsumerGroup = "wa-consumer-group-test"
	con.ConnectionDB = "tcp://locaiton-test.com:9000?debug=true"
	return con
}
