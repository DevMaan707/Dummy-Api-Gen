package generator

import (
    "log"
)


func LogInfo(message string) {
    log.Printf("[INFO] %s\n", message)
}


func LogError(err error) {
    log.Printf("[ERROR] %v\n", err)
}


func ValidateResponseModelName(name string) bool {
    return len(name) > 13 && name[len(name)-13:] == "ResponseModel"
}
