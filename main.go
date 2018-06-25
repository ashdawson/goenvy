package goenvy

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Load() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	defer file.Close()
	setEnvironVars(file)
}

func getSystemEnviron() map[string]string {
	sysEnvMap := make(map[string]string)
	sysEnv := os.Environ()
	for _, envVar := range sysEnv {
		env := strings.Split(envVar, "=")
		sysEnvMap[env[0]] = env[1]
	}

	return sysEnvMap
}

func setEnvironVars(file *os.File) {
	currEnvMap := make(map[string]string)
	sysEnvMap := getSystemEnviron()
	input := bufio.NewScanner(file)
	for input.Scan() {
		if strings.Contains(input.Text(), "=") {
			env := strings.Split(input.Text(), "=")
			currEnvMap[env[0]] = env[1]
		}
	}

	for key, value := range currEnvMap {
		if _, ok := sysEnvMap[key]; !ok {
			os.Setenv(key, value)
		}
	}
}
