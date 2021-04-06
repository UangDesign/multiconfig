package singleconfig

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	util "github.com/UangDesign/multiconfig/utils"

	"github.com/Unknwon/goconfig"
)

type SingleConfig struct {
	filePath         string
	cfg              *goconfig.ConfigFile
	ConfigString     configString
	ConfigBool       configBool
	ConfigInt        configInt
	ConfigUint       configUint
	ConfigInt64      configInt64
	ConfigUint64     configUint64
	ConfigStringList configStringList
	ConfigIntList    configIntList
	ConfigFloat32    configFloat32
	ConfigFloat64    configFloat64
}

func NewSingleConfig(filePath string) (config *SingleConfig) {
	if !util.IsFile(filePath) {
		config = nil
	} else {
		config = &SingleConfig{
			cfg:              getConfHandler(filePath),
			filePath:         filePath,
			ConfigString:     configString{config: make(map[string]string)},
			ConfigBool:       configBool{config: make(map[string]bool)},
			ConfigInt:        configInt{config: make(map[string]int)},
			ConfigUint:       configUint{config: make(map[string]uint)},
			ConfigInt64:      configInt64{config: make(map[string]int64)},
			ConfigUint64:     configUint64{config: make(map[string]uint64)},
			ConfigStringList: configStringList{config: make(map[string][]string)},
			ConfigIntList:    configIntList{config: make(map[string][]int)},
			ConfigFloat32:    configFloat32{config: make(map[string]float32)},
			ConfigFloat64:    configFloat64{config: make(map[string]float64)},
		}
	}
	return config
}

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
func (c *configString) ParseConfig(cfg *goconfig.ConfigFile) map[string]string {
	c.config = getSection(CFG_STRING, cfg)
	return c.config
}

// parseConfig is used to parse the bool configuration
func (c *configBool) ParseConfig(cfg *goconfig.ConfigFile) map[string]bool {
	for k, v := range getSection(CFG_BOOL, cfg) {
		if vb, err := strconv.ParseBool(v); err == nil {
			c.config[k] = vb
		}
	}
	return c.config
}

// parseConfig is used to parse the int configuration
func (c *configInt) ParseConfig(cfg *goconfig.ConfigFile) map[string]int {
	for k, v := range getSection(CFG_INT, cfg) {
		if vInt, err := strconv.Atoi(v); err == nil {
			c.config[k] = vInt
		}
	}
	return c.config
}

// parseConfig is used to parse the uint configuration
func (c *configUint) ParseConfig(cfg *goconfig.ConfigFile) map[string]uint {
	for k, v := range getSection(CFG_UINT, cfg) {
		if vUint, err := strconv.ParseUint(v, 10, 0); err == nil {
			c.config[k] = uint(vUint)
		}
	}
	return c.config
}

// parseConfig is used to parse the int64 configuration
func (c *configInt64) ParseConfig(cfg *goconfig.ConfigFile) map[string]int64 {
	for k, v := range getSection(CFG_INT64, cfg) {
		if vInt64, err := strconv.ParseInt(v, 10, 0); err == nil {
			c.config[k] = vInt64
		}
	}
	return c.config
}

// parseConfig is used to parse the uint64 configuration
func (c *configUint64) ParseConfig(cfg *goconfig.ConfigFile) map[string]uint64 {
	for k, v := range getSection(CFG_UINT64, cfg) {
		if vUint, err := strconv.ParseUint(v, 10, 0); err == nil {
			c.config[k] = vUint
		}
	}
	return c.config
}

func getSection(configType ConfigType, cfg *goconfig.ConfigFile) map[string]string {
	configMap, _ := cfg.GetSection(string(configType))
	return configMap
}

// parseConfig is used to parse the []string configuration
func (c *configStringList) ParseConfig(cfg *goconfig.ConfigFile) map[string][]string {
	for k, v := range getSection(CFG_STRINGLIST, cfg) {
		if isList(v) {
			if trimBracket(&v); v != "" {
				vList := strings.Split(v, ",")
				if len(vList) > 0 {
					c.config[k] = trimSpace(vList)
				}
			}
		}
	}
	return c.config
}

// parseConfig is used to parse the []int configuration
func (c *configIntList) ParseConfig(cfg *goconfig.ConfigFile) map[string][]int {
	for k, v := range getSection(CFG_INTLIST, cfg) {
		if isList(v) {
			if trimBracket(&v); v != "" {
				vList := strings.Split(v, ",")
				if len(vList) > 0 {
					c.config[k] = trimSpaceToIntList(vList)
				}
			}
		}
	}
	return c.config
}

