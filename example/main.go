package main

import (
	"fmt"

	"github.com/UangDesign/multiconfig"
)

var (
	TEST_INT             int
	TEST_TEMP_INT        int
	TEST_INT64           int64
	TEST_TEMP_INT64      int64
	TEST_STRING          string
	TEST_TEMP_STRING     string
	TEST_BOOL            bool
	TEST_TEMP_BOOL_1     bool
	TEST_TEMP_BOOL_0     bool
	TEST_TEMP_BOOL_FALSE bool
	TEST_STRINGLIST      []string
	TEST_INTLIST         []int
)

var (
	intMap        map[string]int
	int64Map      map[string]int64
	uintMap       map[string]uint
	uint64Map     map[string]uint64
	stringMap     map[string]string
	boolMap       map[string]bool
	stringListMap map[string][]string
	intListMap    map[string][]int
)

func init() {
	multiConfig := multiconfig.NewMultiConfig("D:/git/multiconfig/example/config.conf", "D:/git/multiconfig/example/temp.conf")
	intMap = multiConfig.ConfigInt.ParseConfig()
	int64Map = multiConfig.ConfigInt64.ParseConfig()
	uintMap = multiConfig.ConfigUint.ParseConfig()
	uint64Map = multiConfig.ConfigUint64.ParseConfig()
	stringMap = multiConfig.ConfigString.ParseConfig()
	boolMap = multiConfig.ConfigBool.ParseConfig()
	stringListMap = multiConfig.ConfigStringList.ParseConfig()
	intListMap = multiConfig.ConfigIntList.ParseConfig()
}

func outputConfig() {
	for key, value := range intMap {
		fmt.Printf("k: %-40v v: %v\n", key, value)
	}
	for key, value := range int64Map {
		fmt.Printf("k: %-40v v: %v\n", key, value)
	}
	for key, value := range uintMap {
		fmt.Printf("k: %-40v v: %v\n", key, value)
	}
	for key, value := range uint64Map {
		fmt.Printf("k: %-40v v: %v\n", key, value)
	}
	for key, value := range stringMap {
		fmt.Printf("k: %-40v v: %v\n", key, value)
	}
	for key, value := range boolMap {
		fmt.Printf("k: %-40v v: %v\n", key, value)
	}
	for key, value := range stringListMap {
		fmt.Printf("k: %-40v v: %v\n", key, value)
	}
	for key, value := range intListMap {
		fmt.Printf("k: %-40v v: %v\n", key, value)
	}
	// Specify the key output value
	TEST_INT = intMap["TEST_INT"]
	TEST_TEMP_INT = intMap["TEST_TEST_INT"]
	TEST_INT64 = int64Map["TEST_INT64"]
	TEST_TEMP_INT64 = int64Map["TEST_TEMP_INT64"]
	TEST_STRING = stringMap["TEST_STRING"]
	TEST_TEMP_STRING = stringMap["TEST_TEMP_STRING"]
	TEST_BOOL = boolMap["TEST_BOOL"]
	TEST_TEMP_BOOL_1 = boolMap["TEST_TEMP_BOOL_1"]
	TEST_TEMP_BOOL_0 = boolMap["TEST_TEMP_BOOL_0"]
	TEST_TEMP_BOOL_FALSE = boolMap["TEST_TEMP_BOOL_FALSE"]
	TEST_STRINGLIST = stringListMap["TEST_STRINGLIST"]
	TEST_INTLIST = intListMap["TEST_INTLIST"]
	fmt.Printf(" TEST_INT: %v\n TEST_TEMP_INT:%v\n TEST_INT64:%v\n TEST_STRING:%v\n TEST_BOOL:%v\n TEST_STRINGLIST:%v\n TEST_INTLIST:%v\n",
		TEST_INT,
		TEST_TEMP_INT,
		TEST_INT64,
		TEST_STRING,
		TEST_BOOL,
		TEST_STRINGLIST,
		TEST_INTLIST,
	)
}

func main() {
	outputConfig()
}
