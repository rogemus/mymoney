package utils

import "log"

func LogFatal(msg string) {
	log.Fatalf("[FATAL] %s", msg)
}

func LogError(msg string) {
	log.Printf("[ERROR] %s", msg)
}

func LogInfo(msg string) {
	log.Printf("[INFO] %s", msg)
}

func LogWarning(msg string) {
	log.Printf("[WARN] %s", msg)
}
