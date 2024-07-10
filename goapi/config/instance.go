package config

import (
	"os"
	"strconv"
)

var CPU_REQUEST = os.Getenv("CPU_REQUEST")
var MEMORY_REQUEST = os.Getenv("MEMORY_REQUEST")
var CPU_LIMIT = os.Getenv("CPU_LIMIT")
var MEMORY_LIMIT = os.Getenv("MEMORY_LIMIT")
var DISK_REQUEST = os.Getenv("DISK_REQUEST")
var DISK_LIMIT = os.Getenv("DISK_LIMIT")
var TERMINATION_PERIOD = os.Getenv("TERMINATION_PERIOD")

var CONCURRENT_INSTANCES, _ = strconv.Atoi(os.Getenv("CONCURRENT_INSTANCES"))
var INSTANCE_NAME_SECRET = os.Getenv("INSTANCE_NAME_SECRET")
var DEFAULT_USERNAME = os.Getenv("DEFAULT_USERNAME")
var INSTANCE_HOSTNAME = os.Getenv("INSTANCE_HOSTNAME")
var INSTANCE_TIME, _ = strconv.Atoi(os.Getenv("INSTANCE_TIME"))
var MAX_INSTANCE_TIME, _ = strconv.Atoi(os.Getenv("MAX_INSTANCE_TIME"))