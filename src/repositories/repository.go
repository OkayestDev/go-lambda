package repositories

import (
	"golambda/src/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
	Get func(pk string, sk string) T
	// GetAll func(pk string) []T
}

func mapper[T any](items []map[string]*dynamodb.AttributeValue) []T {
	var entities []T

	for i := 0; i < len(items); i++ {
		var entity T
		dynamodbattribute.UnmarshalMap(items[i], &entity)
		entities = append(entities, entity)
	}

	return entities
}

func Repository[T any](tableName string) RepositoryFuncs[T] {
	awsTableName := aws.String(tableName)

	return RepositoryFuncs[T]{
		Get: func(pk string, sk string) T {
			result, err := svc.GetItem(&dynamodb.GetItemInput{
				TableName: awsTableName,
				Key: map[string]*dynamodb.AttributeValue{
					"Pk": {
						S: aws.String(pk),
					},
					"Sk": {
						S: aws.String(sk),
					},
				},
			})

			if err != nil {
				panic(err)
			}
			items :=  []map[string]*dynamodb.AttributeValue { result.Item }
			return mapper[T](items)[0]
		},

		// GetAll: func(pk string) {
		// 	result, err := svc.Query(&dynamodb.QueryInput {
		// 		TableName: awsTableName,
		// 		KeyConditions: {
		// 			"Pk": {
		// 				S: aws.String(pk),
		// 			},
		// 		},
		// 	})

		// 	if err != nil {
		// 		panic(err)
		// 	}

		// 	return result
		// },
	}
}
