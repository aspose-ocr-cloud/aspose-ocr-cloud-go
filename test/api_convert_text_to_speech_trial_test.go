/*
Aspose OCR Cloud 5.0 API

Testing ConvertTextToSpeechTrialApiService

*/


package asposeocrcloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func Test_asposeocrcloud_ConvertTextToSpeechTrialApiService(t *testing.T) {

	clientId := ""
	clientSecret := ""
	configuration := asposeocrcloud.NewConfiguration(clientId, clientSecret)
	apiClient := asposeocrcloud.NewAPIClient(configuration)

	t.Run("Test ConvertTextToSpeechTrialApiService", func(t *testing.T) {
		
		t.Skip("skip test")  // remove to run test

		textToSpeech := "This is a sample text"

		// Step 1: create request body and sent it to OCR Cloud to receive task ID
		requestBody := *asposeocrcloud.NewTTSBody(
			textToSpeech,
			*asposeocrcloud.NewTTSSettings(
				asposeocrcloud.LANGUAGETTS_ENGLISH,
				asposeocrcloud.RESULTTYPETTS_WAV,
			),
		)

		taskId, httpRes, err := apiClient.ConvertTextToSpeechTrialApi.PostConvertTextToSpeechTrial(context.Background()).TTSBody(requestBody).Execute()

		require.Nil(t, err)
		require.NotNil(t, taskId)
		assert.Equal(t, 200, httpRes.StatusCode)

		fmt.Printf("File successfully sent. Your TaskID is %s", taskId)

		// Step 2: request task results using task ID
		ocrResp, httpRes, err := apiClient.ConvertTextToSpeechTrialApi.GetConvertTextToSpeechTrial(context.Background()).Id(taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, ocrResp)
		assert.Equal(t, 200, httpRes.StatusCode)
		
		if *ocrResp.TaskStatus == asposeocrcloud.TTSTASKSTATUS_COMPLETED{
			require.NotNil(t, ocrResp.Results[0].Data)
			
			// Decode results and write to file
			decodedBytes, err := base64.StdEncoding.DecodeString(*ocrResp.Results[0].Data.Get())
			if err != nil {
				fmt.Println("Decode error:", err)
				return
			}
	
			resultFilePath := "../results/" + taskId + ".wav"
			err = os.WriteFile(resultFilePath, decodedBytes, 0644)
			if err != nil {
				fmt.Println("Write file error:", err)
				return
			}
	
			fmt.Printf("Task result successfully saved at %s \n", resultFilePath)
		}else{			
			fmt.Printf("Sorry, task %s is not completed yet. You can request results later. Task status: %s\n", taskId, *ocrResp.TaskStatus)
		}
	})

}
