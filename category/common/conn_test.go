package common

import (
	"log"
	"testing"
)

func TestGetMySQLFromConsul(t *testing.T) {
	consulConfig, err := GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Fatal("Get Consul Config Error")
	}
	log.Println(consulConfig)
	ret := GetMySQLFromConsul(consulConfig)
	if ret.User != "" {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}
	log.Println(ret.User, ret.Pwd, "success")
}

func TestGetMySQLFromSelf(t *testing.T) {
	ret := GetMySQLFromSelf()
	if ret.User == "root" {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}
	log.Println(ret.User, ret.Pwd, "success")

}
