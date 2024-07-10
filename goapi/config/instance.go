package config

import "os"

var CPU_REQUEST = os.Getenv("CPU_REQUEST")
var MEMORY_REQUEST = os.Getenv("MEMORY_REQUEST")
var CPU_LIMIT = os.Getenv("CPU_LIMIT")
var MEMORY_LIMIT = os.Getenv("MEMORY_LIMIT")
var DISK_REQUEST = os.Getenv("DISK_REQUEST")
var DISK_LIMIT = os.Getenv("DISK_LIMIT")
var TERMINATION_PERIOD = os.Getenv("TERMINATION_PERIOD")