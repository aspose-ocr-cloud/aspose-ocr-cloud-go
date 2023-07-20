package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main() {

	clientId := "YOUR_CLIENT_ID"
	clientSecret := "YOUR_CLIENT_SECRET"
	configuration := asposeocrcloud.NewConfiguration(clientId, clientSecret)
	apiClient := asposeocrcloud.NewAPIClient(configuration)

	filePath := "../samples/latin.png" // Path to your file

	// Read your file data and convert it into base64 string
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil || fileBytes == nil {
		fmt.Println("Read file error:", err)
		return
	}
	fileb64Encoded := base64.StdEncoding.EncodeToString(fileBytes)

	// Step 1: create request body and sent it to OCR Cloud to receive task ID
	recognitionSettings := *asposeocrcloud.NewOCRSettingsRecognizeImage()
	recognitionSettings.Language = asposeocrcloud.LANGUAGE_ENGLISH.Ptr()
	recognitionSettings.DsrMode = asposeocrcloud.DSRMODE_NO_DSR_NO_FILTER.Ptr()
	recognitionSettings.DsrConfidence = asposeocrcloud.DSRCONFIDENCE_DEFAULT.Ptr()
	*recognitionSettings.MakeBinarization = false
	*recognitionSettings.MakeSkewCorrect = false
	*recognitionSettings.MakeUpsampling = false
	*recognitionSettings.MakeSpellCheck = false
	*recognitionSettings.MakeContrastCorrection = false
	recognitionSettings.ResultType = asposeocrcloud.RESULTTYPE_TEXT.Ptr()
	recognitionSettings.ResultTypeTable = asposeocrcloud.RESULTTYPETABLE_TEXT.Ptr()

	requestBody := *asposeocrcloud.NewOCRRecognizeImageBody(
		fileb64Encoded,
		recognitionSettings,
	)

	taskId, httpRes, err := apiClient.RecognizeImageApi.PostRecognizeImage(context.Background()).OCRRecognizeImageBody(requestBody).Execute()
	if err != nil || httpRes.StatusCode != 200 {
		fmt.Println("API error:", err)
		return
	}

	fmt.Printf("File successfully sent. Your TaskID is %s \n", taskId)

	// Step 2: request task results using task ID
	ocrResp, httpRes, err := apiClient.RecognizeImageApi.GetRecognizeImage(context.Background()).Id(taskId).Execute()
	if err != nil || httpRes.StatusCode != 200 || ocrResp == nil {
		fmt.Println("API error:", err)
		return
	}

	if *ocrResp.TaskStatus == asposeocrcloud.OCRTASKSTATUS_COMPLETED {
		if !ocrResp.Results[0].Data.IsSet() {
			fmt.Println("Response is empty")
			return
		}

		// Decode results and write to file
		decodedBytes, err := base64.StdEncoding.DecodeString(*ocrResp.Results[0].Data.Get())
		if err != nil {
			fmt.Println("Decode error:", err)
			return
		}

		resultFilePath := "../results/" + taskId + ".txt"
		err = ioutil.WriteFile(resultFilePath, decodedBytes, 0644)
		if err != nil {
			fmt.Println("Write file error:", err)
			return
		}

		fmt.Printf("Task result successfully saved at %s \n", resultFilePath)
	} else {
		fmt.Printf("Sorry, task %s is not completed yet. You can request results later. Task status: %s\n", taskId, *ocrResp.TaskStatus)
	}
}
