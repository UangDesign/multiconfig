package multiconfig

import "github.com/UangDesign/multiconfig/singleconfig"

type MultiConfig struct {
	multiConfig      []singleconfig.SingleConfig
	configString     map[string]string
	configBool       map[string]bool
	configInt        map[string]int
	configInt64      map[string]int64
	configUint       map[string]uint
	configUint64     map[string]uint64
	configFloat32    map[string]float32
	configFloat64    map[string]float64
	configStringList map[string][]string
	configIntList    map[string][]int
}

func NewMultiConfig(confPath string, moreConf ...string) (config *MultiConfig) {
	if len(confPath) < 1 {
		config = nil
	} else {
		config = &MultiConfig{
			multiConfig:      make([]singleconfig.SingleConfig, 0),
			configString:     make(map[string]string),
			configBool:       make(map[string]bool),
			configInt:        make(map[string]int),
			configInt64:      make(map[string]int64),
			configUint:       make(map[string]uint),
			configUint64:     make(map[string]uint64),
			configFloat32:    make(map[string]float32),
			configFloat64:    make(map[string]float64),
			configStringList: make(map[string][]string),
			configIntList:    make(map[string][]int),
		}
		// confPath
		oSingleConfig := singleconfig.NewSingleConfig(confPath)
		if oSingleConfig != nil {
			config.multiConfig = append(config.multiConfig, *oSingleConfig)
		}
		// moreConf
		for i := range moreConf {
			oSingleConfig := singleconfig.NewSingleConfig(moreConf[i])
			if oSingleConfig != nil {
				config.multiConfig = append(config.multiConfig, *oSingleConfig)
			}
		}
	}
	return config
}

func (m *MultiConfig) ParseString() map[string]string {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigString.ParseConfig(singleConfig.GetConfigFile()) {
			m.configString[k] = v
		}
	}
	return m.configString
}

func (m *MultiConfig) ParseBool() map[string]bool {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigBool.ParseConfig(singleConfig.GetConfigFile()) {
			m.configBool[k] = v
		}
	}
	return m.configBool
}

func (m *MultiConfig) ParseInt() map[string]int {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigInt.ParseConfig(singleConfig.GetConfigFile()) {
			m.configInt[k] = v
		}
	}
	return m.configInt
}

func (m *MultiConfig) ParseInt64() map[string]int64 {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigInt64.ParseConfig(singleConfig.GetConfigFile()) {
			m.configInt64[k] = v
		}
	}
	return m.configInt64
}

func (m *MultiConfig) ParseUint() map[string]uint {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigUint.ParseConfig(singleConfig.GetConfigFile()) {
			m.configUint[k] = v
		}
	}
	return m.configUint
}

func (m *MultiConfig) ParseUint64() map[string]uint64 {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigUint64.ParseConfig(singleConfig.GetConfigFile()) {
			m.configUint64[k] = v
		}
	}
	return m.configUint64
}

func (m *MultiConfig) ParseFloat32() map[string]float32 {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigFloat32.ParseConfig(singleConfig.GetConfigFile()) {
			m.configFloat32[k] = v
		}
	}
	return m.configFloat32
}

func (m *MultiConfig) ParseFloat64() map[string]float64 {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigFloat64.ParseConfig(singleConfig.GetConfigFile()) {
			m.configFloat64[k] = v
		}
	}
	return m.configFloat64
}

func (m *MultiConfig) ParseStringList() map[string][]string {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigStringList.ParseConfig(singleConfig.GetConfigFile()) {
			m.configStringList[k] = v
		}
	}
	return m.configStringList
}
func (m *MultiConfig) ParseIntList() map[string][]int {
	for _, singleConfig := range m.multiConfig {
		for k, v := range singleConfig.ConfigIntList.ParseConfig(singleConfig.GetConfigFile()) {
			m.configIntList[k] = v
		}
	}
	return m.configIntList
}

func (m *MultiConfig) SetValue(key string, value interface{}, filePath string) (err error) {
	var valueType string
	if filePath != "" {
		for _, singleConfig := range m.multiConfig {
			if singleConfig.GetConfPath() == filePath {
				valueType, err = singleConfig.SetValue(key, value)
			}
		}
	} else {
		for _, singleConfig := range m.multiConfig {
			if singleConfig.HasKey(key) {
				valueType, err = singleConfig.SetValue(key, value)
			}
		}
	}
	m.reLoadConfig(valueType)
	return err
}

func (m *MultiConfig) FlushToConfig() (err error) {
	for _, singleConfig := range m.multiConfig {
		err = singleConfig.FlushToConfig()
		if err != nil {
			break
		}
	}
	return
}

func (m *MultiConfig) reLoadConfig(valueType string) {
	switch valueType {
	case "string":
		m.ParseString()
	case "bool":
		m.ParseBool()
	case "int":
		m.ParseInt()
	case "int64":
		m.ParseInt64()
	case "uint":
		m.ParseUint()
	case "uint64":
		m.ParseUint64()
	case "float32":
		m.ParseFloat32()
	case "float64":
		m.ParseFloat64()
	case "[]string":
		m.ParseStringList()
	case "[]int":
		m.ParseIntList()
	}
}
