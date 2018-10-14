package config

import (
    "Bookee/infra/logger"
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "strconv"
    "strings"
    "sync"
)

//TODO Should be extracted to be an individual module

type configType map[interface{}]interface{}

var (
    configs     configType
    onceConfigs sync.Once
)

func loadConfigs(data string) configType {
    if len(data) > 0 {
        conf := make(configType)
        err := yaml.Unmarshal([]byte(data), &conf)
        if err == nil {
            return conf
        } else {
            log.Println(err.Error())
            return nil
        }
    } else {
        return nil
    }
}

func loadConfigsByPath(path string) configType {
    data, err := ioutil.ReadFile(path)
    if err == nil {
        return loadConfigs(string(data))
    } else {
        logger.L().Println(err.Error())
        return nil
    }
}

func loadConfigsDefault() configType {
    return loadConfigsByPath(`../../conf/application.yaml`)
}

func getConfigs() configType {
    onceConfigs.Do(func() {
        if configs == nil {
            configs = loadConfigsDefault()
        }
    })
    return configs
}

func GetStringOrEmpty(name string) string {
    return GetString(name, "")
}

func GetString(name string, defaultValue string) string {
    v := GetObject(name)
    if v != nil {
        return fmt.Sprint(v)
    } else {
        return defaultValue
    }
}

func GetIntOrZero(name string) int {
    return GetInt(name, 0)
}

func GetInt(name string, defaultValue int) int {
    v := GetObject(name)
    if v != nil {
        switch v.(type) {
        case int:
            return v.(int)
        case string:
            vStr := v.(string)
            vInt, err := strconv.Atoi(vStr)
            if err == nil {
                return vInt
            } else {
                return defaultValue
            }
        default:
            return defaultValue
        }
    } else {
        return defaultValue
    }
}

func GetObject(name string) (v interface{}) {
    conf := getConfigs()
    if conf != nil {
        keys := strings.Split(name, ".")
        for i := 0; i < len(keys); i++ {
            if conf == nil {
                // TODO log warn
                log.Printf(`%s not exists\n`, name)
                break
            } else {
                if i == len(keys)-1 {
                    v = conf[keys[i]]
                    break
                } else {
                    if sub, ok := conf[keys[i]].(configType); ok {
                        conf = sub
                    } else {
                        // TODO log warn
                        log.Printf(`%s not exists\n`, name[0:i+1])
                        break
                    }
                }
            }
        }
    } else {
        // TODO log warning
        log.Println(`%s not exists`, name)
    }
    return
}
