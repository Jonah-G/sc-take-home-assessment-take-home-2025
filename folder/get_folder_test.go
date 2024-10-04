package folder_test

import (
	"testing"
	"github.com/Jonah-G/sc-take-home-assessment-take-home-2025/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	orgID1 := uuid.Must(uuid.NewV4())
	orgID2 := uuid.Must(uuid.NewV4())
	orgID3 := uuid.Must(uuid.NewV4())

	folders := []folder.Folder{
		{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
		{Name: "bravo", OrgId: orgID1, Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: orgID1, Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
		{Name: "echo", OrgId: orgID1, Paths: "echo"},
		{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
	}

	f := folder.NewDriver(folders)

	t.Run("Valid call", func(t *testing.T) {
		correctOutput := []folder.Folder{
			{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
			{Name: "bravo", OrgId: orgID1, Paths: "alpha.bravo"},
			{Name: "charlie", OrgId: orgID1, Paths: "alpha.bravo.charlie"},
			{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
			{Name: "echo", OrgId: orgID1, Paths: "echo"},
		}
		result := f.GetFoldersByOrgID(orgID1)
		assert.Equal(t, correctOutput, result)
	})

	t.Run("No folders", func(t *testing.T) {
		correctOutput := []folder.Folder{}
		result := f.GetFoldersByOrgID(orgID3)
		assert.Equal(t, correctOutput, result)
	})
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	orgID1 := uuid.Must(uuid.NewV4())
	orgID2 := uuid.Must(uuid.NewV4())

	folders := []folder.Folder{
		{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
		{Name: "bravo", OrgId: orgID1, Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: orgID1, Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
		{Name: "echo", OrgId: orgID1, Paths: "echo"},
		{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
	}

	f := folder.NewDriver(folders)

	t.Run("Valid call", func(t *testing.T) {
		correctOutput := []folder.Folder{
			{Name: "bravo", OrgId: orgID1, Paths: "alpha.bravo"},
			{Name: "charlie", OrgId: orgID1, Paths: "alpha.bravo.charlie"},
			{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
		}
		result, err := f.GetAllChildFolders(orgID1, "alpha")
		assert.NoError(t, err)
		assert.Equal(t, correctOutput, result)
	})

	t.Run("Valid call 2", func(t *testing.T) {
		correctOutput := []folder.Folder{
			{Name: "charlie", OrgId: orgID1, Paths: "alpha.bravo.charlie"},
		}
		result, err := f.GetAllChildFolders(orgID1, "bravo")
		assert.NoError(t, err)
		assert.Equal(t, correctOutput, result)
	})

	t.Run("No children 1", func(t *testing.T) {
		correctOutput := []folder.Folder{}
		result, err := f.GetAllChildFolders(orgID1, "echo")
		assert.NoError(t, err)
		assert.ElementsMatch(t, correctOutput, result)
	})

	t.Run("No children 2", func(t *testing.T) {
		correctOutput := []folder.Folder{}
		result, err := f.GetAllChildFolders(orgID1, "charlie")
		assert.NoError(t, err)
		assert.ElementsMatch(t, correctOutput, result)
	})

	t.Run("Folder does not exist", func(t *testing.T) {
		_, err := f.GetAllChildFolders(orgID1, "invalid_folder")
		assert.EqualError(t, err, "Error: Folder does not exist")
	})

	t.Run("Invalid Organisation", func(t *testing.T) {
		_, err := f.GetAllChildFolders(orgID1, "foxtrot")
		assert.EqualError(t, err, "Error: Folder does not exist in the specified organization")
	})
}