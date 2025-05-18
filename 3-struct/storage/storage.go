package storage

import (
	"dz/3-struct/bins"
	"encoding/json"
	"os"
)

func SaveBinList(bins bins.Bins) (bool, error) {
	binBytes := bins.ToBytes()
	file, err := os.OpenFile("bins.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return false, err
	}
	defer file.Close()
	_, err = file.Write(binBytes)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReadBinList() (bins.Bins, error) {
	file, err := os.ReadFile("bins.json")
	if err != nil {
		return bins.Bins{}, err
	}

	var newBins bins.Bins
	err = json.Unmarshal(file, &newBins)
	if err != nil {
		return bins.Bins{}, err
	}
	return newBins, nil
}
