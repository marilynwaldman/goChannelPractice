package ChannelSync

import "fmt"

func Map() {

	kvStore := make(map[string] string)
    kvStore["hi"] = "hello"
    kvStore["bye"] = "goodbye"

    for k,v := range kvStore{
    	fmt.Println("key ", k, " value ", v)
	}

	kvs := map[string]string {
		"joe" : "smith",
		"ralph" : "hampy",
	}

	for k,v := range kvs {
		fmt.Println(k, " ", v)
	}

}
