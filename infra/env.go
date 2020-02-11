package infra

import "os"

// GetEnvcode Envcodeを返す
func GetEnvcode() string {
	return os.Getenv("ENVCODE")
}

// IsLocal Envcode が localかどうか
func IsLocal() bool {
	return GetEnvcode() == "local" || GetEnvcode() == "test"
}

// IsProduction Envcode が 本番環境かどうか
func IsProduction() bool {
	return GetEnvcode() == "prd"
}

// GetEnvName は環境名を取得
func GetEnvName() string {
	envcode := GetEnvcode()
	switch envcode {
	case "prd":
		return "本番環境"
	case "stg":
		return "ステージ環境"
	case "dev":
		return "開発環境"
	case "local":
		return "ローカル環境"
	}
	return "不明"
}
