package architect_grammar

import (
	"context"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-go/v112/platformclientv2"
	"log"
	"net/http"
	"os"
)

/*
The genesyscloud_architect_grammar_proxy.go file contains the proxy structures and methods that interact
with the Genesys Cloud SDK. We use composition here for each function on the proxy so individual functions can be stubbed
out during testing.
*/

// internalProxy holds a proxy instance that can be used throughout the package
var internalProxy *architectGrammarProxy

// Type definitions for each func on our proxy so we can easily mock them out later
type createArchitectGrammarFunc func(ctx context.Context, p *architectGrammarProxy, grammar *platformclientv2.Grammar) (*platformclientv2.Grammar, error)
type getAllArchitectGrammarFunc func(ctx context.Context, p *architectGrammarProxy) (*[]platformclientv2.Grammar, error)
type getArchitectGrammarByIdFunc func(ctx context.Context, p *architectGrammarProxy, grammarId string) (grammar *platformclientv2.Grammar, responseCode int, err error)
type getArchitectGrammarIdByNameFunc func(ctx context.Context, p *architectGrammarProxy, name string) (grammarId string, retryable bool, err error)
type updateArchitectGrammarFunc func(ctx context.Context, p *architectGrammarProxy, grammarId string, grammar *platformclientv2.Grammar) (*platformclientv2.Grammar, error)
type deleteArchitectGrammarFunc func(ctx context.Context, p *architectGrammarProxy, grammarId string) (responseCode int, err error)

// architectGrammarProxy contains all of the methods that call genesys cloud APIs.
type architectGrammarProxy struct {
	clientConfig                    *platformclientv2.Configuration
	architectApi                    *platformclientv2.ArchitectApi
	createArchitectGrammarAttr      createArchitectGrammarFunc
	getAllArchitectGrammarAttr      getAllArchitectGrammarFunc
	getArchitectGrammarByIdAttr     getArchitectGrammarByIdFunc
	getArchitectGrammarIdByNameAttr getArchitectGrammarIdByNameFunc
	updateArchitectGrammarAttr      updateArchitectGrammarFunc
	deleteArchitectGrammarAttr      deleteArchitectGrammarFunc
}

// newArchitectGrammarProxy initializes the grammar proxy with all of the data needed to communicate with Genesys Cloud
func newArchitectGrammarProxy(clientConfig *platformclientv2.Configuration) *architectGrammarProxy {
	api := platformclientv2.NewArchitectApiWithConfig(clientConfig)
	return &architectGrammarProxy{
		clientConfig:                    clientConfig,
		architectApi:                    api,
		createArchitectGrammarAttr:      createArchitectGrammarFn,
		getAllArchitectGrammarAttr:      getAllArchitectGrammarFn,
		getArchitectGrammarByIdAttr:     getArchitectGrammarByIdFn,
		getArchitectGrammarIdByNameAttr: getArchitectGrammarIdByNameFn,
		updateArchitectGrammarAttr:      updateArchitectGrammarFn,
		deleteArchitectGrammarAttr:      deleteArchitectGrammarFn,
	}
}

// getArchitectGrammarProxy acts as a singleton for the internalProxy. It also ensures
// that we can still proxy our tests by directly setting internalProxy package variable
func getArchitectGrammarProxy(clientConfig *platformclientv2.Configuration) *architectGrammarProxy {
	if internalProxy == nil {
		internalProxy = newArchitectGrammarProxy(clientConfig)
	}

	return internalProxy
}

// createArchitectGrammar creates a Genesys Cloud Architect Grammar
func (p *architectGrammarProxy) createArchitectGrammar(ctx context.Context, grammar *platformclientv2.Grammar) (*platformclientv2.Grammar, error) {
	return p.createArchitectGrammarAttr(ctx, p, grammar)
}

// getAllArchitectGrammar retrieves all Genesys Cloud Architect Grammar
func (p *architectGrammarProxy) getAllArchitectGrammar(ctx context.Context) (*[]platformclientv2.Grammar, error) {
	return p.getAllArchitectGrammarAttr(ctx, p)
}

// getArchitectGrammarById returns a single Genesys Cloud Architect Grammar by Id
func (p *architectGrammarProxy) getArchitectGrammarById(ctx context.Context, grammarId string) (grammar *platformclientv2.Grammar, statusCode int, err error) {
	return p.getArchitectGrammarByIdAttr(ctx, p, grammarId)
}

// getArchitectGrammarIdByName returns a single Genesys Cloud Architect Grammar by a name
func (p *architectGrammarProxy) getArchitectGrammarIdByName(ctx context.Context, name string) (grammarId string, retryable bool, err error) {
	return p.getArchitectGrammarIdByNameAttr(ctx, p, name)
}

// updateArchitectGrammar updates a Genesys Cloud Architect Grammar
func (p *architectGrammarProxy) updateArchitectGrammar(ctx context.Context, grammarId string, grammar *platformclientv2.Grammar) (*platformclientv2.Grammar, error) {
	return p.updateArchitectGrammarAttr(ctx, p, grammarId, grammar)
}