// parseConfig is used to parse the float32 configuration
func (c *configFloat32) ParseConfig(cfg *goconfig.ConfigFile) map[string]float32 {
	for k, v := range getSection(CFG_FLOAT32, cfg) {
		if vFloat32, err := strconv.ParseFloat(v, 32); err == nil {
			c.config[k] = float32(vFloat32)
		}
	}
	return c.config
}

// parseConfig is used to parse the float64 configuration
func (c *configFloat64) ParseConfig(cfg *goconfig.ConfigFile) map[string]float64 {
	for k, v := range getSection(CFG_FLOAT64, cfg) {
		if vFloat64, err := strconv.ParseFloat(v, 64); err == nil {
			c.config[k] = vFloat64
		}
	}
	return c.config
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

func (s *SingleConfig) GetConfigFile() *goconfig.ConfigFile {
	return s.cfg
}

func (s *SingleConfig) GetConfPath() (filePath string) {
	return s.filePath
}

func (s *SingleConfig) HasKey(key string) (has bool) {
	if _, ok := s.ConfigBool.config[key]; ok {
		has = true
	} else if _, ok := s.ConfigString.config[key]; ok {
		has = true
	} else if _, ok := s.ConfigInt.config[key]; ok {
		has = true
	} else if _, ok := s.ConfigInt64.config[key]; ok {
		has = true
	} else if _, ok := s.ConfigUint.config[key]; ok {
		has = true
	} else if _, ok := s.ConfigInt64.config[key]; ok {
		has = true
	} else if _, ok := s.ConfigFloat32.config[key]; ok {
		has = true
	} else if _, ok := s.ConfigFloat64.config[key]; ok {
		has = true
	} else if _, ok := s.ConfigIntList.config[key]; ok {
		has = true
	} else if _, ok := s.ConfigStringList.config[key]; ok {
		has = true
	}
	return has
}

func (s *SingleConfig) SetValue(key string, value interface{}) (valueType string, err error) {
	valueType = reflect.TypeOf(value).Name()
	switch valueType {
	case "string":
		s.cfg.SetValue(string(CFG_STRING), key, value.(string))
		s.ConfigString.ParseConfig(s.cfg)
	case "bool":
		s.cfg.SetValue(string(CFG_BOOL), key, fmt.Sprintf("%v", value.(bool)))
		s.ConfigBool.ParseConfig(s.cfg)
	case "int":
		s.cfg.SetValue(string(CFG_INT), key, fmt.Sprintf("%v", value.(int)))
		s.ConfigInt.ParseConfig(s.cfg)
	case "int64":
		s.cfg.SetValue(string(CFG_INT64), key, fmt.Sprintf("%v", value.(int64)))
		s.ConfigInt64.ParseConfig(s.cfg)
	case "uint":
		s.cfg.SetValue(string(CFG_UINT), key, fmt.Sprintf("%v", value.(uint)))
		s.ConfigUint.ParseConfig(s.cfg)
	case "uint64":
		s.cfg.SetValue(string(CFG_UINT64), key, fmt.Sprintf("%v", value.(uint64)))
		s.ConfigUint64.ParseConfig(s.cfg)
	case "float32":
		s.cfg.SetValue(string(CFG_FLOAT32), key, fmt.Sprintf("%v", value.(float32)))
		s.ConfigFloat32.ParseConfig(s.cfg)
	case "float64":
		s.cfg.SetValue(string(CFG_FLOAT64), key, fmt.Sprintf("%v", value.(float64)))
		s.ConfigFloat64.ParseConfig(s.cfg)
	default:
		if v, ok := value.([]string); ok {
			s.cfg.SetValue(string(CFG_STRING), key, fmt.Sprintf("[%v]", strings.Join(v, ",")))
			s.ConfigStringList.ParseConfig(s.cfg)
			valueType = "[]string"
		} else if v, ok := value.([]int); ok {
			intToString := []string{}
			for _, intV := range v {
				intToString = append(intToString, strconv.Itoa(intV))
			}
			s.cfg.SetValue(string(CFG_INTLIST), key, fmt.Sprintf("[%v]", strings.Join(intToString, ",")))
			s.ConfigIntList.ParseConfig(s.cfg)
			valueType = "[]int"
		}
	}
	return valueType, err
}

func (s *SingleConfig) FlushToConfig() (err error) {
	return goconfig.SaveConfigFile(s.cfg, s.filePath)
}

func getConfHandler(filename string) (cfg *goconfig.ConfigFile) {
	cfg, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		panic(fmt.Sprintf("getcfg failed err is:%v", err))
	}
	return cfg
}
