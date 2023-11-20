package office_repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/evgeny-tokarev/office_app/backend/util"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = util.InitTestDB()
	if err != nil {
		fmt.Println("Unable to create testing Database:", err.Error())
		return
	}

	os.Exit(m.Run())
}

func TestCreateOffice(t *testing.T) {
	arg, office, err := CreateRandomOffice(testDb)

	require.NoError(t, err)
	require.Equal(t, arg.Name, office.Name)
	require.Equal(t, arg.Address, office.Address)
	require.NotZero(t, office.ID)
	require.NotZero(t, office.CreatedAt)
	require.NotZero(t, office.UpdatedAt)
}

func TestGetOffice(t *testing.T) {
	_, office1, err1 := CreateRandomOffice(testDb)
	require.NoError(t, err1)
	office2, err2 := TestOfficeQueries.GetOffice(context.Background(), office1.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, office2)

	require.Equal(t, office1.ID, office2.ID)
	require.Equal(t, office1.Name, office2.Name)
	require.WithinDuration(t, office1.CreatedAt, office2.CreatedAt, time.Second)
	require.WithinDuration(t, office1.UpdatedAt, office2.UpdatedAt, time.Second)
}

func TestDeleteOffice(t *testing.T) {
	_, office1, err1 := CreateRandomOffice(testDb)
	require.NoError(t, err1)
	err := TestOfficeQueries.DeleteOffice(context.Background(), office1.ID)
	require.NoError(t, err)
	office2, err := TestOfficeQueries.GetOffice(context.Background(), office1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, office2)
}

func TestListOffice(t *testing.T) {
	var offices1 []Office
	for i := 0; i < 5; i++ {
		_, office, err := CreateRandomOffice(testDb)
		require.NoError(t, err)
		offices1 = append(offices1, office)
	}
	offices2, err := TestOfficeQueries.ListOffices(context.Background())
	require.NoError(t, err)
	require.True(t, len(offices2) >= 5)
	for i, office1 := range offices1 {
		office2 := offices2[len(offices2)-5+i]
		require.NotEmpty(t, office1)
		require.Equal(t, office1.ID, office2.ID)
		require.Equal(t, office1.Name, office2.Name)
		require.WithinDuration(t, office1.CreatedAt, office2.CreatedAt, time.Second)
		require.WithinDuration(t, office1.UpdatedAt, office2.UpdatedAt, time.Second)
	}
}

func TestAttachePhoto(t *testing.T) {
	_, office1, err := CreateRandomOffice(testDb)
	require.NoError(t, err)
	testParams := AttachePhotoParams{
		ImgFile: sql.NullString{String: util.RandomString(10), Valid: true},
		ID:      office1.ID,
	}

	err = TestOfficeQueries.AttachePhoto(context.Background(), testParams)

	require.NoError(t, err)
	office2, err := TestOfficeQueries.GetOffice(context.Background(), office1.ID)
	require.NoError(t, err)
	require.Equal(t, testParams.ImgFile, office2.ImgFile)
	TestOfficeQueries.DeleteOffice(context.Background(), office2.ID)
}
