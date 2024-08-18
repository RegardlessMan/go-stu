package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// 定义结构体
type Config struct {
	Server   Server   `ini:"server"`
	Database Database `ini:"database"`
}

type Server struct {
	Host string `ini:"host"`
	Port int    `ini:"port"`
}

type Database struct {
	Name     string `ini:"name"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}

// 解析 INI 文件
func parseIniFile(filePath string, config interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	configValue := reflect.ValueOf(config).Elem()
	configType := configValue.Type()
	// 每次循环都创建一个新的 scanner
	scanner := bufio.NewScanner(file)
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		sectionName := field.Tag.Get("ini")

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || strings.HasPrefix(line, ";") {
				break
			}

			if strings.HasPrefix(line, "["+sectionName+"]") {
				if err := parseSection(file, sectionName, configValue.Field(i), field, scanner); err != nil {
					return err
				}
				break
			}
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}

// 解析一个特定的部分
func parseSection(file *os.File, sectionName string, sectionValue reflect.Value, sectionType reflect.StructField,
	scanner *bufio.Scanner) error {
	//scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ";") {
			break
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		for j := 0; j < sectionType.Type.NumField(); j++ {
			subField := sectionType.Type.Field(j)
			tagName := subField.Tag.Get("ini")
			if tagName == key {
				err := setValue(sectionValue.Field(j), value)
				if err != nil {
					return err
				}
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// 设置字段的值
func setValue(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		intVal, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(intVal))
	default:
		return fmt.Errorf("unsupported type: %s", field.Type().String())
	}

	return nil
}
