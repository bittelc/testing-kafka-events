package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKafkaevents(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kafkaevents Suite")
}
