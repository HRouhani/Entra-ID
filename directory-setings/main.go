package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-beta-sdk-go"
)

func main() {
	// Azure AD application credentials
	tenantID := "d.....2"
	clientID := "c.....2"
	clientSecret := ""

	// Step 1: Create a client secret credential for Azure AD authentication
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		log.Fatalf("Failed to create client secret credential: %v", err)
	}

	// Step 2: Create a GraphServiceClient with the credential and desired scope
	graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"https://graph.microsoft.com/.default"})
	if err != nil {
		log.Fatalf("Failed to create GraphServiceClient: %v", err)
	}

	// Step 3: Make a request to get directory setting templates
	directorySettingTemplates, err := graphClient.DirectorySettingTemplates().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to retrieve directory setting templates: %v", err)
	}

	// Step 4: Display settings for all the specified templates
	for _, template := range directorySettingTemplates.GetValue() {
		displayName := template.GetDisplayName()
		if displayName != nil {
			fmt.Printf("\nTemplate: %s\n", *displayName)
			fmt.Println("----------------------------------------------------")

			// Iterate through the settings for each template
			for _, setting := range template.GetValues() {
				name := setting.GetName()
				value := setting.GetDefaultValue()
				description := setting.GetDescription()
				if name != nil && value != nil && description != nil {
					fmt.Printf("Setting: %s\n  Default Value: %s\n  Description: %s\n", *name, *value, *description)
				} else {
					fmt.Printf("Setting: %s\n  Default Value: (nil)\n  Description: (nil)\n", *name)
				}
			}
			fmt.Println("----------------------------------------------------")
		}
	}
}
