package csv2json

/*
陽性患者数

csv
No,全国地方公共団体コード,都道府県名,市区町村名,公表_年月日,曜日,陽性患者人数,死亡者人数
1,221309,静岡県,浜松市,2020-01-29,水,0,0
2,221309,静岡県,浜松市,2020-01-30,木,0,0
3,221309,静岡県,浜松市,2020-01-31,金,0,0
:
603,221309,静岡県,浜松市,2021-09-22,水,9,0
604,221309,静岡県,浜松市,2021-09-23,木,21,0
605,221309,静岡県,浜松市,2021-09-24,金,3,0
json
  "patients_summary": {
    "date": "2021/06/12 15:01",
    "data": [
      {
        "日付": "2020-01-29T08:00:00.000Z",
        "小計": 0
      },
      {
        "日付": "2020-01-30T08:00:00.000Z",
        "小計": 0
      },
      :
      :
      {
      	"日付": "2021-06-12T08:00:00.000Z",
        "小計": 3
      },
      {
        "日付": "2021-06-13T08:00:00.000Z",
        "小計": 0
      }
    ]
  },
*/

import (
	"errors"
	"time"

	"github.com/go-gota/gota/dataframe"
)

const keyPatientsSummaryDateOfPublicate = "公表_年月日"
const keyPatientsSummaryNumberOfPatients = "陽性患者人数"

type (
	PatientSummaryData struct {
		Date     string `json:"日付"`
		Subtotal int    `json:"小計"`
	}
	PatientsSummary struct {
		Date string               `json:"date"`
		Data []PatientSummaryData `json:"data"`
	}
)

func patientsSummary(df *dataframe.DataFrame, dtUpdated time.Time, dtEnd time.Time) (*PatientsSummary, error) {
	dfSelected := df.Select([]string{keyPatientsSummaryDateOfPublicate, keyPatientsSummaryNumberOfPatients})
	if df.Err != nil {
		return nil, df.Err
	}

	var dataList []PatientSummaryData
	for _, v := range dfSelected.Maps() {
		date, ok := v[keyPatientsSummaryDateOfPublicate].(string)
		if !ok {
			return nil, errors.New("unable to cast data")
		}
		number, ok := v[keyPatientsSummaryNumberOfPatients].(int)
		if !ok {
			return nil, errors.New("unable to cast data")
		}
		dataList = append(dataList, PatientSummaryData{
			Date:     date,
			Subtotal: number,
		})
	}

	ps := &PatientsSummary{
		Date: dtUpdated.Format("2006/01/02 15:04"),
		Data: dataList,
	}

	return ps, nil
}
