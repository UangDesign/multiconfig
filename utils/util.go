package util

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

// type CommonJsonMapping map[string]interface{}

func GetJsonIterator() (oJsonIter jsoniter.API) {
	oJsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
	return oJsonIter
}

func GetCommonJsonMapping(oData []byte) (oJsonMapping map[string]interface{}) {
	oJsonMapping = make(map[string]interface{})
	oJsonError := GetJsonIterator().Unmarshal(oData, &oJsonMapping)
	if oJsonError != nil {
		//err = Error.MakeError(Error.ERR_CRITICAL)
		oJsonMapping = nil
	}
	return oJsonMapping
}

func GetCurrentDirectory() string {
	cwdPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		cwdPath = ""
		fmt.Println(err)
	}
	basePath := strings.Replace(cwdPath, "\\", "/", -1)
	return basePath
}

func GetRealDirectory() string {
	curr_path := GetCurrentDirectory()
	myslice := strings.Split(curr_path, "/")
	uppath := strings.Join(myslice[:len(myslice)-1], "/")
	return uppath
}

func IsFile(path string) (blFile bool) {
	f, err := os.Stat(path)
	if err == nil {
		if !f.IsDir() {
			blFile = true
		} else {
			blFile = false
		}
	} else {
		if os.IsNotExist(err) {
			blFile = false
		} else {
			blFile = true
		}
	}
	return blFile
}

func IsDir(path string) (blDir bool) {
	f, err := os.Stat(path)
	if err == nil {
		if f.IsDir() {
			blDir = true
		} else {
			blDir = false
		}
	} else {
		if os.IsNotExist(err) {
			blDir = false
		} else {
			blDir = true
		}
	}
	return blDir
}

func BinQuery(st []string, query string) bool {
	lenth := len(st)
	sindex := lenth / 2
	if sindex == 0 {
		return false
	}
	for i := 0; i < sindex; i++ {
		if st[i] == query {
			return true
		}
	}
	isQue := BinQuery(st[sindex:], query)
	return isQue
}

func UniqSlice(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

func IsExist(path string) (blExist bool) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			blExist = false
		} else {
			blExist = true
		}
	} else {
		blExist = true
	}
	return blExist
}

func IsLittleEndian() (IsLittle bool) {
	Example := 0x1234
	IsLittle = *(*byte)(unsafe.Pointer(&Example)) == 0x34
	return IsLittle
}

func GetSystemEndian() (Order binary.ByteOrder) {
	if IsLittleEndian() {
		Order = binary.LittleEndian
	} else {
		Order = binary.BigEndian
	}
	return Order
}

func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
