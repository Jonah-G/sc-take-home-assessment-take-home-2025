package folder_test

import (
	"testing"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/Jonah-G/sc-take-home-assessment-take-home-2025/folder"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	orgID1 := uuid.Must(uuid.NewV4())
	orgID2 := uuid.Must(uuid.NewV4())

	folders := []folder.Folder{
		{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
		{Name: "bravo", OrgId: orgID1, Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: orgID1, Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
		{Name: "echo", OrgId: orgID1, Paths: "alpha.delta.echo"},
		{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
		{Name: "golf", OrgId: orgID1, Paths: "golf"},
	}

	f := folder.NewDriver(folders)

	t.Run("Valid move", func(t *testing.T) {
		correctOutput := []folder.Folder{
			{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
			{Name: "bravo", OrgId: orgID1, Paths: "alpha.delta.bravo"},
			{Name: "charlie", OrgId: orgID1, Paths: "alpha.delta.bravo.charlie"},
			{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
			{Name: "echo", OrgId: orgID1, Paths: "alpha.delta.echo"},
			{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
			{Name: "golf", OrgId: orgID1, Paths: "golf"},
		}
		updatedFolders, err := f.MoveFolder("bravo", "delta")
		assert.NoError(t, err)
		assert.Equal(t, correctOutput, updatedFolders)
	})

	t.Run("Valid move 2", func(t *testing.T) {
		correctOutput := []folder.Folder{
			{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
			{Name: "bravo", OrgId: orgID1, Paths: "golf.bravo"},
			{Name: "charlie", OrgId: orgID1, Paths: "golf.bravo.charlie"},
			{Name: "delta", OrgId: orgID1, Paths: "alpha.delta"},
			{Name: "echo", OrgId: orgID1, Paths: "alpha.delta.echo"},
			{Name: "foxtrot", OrgId: orgID2, Paths: "foxtrot"},
			{Name: "golf", OrgId: orgID1, Paths: "golf"},
		}
		updatedFolders, err := f.MoveFolder("bravo", "golf")
		assert.NoError(t, err)
		assert.Equal(t, correctOutput, updatedFolders)
	})

	t.Run("Move to a child of itself", func(t *testing.T) {
		_, err := f.MoveFolder("alpha", "bravo")
		assert.EqualError(t, err, "Error: Cannot move folder to a child of itself")
	})

	t.Run("Move to self", func(t *testing.T) {
		_, err := f.MoveFolder("bravo", "bravo")
		assert.EqualError(t, err, "Error: Cannot move a folder to itself")
	})

	t.Run("Invalid organisation", func(t *testing.T) {
		_, err := f.MoveFolder("alpha", "foxtrot")
		assert.EqualError(t, err, "Error: Cannot move a folder to a different organization")
	})

	t.Run("Source folder does not exist", func(t *testing.T) {
		_, err := f.MoveFolder("invalid_folder", "delta")
		assert.EqualError(t, err, "Error: Source folder does not exist")
	})

	t.Run("Destination folder does not exist", func(t *testing.T) {
		_, err := f.MoveFolder("alpha", "invalid_folder")
		assert.EqualError(t, err, "Error: Destination folder does not exist")
	})
}