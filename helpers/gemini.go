package helpers

// func GenerateWithGemini(prompt string) (string, error) {
// 	apiKey := os.Getenv("GEMINI_API_KEY")
func GenerateWithGemini(prompt string) (string, error) {
	return `
	Day 1: Explore Myeongdong & N Seoul Tower
	Day 2: Visit Gyeongbokgung Palace
	Day 3: Shopping & return
	`, nil
}

// url := fmt.Sprintf(
// 	"https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=%s",
// 	apiKey,
// )

// requestBody := map[string]interface{}{
// 	"contents": []map[string]interface{}{
// 		{
// 			"parts": []map[string]string{
// 				{"text": prompt},
// 			},
// 		},
// 	},
// }

// jsonData, _ := json.Marshal(requestBody)

// resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
// if err != nil {
// 	return "", err
// }
// defer resp.Body.Close()

// body, _ := io.ReadAll(resp.Body)

// fmt.Println("STATUS:", resp.StatusCode)
// fmt.Println("BODY:", string(body))
// fmt.Println(url)

// type GeminiResponse struct {
// 	Candidates []struct {
// 		Content struct {
// 			Parts []struct {
// 				Text string `json:"text"`
// 			} `json:"parts"`
// 		} `json:"content"`
// 	} `json:"candidates"`
// }

// var result GeminiResponse
// err = json.Unmarshal(body, &result)
// if err != nil {
// 	return "", err
// }

// if len(result.Candidates) == 0 {
// 	return "", fmt.Errorf("no response from gemini")
// }

// return result.Candidates[0].Content.Parts[0].Text, nil
