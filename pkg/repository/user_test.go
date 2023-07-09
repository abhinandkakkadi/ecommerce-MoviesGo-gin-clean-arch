package repository

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_user_UserSignUp(t *testing.T) {

	type fields struct {
		db *gorm.DB
	}

	type args struct {
		input models.UserDetails
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       models.UserDetailsResponse
		wantErr    error
	}{
		{
			name: "success signup user",
			args: args{
				input: models.UserDetails{Name: "Abhinand", Email: "nanduttanvsabhi@gmail.com", Phone: "9961088604", Password: "132457689"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("Abhinand", "nanduttanvsabhi@gmail.com", "9961088604", "132457689").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, "Abhinand", "nanduttanvsabhi@gmail.com", "9961088604"))

			},

			want:    models.UserDetailsResponse{Id: 1, Name: "Abhinand", Email: "nanduttanvsabhi@gmail.com", Phone: "9961088604"},
			wantErr: nil,
		},
		{
			name: "duplicate user",
			args: args{
				input: models.UserDetails{Name: "Abhinand", Email: "nanduttanvsabhi@gmail.com", Phone: "9961088604", Password: "132457689"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("Abhinand", "nanduttanvsabhi@gmail.com", "9961088604", "132457689").
					WillReturnRows(sqlmock.NewRows([]string{}).AddRow()).
					WillReturnError(errors.New("email should be unique"))

			},

			want:    models.UserDetailsResponse{},
			wantErr: errors.New("email should be unique"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.beforeTest(mockSQL)

			u := NewUserRepository(gormDB)

			got, err := u.UserSignUp(tt.args.input)

			assert.Equal(t, tt.wantErr, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.UserSignUp() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_user_GetWalletDetails(t *testing.T) {

	type fields struct {
		db *gorm.DB
	}

	type args struct {
		userID int
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       models.Wallet
		wantErr    error
	}{
		{
			name: "success retrieved wallet",
			args: args{
				userID: 1,
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {

				columns := []string{"wallet_amount"}

				expectedQuery := `^select wallet_amount from wallets (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs(1).WillReturnRows(sqlmock.NewRows(columns).FromCSVString("10000"))

			},

			want:    models.Wallet{10000},
			wantErr: nil,
		},

		{
			name: "user_id does not exists",
			args: args{
				userID: 2,
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {

				columns := []string{"wallet_amount"}

				expectedQuery := `^select wallet_amount from wallets (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs(2).WillReturnRows(sqlmock.NewRows(columns).FromCSVString("10000")).WillReturnRows(sqlmock.NewRows([]string{}).AddRow()).
					WillReturnError(errors.New("user_id does not exists"))

			},

			want:    models.Wallet{},
			wantErr: errors.New("user_id does not exists"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.beforeTest(mockSQL)

			u := NewUserRepository(gormDB)

			got, err := u.GetWalletDetails(tt.args.userID)
			assert.Equal(t, tt.wantErr, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.UserSignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_user_GetAllAddresses(t *testing.T) {

// 	mockDb, mock, err := sqlmock.New()

// 	if assert.NoError(t, err) {
// 		log.Println("Mock DB created")
// 	}

// 	defer mockDb.Close()

// 	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDb, DriverName: "postgres"}), &gorm.Config{})
// 	if assert.NoError(t, err) {
// 		log.Println("Mock DB connected to gorm")
// 	}

// 	userRepo := NewUserRepository(db)

// 	userID := 1
// 	expectedQuery := `^select * from addresses (.+)$`
// 	rows := sqlmock.NewRows([]string{"id", "name", "house_name", "state","pin","street","city"}).
// 		AddRow("1", "Abhinand", "Kakkadi House", "Kerala","670562","Kanul","Kannur").
// 		AddRow("2", "Abhinand", "Kadankot Valappu", "Kerala","670567","Kanul","Kannur")

// 	mock.ExpectQuery(expectedQuery).WithArgs(userID).WillReturnRows(rows)

// 	addresses, err := userRepo.GetAllAddresses(userID)
// 	if err != nil {
// 		t.Fatalf("Failed to get addresses: %v", err)
// 	}

// 	expectedAddresses := []models.AddressInfoResponse{
// 		{ID: 1, Name: "Abhinand", HouseName: "Kakkadi House", State: "Kerala", Pin: "670562", Street: "Kanul",City: "Kannur"},
// 		{ID: 2, Name: "Abhinand", HouseName: "Kadankot Valappu", State: "Kerala", Pin: "670567", Street: "Kanul",City: "Kannur"},
// 	}

// 	if !reflect.DeepEqual(addresses, expectedAddresses) {
// 		t.Errorf("Mismatched Addresses. Expected: %v, but got: %v", addresses, expectedAddresses)
// 	}

// 	// Ensure all expectations were met
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		t.Errorf("Unfulfilled expectations: %v", err)
// 	}

// }
