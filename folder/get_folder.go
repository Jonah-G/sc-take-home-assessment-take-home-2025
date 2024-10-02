package folder

import "github.com/gofrs/uuid"
import "strings"
import "errors"

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folders := f.folders
	var parent *Folder
	correctOrg := true
	for _, f := range folders {
		if strings.HasSuffix(f.Paths, name) {
			if f.OrgId == orgID {
				parent = &f
				break
			} else {
				correctOrg = false
			}
		}
	}

	if parent == nil {
		if !correctOrg {
			return nil, errors.New("Error: Folder does not exist in the specified organization")
		}
		return nil, errors.New("Error: Folder does not exist")
	}

	var children []Folder
	for _, f := range folders {
		if strings.Contains(f.Paths, parent.Paths+".") {
			children = append(children, f)
		}
	}
	return children, nil
}
