package main

import "dz/3-struct/bins"

type FileService interface {
	ReadFile(path string) ([]byte, error)
}

type StorageService interface {
	SaveBinList(path string, bins bins.Bins) error
	ReadBinList(path string) (bins.Bins, error)
}

func main() {

}