// deleteArchitectGrammar deletes a Genesys Cloud Architect Grammar by Id
func (p *architectGrammarProxy) deleteArchitectGrammar(ctx context.Context, grammarId string) (statusCode int, err error) {
	return p.deleteArchitectGrammarAttr(ctx, p, grammarId)
}

// createArchitectGrammarFn is an implementation function for creating a Genesys Cloud Architect Grammar
func createArchitectGrammarFn(ctx context.Context, p *architectGrammarProxy, grammar *platformclientv2.Grammar) (*platformclientv2.Grammar, error) {
	grammarSdk, _, err := p.architectApi.PostArchitectGrammars(*grammar)
	if err != nil {
		return nil, fmt.Errorf("Failed to create grammar: %s", err)
	}

	// Create each language associated with the grammar
	for _, language := range *grammar.Languages {
		_, _, err := p.architectApi.PostArchitectGrammarLanguages(*grammarSdk.Id, language)
		if err != nil {
			return grammarSdk, fmt.Errorf("Failed to create grammar language: %s", err)
		}

		// Upload grammar voice file
		if language.VoiceFileMetadata != nil && language.VoiceFileMetadata.FileName != nil {
			uploadRequest := platformclientv2.Grammarfileuploadrequest{
				FileType: language.VoiceFileMetadata.FileType,
			}
			err = uploadGrammarLanguageFile(p, *grammarSdk.Id, *language.Language, language.VoiceFileMetadata.FileName, &uploadRequest, "voice")
			if err != nil {
				return grammarSdk, fmt.Errorf("Failed to upload language voice file: %s", err)
			}
		}

		// Upload grammar dtmf file
		if language.DtmfFileMetadata != nil && language.DtmfFileMetadata.FileName != nil {
			uploadRequest := platformclientv2.Grammarfileuploadrequest{
				FileType: language.DtmfFileMetadata.FileType,
			}
			err = uploadGrammarLanguageFile(p, *grammarSdk.Id, *language.Language, language.DtmfFileMetadata.FileName, &uploadRequest, "dtmf")
			if err != nil {
				return grammarSdk, fmt.Errorf("Failed to upload language dtmf file: %s", err)
			}
		}
	}

	return grammarSdk, nil
}

// getAllArchitectGrammarFn is the implementation for retrieving all Architect Grammars in Genesys Cloud
func getAllArchitectGrammarFn(ctx context.Context, p *architectGrammarProxy) (*[]platformclientv2.Grammar, error) {
	var allGrammars []platformclientv2.Grammar

	for pageNum := 1; ; pageNum++ {
		const pageSize = 100

		grammars, _, err := p.architectApi.GetArchitectGrammars(pageNum, pageSize, "", "", []string{}, "", "", "", true)
		if err != nil {
			return nil, fmt.Errorf("Failed to get architect grammars: %v", err)
		}

		if grammars.Entities == nil || len(*grammars.Entities) == 0 {
			break
		}

		for _, grammar := range *grammars.Entities {
			log.Printf("Dealing with grammar id : %s", *grammar.Id)
			allGrammars = append(allGrammars, grammar)
		}
	}

	return &allGrammars, nil
}

// getArchitectGrammarByIdFn is an implementation of the function to get a Genesys Cloud Architect Grammar by Id
func getArchitectGrammarByIdFn(ctx context.Context, p *architectGrammarProxy, grammarId string) (grammar *platformclientv2.Grammar, statusCode int, err error) {
	grammar, resp, err := p.architectApi.GetArchitectGrammar(grammarId, true)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("Failed to retrieve grammar by id %s: %s", grammarId, err)
	}
	return grammar, resp.StatusCode, nil
}

// getArchitectGrammarIdByNameFn is an implementation of the function to get a Genesys Cloud Architect Grammar by name
func getArchitectGrammarIdByNameFn(ctx context.Context, p *architectGrammarProxy, name string) (grammarId string, retryable bool, err error) {
	const pageNum = 1
	const pageSize = 100
	grammars, _, err := p.architectApi.GetArchitectGrammars(pageNum, pageSize, "", "", []string{}, name, "", "", true)
	if err != nil {
		return "", false, fmt.Errorf("Error searching architect grammar %s: %s", name, err)
	}

	if grammars.Entities == nil || len(*grammars.Entities) == 0 {
		return "", true, fmt.Errorf("No architect grammars found with name %s", name)
	}

	var grammar platformclientv2.Grammar
	for _, grammarSdk := range *grammars.Entities {
		if *grammarSdk.Name == name {
			log.Printf("Retrieved the grammar id %s by name %s", *grammarSdk.Id, name)
			grammar = grammarSdk
			return *grammar.Id, false, nil
		}
	}

	return "", false, fmt.Errorf("Unable to find grammar with name %s", name)
}

