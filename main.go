package goenvy

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func HasEnvFile() bool {
	file, err := os.Open(".env")
	defer file.Close()
	if err != nil {
		return false
	}
	return true
}

func Load() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	defer file.Close()
	setEnvironVars(file)
}

func GetEnv(envValue string) (result string, err error){
	if len(os.Getenv(envValue)) == 0 {
		log.Fatalf("Env Var Error: %s does not exist", envValue)
	}
	return os.Getenv(envValue), err
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
