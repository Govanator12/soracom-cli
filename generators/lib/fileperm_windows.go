// +build windows

package lib

import (
	"fmt"

	"github.com/bearmini/go-acl/api"
	"golang.org/x/sys/windows"
)

func IsFilePermissionTooOpen(path string) (bool, error) {
	var (
		ownerSID *windows.SID
		dacl     *api.ACL
		secDesc  windows.Handle
	)
	err := api.GetNamedSecurityInfo(
		path,
		api.SE_FILE_OBJECT,
		api.OWNER_SECURITY_INFORMATION|api.DACL_SECURITY_INFORMATION,
		&ownerSID,
		nil,
		&dacl,
		nil,
		&secDesc,
	)
	defer windows.LocalFree(secDesc)
	if err != nil {
		// This `err` always contains "The operation completed successfully"
		// So we create a new error instance
		return false, fmt.Errorf("unable to get security info for the file: %s", path)
	}

	currProcSID, err := GetCurrentProcessSID()
	if err != nil {
		return false, err
	}
	//fmt.Println(sidToString(currProcSID))

	//fmt.Printf("dacl == %+v\n", dacl)
	aces := dacl.GetACEList()
	//fmt.Printf("ACEs == %+v\n", aces)
	for _, ace := range aces {
		switch ace.(type) {
		case *api.AccessAllowedACE:
		// ok to have this if it's sid == mine
		default:
			return true, nil
		}
		//fmt.Println(sidToString(ace.GetSID()))
		if !windows.EqualSid(ace.GetSID(), currProcSID) {
			return true, nil
		}
	}
	return false, nil
}

func GetCurrentProcessSID() (*windows.SID, error) {
	token, err := windows.OpenCurrentProcessToken()
	if err != nil {
		return nil, err
	}
	defer token.Close()

	tu, err := token.GetTokenUser()
	if err != nil {
		return nil, err
	}
	return tu.User.Sid, nil
}

func sidToString(sid *windows.SID) string {
	str, err := sid.String()
	if err != nil {
		return "<err: " + err.Error()
	}
	return str
}