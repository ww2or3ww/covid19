package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"app/csv2json"
	"app/utils/logger"
)

// クエリパラメータが無かった場合のデフォルト
const defaultTypes = "patients_summary:221309_hamamatsu_covid19_patients_summary,inspection_persons:221309_hamamatsu_covid19_test_people,contacts:221309_hamamatsu_covid19_call_center"

// APIアドレス
const opendataApiUrl = "https://www.city.hamamatsu.shizuoka.jp/api/odpf/opendata/v1"

type Csv2Json interface {
	Process(apiAddress string, queryStrPrm string) (*csv2json.Result, error)
}

var c2j Csv2Json

// AWS Lambda エンドポイント
func handler(filePath string) (bool, error) {
	timeStart := time.Now()

	// クエリパラメータ取得
	queryStrPrm := defaultTypes

	// csv2json
	mapData, err := c2j.Process(opendataApiUrl, queryStrPrm)
	if err != nil {
		return false, err
	}

	// mapをインデント付きのJSONに整形してBodyとして返す
	jsonIndent, err := json.MarshalIndent(mapData, "", "   ")
	if err != nil {
		return false, err
	}

	logger.Debugs(string(jsonIndent))

	if len(filePath) > 0 {
		ioutil.WriteFile(filePath, jsonIndent, 0644)
	}

	logger.Infof("total time = %d milliseconds", time.Since(timeStart).Milliseconds())

	return true, nil
}

// mainメソッドの前に呼ばれる初期化処理
func init() {
	// LogLvを環境変数から取得してLog初期設定する
	logLv := logger.Error
	envLogLv := os.Getenv("LOG_LEVEL")
	logger.Infos(envLogLv)
	if envLogLv != "" {
		n, _ := strconv.Atoi(envLogLv)
		logLv = logger.LogLv(n)
	}
	logger.LogInitialize(logLv, 25)

	// 本番用のCSV2JSONをDIしておく
	c2j = csv2json.NewCsv2Json(csv2json.NewCsvAccessor())
}

// アプリケーションエンドポイント
func main() {
	if os.Getenv("LOG_LEVEL") == "" {
		godotenv.Load(".env")
	}
	logger.Infos("=== START ===")
	handler(os.Args[1])
	logger.Infos("=== COMPLETED ===")
}
