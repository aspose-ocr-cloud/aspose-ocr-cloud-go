# Aspose.OCR Cloud SDK for Go 23.12.0

[![License](https://img.shields.io/github/license/aspose-ocr-cloud/aspose-ocr-cloud-dotnet)](LICENSE)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/aspose-ocr-cloud/aspose-ocr-cloud-go)

[Aspose.OCR Cloud](https://products.aspose.cloud/ocr/) is an optical character recognition as a service. With it, you can easily add OCR functionality to almost any device or platform: cloud, web, PCs, netbooks, or even entry-level smartphones.

Our engine can read text from images, photos, screenshots and scanned PDFs in a wide variety of European, Cyrillic and Oriental fonts, returning results in the most popular document formats. Powerful built-in image processing filters based on neural networks automatically correct skewed and distorted images, automatically remove dirt, smudges, scratches, glare and other image defects that can affect recognition accuracy. To further improve the results, Aspose.OCR Cloud has a built-in spell checker that automatically replaces misspelled words and saves you the trouble of manually correcting the recognition results.

Even the complex recognition tasks can be done with a couple of API calls. To make interacting with Aspose.OCR Cloud services from Go applications even easier, we provide the software development kit (SDK) for Go. It handles all the routine operations such as establishing connections, sending API requests, and parsing responses, wrapping all these tasks into a few simple classes.

Aspose.OCR Cloud SDK for Go is open source under the MIT license. You can freely use it for any projects, including commercial and proprietary applications, as well as modify any part of its code.

## Try Online
[Image to Text](https://products.aspose.app/ocr/scan-image) | [Image to Searchable PDF](https://products.aspose.app/ocr/ocr-to-pdf) | [PDF OCR](https://products.aspose.app/ocr/pdf-ocr)| [Receipt Scanner](https://products.aspose.app/ocr/scan-receipt)
:---: | :---: | :---:| :---:
[![Scan Image](https://products.aspose.app/ocr/scan-image/img/ocr-recognize-48.png)](https://products.aspose.app/ocr/scan-image) | [![Image to Searchable PDF](https://products.aspose.app/ocr/scan-image/img/ocr-to-pdf-4-48.png)](https://products.aspose.app/ocr/ocr-to-pdf) | [![PDF OCR](https://products.aspose.app/ocr/scan-image/img/ocr-to-pdf-2-48.png)](https://products.aspose.app/ocr/pdf-ocr) | [![Receipt Scanner](https://products.aspose.app/ocr/scan-image/img/aspose-scan-receipt-48.png)](https://products.aspose.app/ocr/scan-receipt) 

## System requirements

- Go 1.18 or later.
- Internet connection.
- Access to the **api.aspose.cloud** domain.

Check [go.mod](go.mod) file for the full list of third-party dependencies.

## Get started

Aspose.OCR Cloud is an on-demand service with a free tier. In order to use Aspose.OCR Cloud service, you must create an account at **Aspose Cloud API**:

1. Go to https://dashboard.aspose.cloud/
2. If you are already registered with Aspose, sign in with your user name and password.  
   Otherwise, click **Don’t have an account? Sign Up** link and create a new account.
3. [Check out more information](https://docs.aspose.cloud/ocr/subscription/) about available subscription plans and a free tier limits.

Aspose values your privacy and takes technical, security and organizational measures to protect your data from unauthorized use, accidental loss or disclosure. Read our [Privacy Policy](https://about.aspose.cloud/legal/privacy-policy) and [Terms of Service](https://about.aspose.cloud/legal/tos) for details.

### Authorization

Aspose.OCR Cloud follows industry standards and best practices to keep your data secure. All communication with OCR REST API is done using JWT authentication, which provides an open-standard, highly secure way to exchange information. Time-limited JWT tokens are generated using _Client ID_ and _Client Secret_ credentials that are specific for each application. To obtain the credentials:

1. Sign in to [Aspose Cloud API Dashboard](https://dashboard.aspose.cloud/).
2. Go to [**Applications**](https://dashboard.aspose.cloud/applications) page.
3. Click **Create New Application** button.
4. Give the application an easily recognizable name so it can be quickly found in a long list, and provide an optional detailed description.
5. Create the cloud storage by clicking the _plus_ icon and following the required steps. You can also reuse existing storage, if available.   
   Aspose.OCR Cloud uses its own internal storage, so you can provide the bare minimum storage options:

    - Type: **Internal storage**
    - Storage name: _Any name you like_
    - Storage mode: **Retain files for 24 hours**

6. Click **Save** button.
7. Click the newly created application and copy the values from **Client Id** and **Client Secret** fields.
8. Pass in the values from the **Client ID** and **Client Secret** fields when initializing the required OCR API.

### Running demo

1. Clone this repository or [download](https://github.com/aspose-ocr-cloud/aspose-ocr-cloud-go/archive/refs/heads/master.zip) it as ZIP.
3. Install **Aspose.OCR Cloud SDK for Go** package:
```shell
go get github.com/aspose-ocr-cloud/aspose-ocr-cloud-go
```
3. Open [examples/example.go](examples/example.go) file and replace "YOUR_CLIENT_ID" and "YOUR_CLIENT_SECRET" with credentials obtained during **Authorization** phase:
```go 
func main() {

	clientId := "YOUR_CLIENT_ID"
	clientSecret := "YOUR_CLIENT_SECRET"
```
4. Open the terminal and navigate to the _example_ directory of the downloaded repository.
5. Run the example:
```shell
go run .\example.go
```
6. Recognition results will be saved to the `results` directory of the downloaded repository.

### Running tests

We also provide automated tests for [Testify](https://github.com/stretchr/testify).

1. Clone this repository or [download](https://github.com/aspose-ocr-cloud/aspose-ocr-cloud-go/archive/refs/heads/master.zip) it as ZIP.
2. Install dependencies:
```shell
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/require
go get golang.org/x/oauth2
```
3. Open [test/test_config.go](test/test_config.go) file and replace "YOUR_CLIENT_ID" and "YOUR_CLIENT_SECRET" with credentials obtained during **Authorization** phase:
```go
package asposeocrcloud

var (
	ConfigClientID = "YOUR_CLIENT_ID"
	ConfigClientSecret = "YOUR_CLIENT_SECRET"
)
```
4. Open the terminal and navigate to the _test_ directory of the downloaded repository.
5. Run tests:
```
go test github.com/aspose-ocr-cloud/aspose-ocr-cloud-go/test
```
6. Test results will be saved to the `results` directory of the downloaded repository.

## What was changed in version 23.12.0

A summary of recent changes, enhancements and bug fixes in **Aspose.OCR Cloud SDK for .NET 23.12.0** release:

Key | Summary | Category
--- | ------- | --------
OCR&#8209;3737 | Added a free API for evaluating image recognition without [authorization](/ocr/authorization/).<br />Some restrictions apply. See below for details. | New feature

### Public API changes and backwards compatibility

This section lists all public API changes introduced in **Aspose.OCR Cloud SDK for .NET 23.12.0** that may affect the code of existing applications.

#### Added public APIs:

The following public APIs have been introduced in this release:

##### Image recognition evaluation

The following new classes have been added:

Class | Description
----- | -----------
`RecognizeImageTrialApi` | Image recognition API that works without authorization.

**Important:** In recognition results, 10% of the words are substituted with asterisks (`*`). The sequence of masked words remains unchanged upon re-submitting the identical image for recognition.

[Learn more...](https://docs.aspose.cloud/ocr/recognize-image/)

#### Updated public APIs:

_No changes_

#### Removed public APIs:

_No changes._

## Examples

The example below illustrates how to use Aspose.OCR Cloud SDK for Go to extract text from an image:

```go
package main


import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	asposeocrcloud "github.com/aspose-ocr-cloud/aspose-ocr-cloud-go"
)

func main(){
	
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
		if err != nil|| httpRes.StatusCode != 200 || ocrResp == nil {
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
```

## Other Aspose.OCR Cloud SDKs

- [Aspose.OCR Cloud for Go](https://github.com/aspose-ocr-cloud/aspose-ocr-cloud-dotnet/)  
  Extract text with Aspose.OCR Cloud API in your Go cloud, online and on-premises applications.
- [Aspose.OCR Cloud for Java](https://github.com/aspose-ocr-cloud/aspose-ocr-cloud-java)  
  Call Aspose.OCR Cloud API from cross-platform Java apps.
- [Aspose.OCR Cloud for Python](https://github.com/aspose-ocr-cloud/aspose-ocr-cloud-python)  
  Natively integrate OCR features into Python applications.
- [Aspose.OCR Cloud for Node.js](https://github.com/aspose-ocr-cloud/aspose-ocr-cloud-nodejs)  
  Add OCR functionality to AWS Lambda, Azure Functions, services, and applications written in Node.js by querying our REST API.
- [Aspose.OCR Cloud for Android](https://github.com/aspose-ocr-cloud/aspose-ocr-cloud-android)  
  Turn a smartphone into a full featured OCR scanner. Aspose.OCR service runs in the cloud and supports even entry-level and legacy smartphones.

## Resources

Find more information on Aspose.OCR Cloud and get professional help:

- [Documentation](https://docs.aspose.cloud/ocr/)
- [GitHub repositories](https://github.com/aspose-ocr-cloud)
- [Free Support Forum](https://forum.aspose.cloud/c/ocr/12)
- [Paid Support Helpdesk](https://helpdesk.aspose.cloud/)
- [Blog](https://blog.aspose.cloud/category/ocr/)
- [Free OCR Apps](https://products.aspose.app/ocr/family/)
- [Aspose Cloud Pricing](https://purchase.aspose.cloud/pricing)
- [Terms of Service](https://about.aspose.cloud/legal/tos/)
