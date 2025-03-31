package business_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tapiaw38/reservation-service-be/internal/adapters/datasources/repositories/business"
	domain "github.com/tapiaw38/reservation-service-be/internal/domain/business"
	"github.com/tapiaw38/reservation-service-be/internal/platform/nosql/mocks"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	type fields struct {
		client *mocks.MockClient
	}

	tests := map[string]struct {
		prepare   func(*fields)
		input     domain.Business
		expect    string
		expectErr error
	}{
		"when everything works fine": {
			prepare: func(f *fields) {
				f.client.EXPECT().
					InsertOne(gomock.Any(), gomock.Any()).
					Return(&mongo.InsertOneResult{InsertedID: "some-id"}, nil)
			},
			input: domain.Business{
				ID:   "some-id",
				Name: "Test Business",
			},
			expect:    "some-id",
			expectErr: nil,
		},
		"when insertion fails": {
			prepare: func(f *fields) {
				f.client.EXPECT().
					InsertOne(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("insertion error"))
			},
			input: domain.Business{
				ID:   "some-id",
				Name: "Test Business",
			},
			expect:    "",
			expectErr: errors.New("insertion error"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := &fields{
				client: mocks.NewMockClient(ctrl),
			}

			tc.prepare(f)

			repo := business.NewRepository(f.client)
			actual, actualErr := repo.Create(context.Background(), tc.input)

			assert.Equal(t, tc.expect, actual)
			assert.Equal(t, tc.expectErr, actualErr)
		})
	}
}
