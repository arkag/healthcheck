package main

import (
	"reflect"
	"regexp"
	"testing"
)

func TestParseYaml(t *testing.T) {
	data, err := ParseYaml("test_input.yml")
	want := regexp.MustCompile("HealthCheckItem")

	typeObject := reflect.TypeOf(data)
	typeName := typeObject.Elem().Name()

	if !want.MatchString(typeName) || err != nil {
		t.Fatalf(`ParseYaml("test_input.yml") = %q, %v, want match for %#q, nil`, typeName, err, want)
	}
}

func TestGetEndpointHosts(t *testing.T) {
	items, err := ParseYaml("test_input.yml")
	data := GetEndpointHosts(items)

	if data == nil || err != nil {
		t.Fatalf(`GetEndpointHosts(HealthCheckItem) = nil, want populated`)
	}
}

func TestSendHealthCheckUp(t *testing.T) {
	want := "UP"
	hci := HealthCheckItem{
		Name: "Google test",
		Url:  "https://www.google.com",
	}
	status := SendHealthCheck(hci)
	if status != want {
		t.Fatalf(`HCI for Google != %v, want %v`, want, want)
	}
}

func TestSendHealthCheckDown(t *testing.T) {
	want := "DOWN"
	hci := HealthCheckItem{
		Name:   "Google test",
		Method: "POST",
		Url:    "https://www.google.com/v4/somefunc",
	}
	status := SendHealthCheck(hci)
	if status != want {
		t.Fatalf(`HCI for Google != %v, want %v`, want, want)
	}
}
