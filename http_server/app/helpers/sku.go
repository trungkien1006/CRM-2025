package helpers

func GenerateSKU(upc string, ctsp_id int, counter int64) string {
	nowDate := GetCurrentTimeVN().Format("2006-01-02")

	return upc + "-" + string(ctsp_id) + "-" + nowDate + "-" + string(counter)
}