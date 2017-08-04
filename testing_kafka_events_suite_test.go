package testing_kafka_events_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTestingKafkaEvents(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestingKafkaEvents Suite")
}
