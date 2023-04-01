package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"golambda/src/constants"
)

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
	Config: aws.Config{
		Endpoint: &constants.DYNAMO_ENDPOINT,
	},
}))

// Create DynamoDB client
var svc = dynamodb.New(sess)

type RepositoryFuncs[T any] struct {
	Get func(pk string, sk string) *dynamodb.GetItemOutput
}

func Repository[T any](tableName string) RepositoryFuncs[T] {
	var repo = RepositoryFuncs[T]{
		Get: func(pk string, sk string) *dynamodb.GetItemOutput {
			result, err := svc.GetItem(&dynamodb.GetItemInput{
				TableName: aws.String(tableName),
				Key: map[string]*dynamodb.AttributeValue{
					"Pk": {
						N: aws.String(pk),
					},
					"Sk": {
						S: aws.String(sk),
					},
				},
			})

			if err != nil {
				panic(err)
			}

			return result
		},
	}

	return repo
}
