/*
Aspose OCR Cloud 5.0 API

Testing RecognizeRegionsApiService

*/

package asposeocrcloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"

	asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_asposeocrcloud_RecognizeRegionsApiService(t *testing.T) {

	clientId := ConfigClientID
	clientSecret := ConfigClientSecret
	configuration := asposeocrcloud.NewConfiguration(clientId, clientSecret)
	apiClient := asposeocrcloud.NewAPIClient(configuration)

	t.Run("Test RecognizeRegionsApiService", func(t *testing.T) {

		// t.Skip("skip test")  // remove to run test

		filePath := "../samples/greek.PNG" // Path to your file

		// Read your file data and convert it into base64 string
		fileBytes, err := ioutil.ReadFile(filePath)
		require.Nil(t, err)
		require.NotNil(t, fileBytes)
		fileb64Encoded := base64.StdEncoding.EncodeToString(fileBytes)

		// Step 1: create request body and sent it to OCR Cloud to receive task ID
		recognitionSettings := *asposeocrcloud.NewOCRSettingsRecognizeRegionsWithDefaults()
		recognitionSettings.Language = asposeocrcloud.LANGUAGE_GREEK.Ptr()
		recognitionSettings.DsrMode = asposeocrcloud.DSRMODE_REGIONS.Ptr()
		*recognitionSettings.MakeBinarization = false
		*recognitionSettings.MakeSkewCorrect = false
		*recognitionSettings.MakeUpsampling = false
		*recognitionSettings.MakeSpellCheck = false
		*recognitionSettings.MakeContrastCorrection = false
		recognitionSettings.ResultType = asposeocrcloud.RESULTTYPE_TEXT.Ptr()

		// Define zone of image to detect text. To recognize full document, you may omit this part
		regionOrderValue := int32(0) // order of recognition region
		topLeftXValue := int32(117) // bounding box pixel coordinates of image part to recognize
		TopLeftYValue := int32(81)
		BottomRightXValue := int32(611)
		BottomRightYValue := int32(98)
		recognitionRegionRectZone := asposeocrcloud.NewOCRRect()
		recognitionRegionRectZone.TopLeftX = &topLeftXValue
		recognitionRegionRectZone.TopLeftY = &TopLeftYValue
		recognitionRegionRectZone.BottomRightX = &BottomRightXValue
		recognitionRegionRectZone.BottomRightY = &BottomRightYValue
		recognitionRegion := *asposeocrcloud.NewOCRRegion()
		recognitionRegion.Rect = recognitionRegionRectZone
		recognitionRegion.Order = &regionOrderValue
		recognitionSettings.Regions = append(recognitionSettings.Regions, recognitionRegion)

		requestBody := *asposeocrcloud.NewOCRRecognizeRegionsBody(
			fileb64Encoded,
			recognitionSettings,
		)
		taskId, httpRes, err := apiClient.RecognizeRegionsApi.PostRecognizeRegions(context.Background()).OCRRecognizeRegionsBody(requestBody).Execute()

		require.Nil(t, err)
		require.NotNil(t, taskId)
		assert.Equal(t, 200, httpRes.StatusCode)

		fmt.Printf("File successfully sent. Your TaskID is %s \n", taskId)

		// Step 2: request task results using task ID
		ocrResp, httpRes, err := apiClient.RecognizeRegionsApi.GetRecognizeRegions(context.Background()).Id(taskId).Execute()

		require.Nil(t, err)
		require.NotNil(t, ocrResp)
		assert.Equal(t, 200, httpRes.StatusCode)
		if *ocrResp.TaskStatus == asposeocrcloud.OCRTASKSTATUS_COMPLETED {
			require.NotNil(t, ocrResp.Results[0].Data)

			// Decode results and write to file
			// The detected regions are presented as a text string containing an entry of the form [[topLeftPixelX, topLeftPixelY, bottomRightPixelX, bottomRightPixelY], ...]
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

	})

}
