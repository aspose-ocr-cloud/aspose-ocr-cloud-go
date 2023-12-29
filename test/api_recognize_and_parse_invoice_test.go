
package asposeocrcloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_asposeocrcloud_RecognizeAndParseInvoiceApiService(t *testing.T) {

	clientId := ConfigClientID
	clientSecret := ConfigClientSecret
	configuration := asposeocrcloud.NewConfiguration(clientId, clientSecret)
	apiClient := asposeocrcloud.NewAPIClient(configuration)

	t.Run("Test RecognizeAndParseInvoiceApiService", func(t *testing.T) {

		// t.Skip("skip test")  // remove to run test

		filePath := "../samples/invoice_english_01.jpg" // Path to your file

		// Read your file data and convert it into base64 string
		fileBytes, err := os.ReadFile(filePath)
		require.Nil(t, err)
		require.NotNil(t, fileBytes)
		fileb64Encoded := base64.StdEncoding.EncodeToString(fileBytes)

		// Step 1: create request body and sent it to OCR Cloud to receive task ID
		recognitionSettings := *asposeocrcloud.NewOCRSettingsRecognizeAndParseInvoice()
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

		requestBody := *asposeocrcloud.NewOCRRecognizeAndParseInvoiceBody(
			fileb64Encoded,
			recognitionSettings,
		)
		taskId, httpRes, err := apiClient.RecognizeAndParseInvoiceApi.PostRecognizeAndParseInvoice(context.Background()).OCRRecognizeAndParseInvoiceBody(requestBody).Execute()

		require.Nil(t, err)
		require.NotNil(t, taskId)
		assert.Equal(t, 200, httpRes.StatusCode)

		fmt.Printf("File successfully sent. Your TaskID is %s \n", taskId)

		// Step 2: request task results using task ID
		ocrResp, httpRes, err := apiClient.RecognizeAndParseInvoiceApi.GetRecognizeAndParseInvoice(context.Background()).Id(taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, ocrResp)
		assert.Equal(t, 200, httpRes.StatusCode)
		if *ocrResp.TaskStatus == asposeocrcloud.OCRTASKSTATUS_COMPLETED {
			require.NotNil(t, ocrResp.Results[0].Data)

			// Decode results and write to file
			decodedBytes, err := base64.StdEncoding.DecodeString(*ocrResp.Results[0].Data.Get())
			if err != nil {
				fmt.Println("Decode error:", err)
				return
			}

			resultFilePath := "../results/" + taskId + ".json"
			err = os.WriteFile(resultFilePath, decodedBytes, 0644)
			if err != nil {
				fmt.Println("Write file error:", err)
				return
			}

			fmt.Printf("Task result successfully saved at %s \n", resultFilePath)
		} else {
			fmt.Printf("Sorry, task %s is not completed yet. You can request results later. Task status: %s\n", taskId, *ocrResp.TaskStatus)
		}

	})

}
