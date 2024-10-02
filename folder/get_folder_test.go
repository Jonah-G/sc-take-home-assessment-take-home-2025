package folder_test

import (
	"testing"
	"errors"
	"github.com/Jonah-G/sc-take-home-assessment-take-home-2025/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	orgID1 := uuid.Must(uuid.NewV4())
	orgID2 := uuid.Must(uuid.NewV4())

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name: "Valid orgId with folders",
			orgID: orgID1,
			folders: []folder.Folder{
				{Name: "folder1", OrgId: orgID1, Paths: "folder1"},
				{Name: "folder2", OrgId: orgID1, Paths: "folder2"},
				{Name: "folder3", OrgId: orgID2, Paths: "folder3"},
			},
			want: []folder.Folder{
				{Name: "folder1", OrgId: orgID1, Paths: "folder1"},
				{Name: "folder2", OrgId: orgID1, Paths: "folder2"},
			},
		},
		{
			name: "Valid orgId with sub folders",
			orgID: orgID1,
			folders: []folder.Folder{
				{Name: "folder1", OrgId: orgID1, Paths: "folder1"},
				{Name: "folder2", OrgId: orgID1, Paths: "folder2"},
				{Name: "folder11", OrgId: orgID1, Paths: "folder1.folder11"},
				{Name: "folder3", OrgId: orgID2, Paths: "folder3"},
			},
			want: []folder.Folder{
				{Name: "folder1", OrgId: orgID1, Paths: "folder1"},
				{Name: "folder2", OrgId: orgID1, Paths: "folder2"},
				{Name: "folder11", OrgId: orgID1, Paths: "folder1.folder11"},
			},
		},
		{
			name: "No folders with orgId",
			orgID: orgID2,
			folders: []folder.Folder{
				{Name: "folder1", OrgId: orgID1, Paths: "folder1"},
				{Name: "folder2", OrgId: orgID1, Paths: "folder2"},
			},
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)

			assert.Equal(t, tt.want, get)

		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	orgID1 := uuid.Must(uuid.NewV4())
	orgID2 := uuid.Must(uuid.NewV4())

	tests := [...]struct {
		name string
		folderName string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
		wantErr error
	}{
		{
			name: "Valid orgId with folders",
			folderName:  "folder1",
			orgID: orgID1,
			folders: []folder.Folder{
				{Name: "folder1", OrgId: orgID1, Paths: "folder1"},
				{Name: "folder11", OrgId: orgID1, Paths: "folder1.folder11"},
				{Name: "folder12", OrgId: orgID1, Paths: "folder1.folder12"},
				{Name: "folder2", OrgId: orgID2, Paths: "folder2"},
				{Name: "folder21", OrgId: orgID2, Paths: "folder2.21"},
			},
			want: []folder.Folder{
				{Name: "folder11", OrgId: orgID1, Paths: "folder1.folder11"},
				{Name: "folder12", OrgId: orgID1, Paths: "folder1.folder12"},
			},
			wantErr: nil,
		},
		{
			name:  "Valid orgID with deeper folders",
			folderName:  "folder1",
			orgID: orgID1,
			folders: []folder.Folder{
				{Name: "folder1", OrgId: orgID1, Paths: "folder1"},
				{Name: "folder11", OrgId: orgID1, Paths: "folder1.folder11"},
				{Name: "folder12", OrgId: orgID1, Paths: "folder1.folder12"},
				{Name: "folder111", OrgId: orgID2, Paths: "folder1.folder11.folder111"},
				{Name: "folder2", OrgId: orgID2, Paths: "folder2"},
				{Name: "folder21", OrgId: orgID2, Paths: "folder2.21"},
			},
			want: []folder.Folder{
				{Name: "folder11", OrgId: orgID1, Paths: "folder1.folder11"},
				{Name: "folder12", OrgId: orgID1, Paths: "folder1.folder12"},
				{Name: "folder111", OrgId: orgID2, Paths: "folder1.folder11.folder111"},
			},
			wantErr: nil,
		},
		{
			name:  "Folder exists but no children exist",
			folderName:  "folder1",
			orgID: orgID1,
			folders: []folder.Folder{
				{Name: "folder1", OrgId: orgID1, Paths: "folder1"},
			},
			want: []folder.Folder{},
			wantErr: nil,
		},
		{
			name:  "Folder exists deeper but no children exist",
			folderName:  "folder11",
			orgID: orgID1,
			folders: []folder.Folder{
				{Name: "folder11", OrgId: orgID1, Paths: "folder1.folder11"},
			},
			want: []folder.Folder{},
			wantErr: nil,
		},
		{
			name:  "Folder does not exist",
			folderName:  "invalid_folder",
			orgID: orgID1,
			folders: []folder.Folder{
				{Name: "folder11", OrgId: orgID1, Paths: "folder1.folder11"},
			},
			want: []folder.Folder{},
			wantErr: errors.New("Error: Folder does not exist"),
		},
		{
			name:  "Folder does not exist in the specified organisation",
			folderName:  "folder21",
			orgID: orgID1,
			folders: []folder.Folder{
				{Name: "folder1", OrgId: orgID1, Paths: "folder1"},
				{Name: "folder11", OrgId: orgID1, Paths: "folder1.folder11"},
				{Name: "folder2", OrgId: orgID2, Paths: "folder2"},
				{Name: "folder21", OrgId: orgID2, Paths: "folder2.folder21"},
			},
			want: []folder.Folder{},
			wantErr: errors.New("Error: Folder does not exist in the specified organization"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(tt.orgID, tt.folderName)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.want, get)
			}
		})
	}
}