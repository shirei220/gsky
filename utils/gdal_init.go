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
	// This is a bit nasty, but this is one way to work out which
	// drivers are present in the GDAL shared library. We then
	// load the drivers of interest and then load all of the
	// drivers. This places common drivers at the front of the
	// driver list.
	var haveNetCDF, haveHDF4, haveHDF5, haveJP2OpenJPEG bool
	var haveGTiff, haveGSKYNetCDF bool

	// Find out which drivers are present
	C.GDALAllRegister()
	for i := 0; i < int(C.GDALGetDriverCount()); i++ {
		driver := C.GDALGetDriver(C.int(i))
		switch C.GoString(C.GDALGetDriverShortName(driver)) {
		case "GSKY_netCDF":
			haveGSKYNetCDF = true
		case "netCDF":
			haveNetCDF = true
		case "HDF4":
			haveHDF4 = true
		case "HDF5":
			haveHDF5 = true
		case "JP2OpenJPEG":
			haveJP2OpenJPEG = true
		case "GTiff":
			haveGTiff = true
		}
	}
}
