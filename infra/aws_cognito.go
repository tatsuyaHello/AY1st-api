package infra

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// NewCognitoClient は Cognito URL Signer
func NewCognitoClient() *cognitoidentityprovider.CognitoIdentityProvider {
	session := GetAWSSession()
	svc := cognitoidentityprovider.New(session, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})

	return svc
}

// ListUsers はsubからcognitoユーザーを検索
func ListUsers(sub string) (*cognitoidentityprovider.ListUsersOutput, error) {
	cognito := NewCognitoClient()
	getAttributes := []string{}
	filter := "sub = \"" + sub + "\""

	input := cognitoidentityprovider.ListUsersInput{
		AttributesToGet: aws.StringSlice(getAttributes),
		UserPoolId:      aws.String(os.Getenv("COGNITO_USER_POOL_ID")),
		Filter:          aws.String(filter),
	}

	return cognito.ListUsers(&input)
}

// GetCognitoUserName はsubからcognitoユーザーを検索
func GetCognitoUserName(sub string) (string, error) {
	output, err := ListUsers(sub)
	if err != nil {
		return "", err
	}

	if len(output.Users) == 0 {
		return "", fmt.Errorf("not found userName(sub = %v)", sub)
	}
	userName := *output.Users[0].Username

	return userName, nil
}

// DeleteCognitoUser Cognitoのuserデータの削除
func DeleteCognitoUser(username string) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {
	cognito := NewCognitoClient()

	input := cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(os.Getenv("COGNITO_USER_POOL_ID")),
		Username:   aws.String(username),
	}

	return cognito.AdminDeleteUser(&input)
}
