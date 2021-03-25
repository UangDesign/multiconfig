package multiconfig

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Unknwon/goconfig"
)

/*
	support type：
		int,int16,int32,int64,int8
		uint,uint16,uint32,uint64,uint8
		float32,float64
		string
		bool
		[]int
		[]string
*/

type ConfigType string

const (
	CFG_STRING     ConfigType = "sectionString"
	CFG_BOOL       ConfigType = "sectionBool"
	CFG_INT        ConfigType = "sectionInt"
	CFG_UINT       ConfigType = "sectionUint"
	CFG_INT64      ConfigType = "sectionInt64"
	CFG_UINT64     ConfigType = "sectionUint64"
	CFG_STRINGLIST ConfigType = "sectionStringList" // etc: [one,two,three]
	CFG_INTLIST    ConfigType = "sectionIntList"    // etc: [1,2,3]
	CFG_FLOAT32    ConfigType = "sectionFloat32"
	CFG_FLOAT64    ConfigType = "sectionFloat64"
)

var (
	cfg *goconfig.ConfigFile
)

type configString struct {
	config map[string]string
}

type configBool struct {
	config map[string]bool
}

type configInt struct {
	config map[string]int
}

type configInt64 struct {
	config map[string]int64
}

type configUint struct {
	config map[string]uint
}

type configUint64 struct {
	config map[string]uint64
}

type configStringList struct {
	config map[string][]string
}

type configIntList struct {
	config map[string][]int
}

type configFloat32 struct {
	config map[string]float32
}

type configFloat64 struct {
	config map[string]float64
}

// parseConfig is used to parse the string configuration
func (c *configString) ParseConfig() map[string]string {
	return getSection(CFG_STRING)
}

// parseConfig is used to parse the bool configuration
func (c *configBool) ParseConfig() map[string]bool {
	ret := make(map[string]bool)
	for k, v := range getSection(CFG_BOOL) {
		if vb, err := strconv.ParseBool(v); err == nil {
			ret[k] = vb
		}
	}
	return ret
}

// parseConfig is used to parse the int configuration
func (c *configInt) ParseConfig() map[string]int {
	ret := make(map[string]int)
	for k, v := range getSection(CFG_INT) {
		if vInt, err := strconv.Atoi(v); err == nil {
			ret[k] = vInt
		}
	}
	return ret
}

// parseConfig is used to parse the uint configuration
func (c *configUint) ParseConfig() map[string]uint {
	ret := make(map[string]uint)
	for k, v := range getSection(CFG_UINT) {
		if vUint, err := strconv.ParseUint(v, 10, 0); err == nil {
			ret[k] = uint(vUint)
		}
	}
	return ret
}

// parseConfig is used to parse the int64 configuration
func (c *configInt64) ParseConfig() map[string]int64 {
	ret := make(map[string]int64)
	for k, v := range getSection(CFG_INT64) {
		if vInt64, err := strconv.ParseInt(v, 10, 0); err == nil {
			ret[k] = vInt64
		}
	}
	return ret
}

// parseConfig is used to parse the uint64 configuration
func (c *configUint64) ParseConfig() map[string]uint64 {
	ret := make(map[string]uint64)
	for k, v := range getSection(CFG_UINT64) {
		if vUint, err := strconv.ParseUint(v, 10, 0); err == nil {
			ret[k] = vUint
		}
	}
	return ret
}

func getSection(configType ConfigType) map[string]string {
	configMap, err := cfg.GetSection(string(configType))
	if err != nil {
		fmt.Printf(fmt.Sprintf("sectionString parseFailed, err is:%v\n", err))
	}
	return configMap
}

// parseConfig is used to parse the []string configuration
func (c *configStringList) ParseConfig() map[string][]string {
	ret := make(map[string][]string)
	for k, v := range getSection(CFG_STRINGLIST) {
		if isList(v) {
			if trimBracket(&v); v != "" {
				vList := strings.Split(v, ",")
				if len(vList) > 0 {
					ret[k] = trimSpace(vList)
				}
			}
		}
	}
	return ret
}

// parseConfig is used to parse the []int configuration
func (c *configIntList) ParseConfig() map[string][]int {
	ret := make(map[string][]int)
	for k, v := range getSection(CFG_INTLIST) {
		if isList(v) {
			if trimBracket(&v); v != "" {
				vList := strings.Split(v, ",")
				if len(vList) > 0 {
					ret[k] = trimSpaceToIntList(vList)
				}
			}
		}
	}
	return ret
}

// parseConfig is used to parse the float32 configuration
func (c *configFloat32) ParseConfig() map[string]float32 {
	ret := make(map[string]float32)
	for k, v := range getSection(CFG_FLOAT32) {
		if vFloat32, err := strconv.ParseFloat(v, 32); err == nil {
			ret[k] = float32(vFloat32)
		}
	}
	return ret
}

// parseConfig is used to parse the float64 configuration
func (c *configFloat64) ParseConfig() map[string]float64 {
	ret := make(map[string]float64)
	for k, v := range getSection(CFG_FLOAT64) {
		if vFloat64, err := strconv.ParseFloat(v, 64); err == nil {
			ret[k] = vFloat64
		}
	}
	return ret
}

// isList determine if it is a list
func isList(list string) (is bool) {
	if strings.HasPrefix(list, "[") && strings.HasSuffix(list, "]") {
		is = true
	} else {
		is = false
	}
	return is
}

func trimBracket(list *string) {
	*list = strings.TrimRight(strings.TrimLeft(*list, "["), "]")
}

func trimSpaceToIntList(stringList []string) (ret []int) {
	ret = make([]int, 0)
	for _, v := range stringList {
		if v = strings.Trim(v, " "); v != "" {
			if vInt, err := strconv.Atoi(v); err == nil {
				ret = append(ret, vInt)
			}
		}
	}
	return ret
}

func trimSpace(stringList []string) (ret []string) {
	ret = make([]string, 0)
	for _, v := range stringList {
		if v = strings.Trim(v, " "); v != "" {
			ret = append(ret, v)
		}
	}
	return ret
}

type MultiConfig struct {
	ConfigString     *configString
	ConfigBool       *configBool
	ConfigInt        *configInt
	ConfigUint       *configUint
	ConfigInt64      *configInt64
	ConfigUint64     *configUint64
	ConfigStringList *configStringList
	ConfigIntList    *configIntList
	ConfigFloat32    *configFloat32
	ConfigFloat64    *configFloat64
}

func NewMultiConfig(confPath string, moreConf ...string) (multiConfig *MultiConfig) {
	if len(confPath) < 1 {
		multiConfig = nil
	} else {
		cfg = getConfHandler(confPath, moreConf...)
		multiConfig = &MultiConfig{
			ConfigString: new(configString),
			ConfigBool:   new(configBool),
			ConfigInt:    new(configInt),
			ConfigInt64:  new(configInt64),
		}
	}
	return
}

func getConfHandler(filename string, moreConf ...string) (cfg *goconfig.ConfigFile) {
	cfg, err := goconfig.LoadConfigFile(filename, moreConf...)
	if err != nil {
		panic(fmt.Sprintf("getcfg failed err is:%v", err))
	}
	return cfg
}
