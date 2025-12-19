package utils

// #include "gdal.h"
// #include "gdal_frmts.h"
// #cgo pkg-config: gdal
import "C"

import (
	"os"
	"path/filepath"
)

func InitGdal() {
	setDefaultEnv("GDAL_NETCDF_VERIFY_DIMS", "NO")
	setDefaultEnv("GDAL_PAM_ENABLED", "NO")
	setDefaultEnv("GDAL_DISABLE_READDIR_ON_OPEN", "EMPTY_DIR")
	setDefaultEnv("GDAL_VRT_ENABLE_PYTHON", "YES")
	setDefaultEnv("GDAL_MAX_DATASET_POOL_SIZE", "10")

	exeFilePath, err := os.Executable()
	if err == nil {
		setDefaultEnv("GDAL_DRIVER_PATH", filepath.Dir(exeFilePath))
	}

	registerGDALDrivers()
}

func setDefaultEnv(envVar string, defaultVal string) {
	if _, ok := os.LookupEnv(envVar); !ok {
		os.Setenv(envVar, defaultVal)
	}
}

func registerGDALDrivers() {
	C.GDALAllRegister()
}
