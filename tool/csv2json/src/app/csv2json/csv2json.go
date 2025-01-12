package csv2json

import (
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"

	"app/utils/logger"
)

type (
	CsvData struct {
		DfCsv     *dataframe.DataFrame
		DtUpdated time.Time
	}
)

type Result struct {
	Contacts          *Contacts          `json:"contacts,omitempty"`
	InspectionPersons *InspectionPersons `json:"inspection_persons,omitempty"`
	MainSummary       *MainSummary       `json:"main_summary,omitempty"`
	Patients          *Patients          `json:"patients,omitempty"`
	PatientsSummary   *PatientsSummary   `json:"patients_summary,omitempty"`
	Value             int                `json:"value"`
	HasError          bool               `json:"hasError"`
	LastUpdate        string             `json:"lastUpdate"`
}

type accessor interface {
	GetCSVDataFrameFromApi(apiAddress string, apiId string) (*dataframe.DataFrame, time.Time, error)
	GetTimeNow() time.Time
}

type Csv2Json struct {
	csvAccessor accessor
}

func NewCsv2Json(csvAccessorIn accessor) *Csv2Json {
	return &Csv2Json{csvAccessor: csvAccessorIn}
}

// 同じCSVデータを何度も読みにいかないためにバックアップしておくための変数
// key	: csv address
var mapCSVDataBackup = make(map[string](*CsvData))

// Process はオープンデータのCSVをJSONに変換する
func (c2j *Csv2Json) Process(apiAddress string, queryStrPrm string) (*Result, error) {
	r := &Result{
		Value:    0,
		HasError: false,
	}
	logger.Infos(apiAddress, queryStrPrm)

	dtLastUpdate := time.Date(2000, 1, 1, 1, 1, 0, 0, time.Local)
	types := strings.Split(queryStrPrm, ",")
	for index, value := range types {
		timeStart := time.Now()

		values := strings.Split(value, ":")
		if len(values) != 2 {
			r.HasError = true
			logger.Errors(value, "invalid query param...")
			continue
		}

		key := values[0]
		apiId := values[1]
		logger.Infof("%d, key=%s, id=%s", index, key, apiId)
		csvData, err := getCSVDataFrame(apiAddress, apiId, c2j.csvAccessor)

		if err != nil {
			r.HasError = true
			setEmptyStructToResult(key, r)
			logger.Errors(key, err)
			continue
		}

		switch key {
		case "patients_summary":
			totalCount := 0
			r.PatientsSummary, totalCount, err = patientsSummary(csvData.DfCsv, csvData.DtUpdated, c2j.csvAccessor.GetTimeNow())
			if err != nil {
				r.HasError = true
				setEmptyStructToResult(key, r)
				logger.Errors(key, err)
				continue
			}
			r.MainSummary = getEmptyMainSummary(csvData.DtUpdated)
			err = mainSummaryTry2Merge4Deth(csvData.DfCsv, r.MainSummary)
			r.MainSummary.Children[0].Value = totalCount
		case "inspection_persons":
			r.InspectionPersons, err = inspectionPersons(csvData.DfCsv, csvData.DtUpdated)
		case "contacts":
			r.Contacts, err = contacts(csvData.DfCsv, csvData.DtUpdated)
		default:
			r.HasError = true
			logger.Errors(key, "key not supported...")
			continue
		}

		if err != nil {
			r.HasError = true
			setEmptyStructToResult(key, r)
			logger.Errors(key, err)
			continue
		}

		if csvData.DtUpdated.After(dtLastUpdate) {
			dtLastUpdate = csvData.DtUpdated
		}

		logger.Infof("%s time = %d milliseconds", value, time.Since(timeStart).Milliseconds())
	}

	r.Patients = getEmptyPatients(dtLastUpdate)

	r.LastUpdate = dtLastUpdate.Format("2006/01/02 15:04")

	return r, nil
}

func getCSVDataFrame(apiAddress string, apiId string, csvAccessor accessor) (*CsvData, error) {
	data := mapCSVDataBackup[apiId]
	var err error
	if data == nil {
		data = &CsvData{}
		data.DfCsv, data.DtUpdated, err = csvAccessor.GetCSVDataFrameFromApi(apiAddress, apiId)
		mapCSVDataBackup[apiId] = data
	}
	return data, err
}

func setEmptyStructToResult(key string, r *Result) {
	switch key {
	case "main_summary":
		r.MainSummary = &MainSummary{}
	case "patients":
		r.Patients = &Patients{}
	case "patients_summary":
		r.PatientsSummary = &PatientsSummary{}
	case "inspection_persons":
		r.InspectionPersons = &InspectionPersons{}
	case "contacts":
		r.Contacts = &Contacts{}
	default:
		logger.Errors(key, "key not supported...")
	}
}
