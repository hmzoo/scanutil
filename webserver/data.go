package webserver

import (
  "github.com/gocarina/gocsv"
  "encoding/json"
  "os"
)

type Item struct{
  Inv string `csv:"inv" json:"inv"`
  Mod string `csv:"mod" json:"mod"`
  Mac string `csv:"mac" json:"mac"`
  Ser string `csv:"ser" json:"ser"`
}

type Data []*Item

func newData() Data {
  return Data{}
}

func (data *Data) JSON() (error,[]byte) {
	datajson, err := json.Marshal(*data)
  return err,datajson
}

func (data *Data) AddItem( i Item)  {
  for k,v :=range(*data){
    if (v.Inv == i.Inv){
      (*data)[k]=&i
      return
      }
  }
	*data=append(*data,&i)
}


func LoadCSV() (error, *Data) {
	var data Data
	csvFile, err := os.OpenFile("data.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err, nil
	}
	defer csvFile.Close()
	if err := gocsv.UnmarshalFile(csvFile, &data); err != nil {
		return err, nil
	}
	return nil, &data
}

func (data *Data) SaveCSV() error {
	os.Remove("data.csv")
	resultFile, err := os.OpenFile("data.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer resultFile.Close()

	err = gocsv.MarshalFile(data, resultFile)
	if err != nil {
		return err
	}

	return nil

}
