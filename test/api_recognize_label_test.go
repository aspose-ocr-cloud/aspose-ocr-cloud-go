/*
Aspose OCR Cloud 5.0 API

Testing RecognizeLabelApiService

*/

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

func Test_asposeocrcloud_RecognizeLabelApiService(t *testing.T) {

	clientId := ConfigClientID
	clientSecret := ConfigClientSecret
	configuration := asposeocrcloud.NewConfiguration(clientId, clientSecret)
	apiClient := asposeocrcloud.NewAPIClient(configuration)

	t.Run("Test RecognizeLabelApiService", func(t *testing.T) {

		// t.Skip("skip test")  // remove to run test

		filePath := "../samples/street_sign_01.jpg" // Path to your file

		// Read your file data and convert it into base64 string
		fileBytes, err := os.ReadFile(filePath)
		require.Nil(t, err)
		require.NotNil(t, fileBytes)
		fileb64Encoded := base64.StdEncoding.EncodeToString(fileBytes)

		// Step 1: create request body and sent it to OCR Cloud to receive task ID
		requestBody := *asposeocrcloud.NewOCRRecognizeLabelBody(
			fileb64Encoded,
			*asposeocrcloud.NewOCRSettingsRecognizeLabelWithDefaults())
		taskId, httpRes, err := apiClient.RecognizeLabelApi.PostRecognizeLabel(context.Background()).OCRRecognizeLabelBody(requestBody).Execute()

		require.Nil(t, err)
		require.NotNil(t, taskId)
		assert.Equal(t, 200, httpRes.StatusCode)

		fmt.Printf("File successfully sent. Your TaskID is %s \n", taskId)

		// Step 2: request task results using task ID
		ocrResp, httpRes, err := apiClient.RecognizeLabelApi.GetRecognizeLabel(context.Background()).Id(taskId).Execute()

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

			resultFilePath := "../results/" + taskId + ".txt"
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