// updateArchitectGrammarFn is an implementation of the function to update a Genesys Cloud Architect Grammar
func updateArchitectGrammarFn(ctx context.Context, p *architectGrammarProxy, grammarId string, grammar *platformclientv2.Grammar) (*platformclientv2.Grammar, error) {
	grammarSdk, _, err := p.architectApi.PatchArchitectGrammar(grammarId, *grammar)
	if err != nil {
		return nil, fmt.Errorf("Failed to update grammar %s: %s", grammarId, err)
	}

	// Update each language associated with the grammar
	for _, language := range *grammar.Languages {
		_, err := updateArchitectGrammarLanguage(ctx, p, *grammarSdk.Id, *language.Language, &language)
		if err != nil {
			return grammarSdk, fmt.Errorf("Failed to update grammar language: %s", err)
		}

		// Upload grammar voice file
		if language.VoiceFileMetadata != nil && language.VoiceFileMetadata.FileName != nil {
			uploadRequest := platformclientv2.Grammarfileuploadrequest{
				FileType: language.VoiceFileMetadata.FileType,
			}
			err = uploadGrammarLanguageFile(p, *grammarSdk.Id, *language.Language, language.VoiceFileMetadata.FileName, &uploadRequest, "voice")
			if err != nil {
				return grammarSdk, fmt.Errorf("Failed to upload language voice file: %s", err)
			}
		}

		// Upload grammar dtmf file
		if language.DtmfFileMetadata != nil && language.DtmfFileMetadata.FileName != nil {
			uploadRequest := platformclientv2.Grammarfileuploadrequest{
				FileType: language.DtmfFileMetadata.FileType,
			}
			err = uploadGrammarLanguageFile(p, *grammarSdk.Id, *language.Language, language.DtmfFileMetadata.FileName, &uploadRequest, "dtmf")
			if err != nil {
				return grammarSdk, fmt.Errorf("Failed to upload language dtmf file: %s", err)
			}
		}
	}

	return grammarSdk, nil
}

// updateArchitectGrammarLanguage is a function for updating a Genesys Cloud Architect Grammarlanguage
func updateArchitectGrammarLanguage(ctx context.Context, p *architectGrammarProxy, grammarId string, languageCode string, language *platformclientv2.Grammarlanguage) (*platformclientv2.Grammarlanguage, error) {
	// Need to check if the language exists, if not, create it
	_, _, err := p.architectApi.GetArchitectGrammarLanguage(grammarId, languageCode)
	if err != nil {
		log.Printf("Laguage %s not found, creating language", languageCode)
		language, _, err := p.architectApi.PostArchitectGrammarLanguages(grammarId, *language)
		if err != nil {
			return nil, fmt.Errorf("Failed to create grammar language: %s", err)
		}

		return language, nil
	}

	languageUpdate := platformclientv2.Grammarlanguageupdate{
		DtmfFileMetadata:  language.DtmfFileMetadata,
		VoiceFileMetadata: language.VoiceFileMetadata,
	}
	languageSDK, _, err := p.architectApi.PatchArchitectGrammarLanguage(grammarId, languageCode, languageUpdate)
	if err != nil {
		return nil, fmt.Errorf("Failed to update grammar language: %s", err)
	}

	return languageSDK, nil
}

// deleteArchitectGrammarFn is an implementation function for deleting a Genesys Cloud Architect Grammar
func deleteArchitectGrammarFn(ctx context.Context, p *architectGrammarProxy, grammarId string) (statusCode int, err error) {
	_, resp, err := p.architectApi.DeleteArchitectGrammar(grammarId)
	if err != nil {
		return resp.StatusCode, fmt.Errorf("Failed to delete grammar: %s", err)
	}

	return resp.StatusCode, nil
}

// uploadGrammarLanguageFile is a function for uploading a grammar language file to Genesys cloud
func uploadGrammarLanguageFile(p *architectGrammarProxy, grammarId string, languageCode string, filename *string, uploadBody *platformclientv2.Grammarfileuploadrequest, fileType string) error {
	var uploadResponse *platformclientv2.Uploadurlresponse
	var err error
	if fileType == "voice" {
		uploadResponse, _, err = p.architectApi.PostArchitectGrammarLanguageFilesVoice(grammarId, languageCode, *uploadBody)
	} else if fileType == "dtmf" {
		uploadResponse, _, err = p.architectApi.PostArchitectGrammarLanguageFilesDtmf(grammarId, languageCode, *uploadBody)
	} else {
		return fmt.Errorf("Invalid file type given. Specify either voice of dtmf")
	}
	if err != nil {
		return fmt.Errorf("Failed to get language file presignedUri: %s for file %s", err, *filename)
	}

	file, err := os.Open(*filename)
	if err != nil {
		return fmt.Errorf("Failed to find file: %s", err)
	}
	defer file.Close()

	request, err := http.NewRequest(http.MethodPut, *uploadResponse.Url, file)
	if err != nil {
		return err
	}

	for key, value := range *uploadResponse.Headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}
