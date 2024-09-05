# Retrieve Microsoft Entra ID Directory Setting Templates with Go

This Go program uses the Microsoft Graph API to authenticate with Microsoft Entra ID (Azure AD) and retrieve various directory setting templates. These templates provide configurations and settings that define policies for different aspects of the directory, such as user consent, group management, application settings, and more.

## Features

The program retrieves and displays the settings and their descriptions for the following templates:

- **Group.Unified.Guest**: Settings for managing guest users in unified groups.
- **Application**: Settings related to application management.
- **Password Rule Settings**: Settings defining password policies.
- **Group.Unified**: Settings for unified group management.
- **Prohibited Names Settings**: Templates defining reserved or prohibited names in the directory.
- **Custom Policy Settings**: Templates for custom policies within the directory.
- **Prohibited Names Restricted Settings**: Additional settings for managing prohibited names.
- **Consent Policy Settings**: Settings related to user consent policies for applications.

## Prerequisites

- Go 1.16 or later.
- Azure Active Directory (Entra ID) APP with credentials:
  - Tenant ID
  - Client ID
  - Client Secret
- Microsoft Graph API permissions: `Directory.Read.All` or `Directory.ReadWrite.All`.

## Setup

1. **Clone the Repository:**

    ```sh
    git clone https://github.com/yourusername/your-repo-name.git
    cd your-repo-name
    ```

2. **Install Dependencies:**

    Ensure you have the necessary Go packages installed:

    ```sh
    go mod tidy
    ```

3. **Set Your Azure AD Credentials:**

    Open the `main.go` file and update the following variables with your Azure AD credentials:

    ```go
    tenantID := "your-tenant-id"
    clientID := "your-client-id"
    clientSecret := "your-client-secret"
    ```

## Usage

Run the program to retrieve the directory setting templates:

```sh
go run main.go
