package helperfunc

import (
	"errors"
	"fmt"
	"strconv"
)

//UIDStrX : cheack the uid of type string and return the value in the format 0xXXXX and err
func UIDStrX(uid string) (string, error) {
	if len(uid) < 1 {
		return "", errors.New("invalid Uid")
	}
	bid := []byte(uid)
	if bid[0] == '0' && bid[1] == 'x' {
		_, err := strconv.ParseUint(string(bid[2:]), 16, 64)
		if err != nil {
			return "", errors.New("invalid hexnumber")
		}
		return uid, nil
	} else {
		_, err := strconv.ParseUint(uid, 16, 64)
		if err != nil {
			return "", errors.New("invalid hexnumber")
		}
		return fmt.Sprintf("0x%s", uid), nil
	}
}

//UIDStrXNR : cheack the uid of type string and return the value without any error
func UIDStrXNR(uid string) string {
	if len(uid) < 1 {
		return ""
	}
	bid := []byte(uid)
	if bid[0] == '0' && bid[1] == 'x' {
		_, err := strconv.ParseUint(string(bid[2:]), 16, 64)
		if err != nil {
			return ""
		}
		return uid
	} else {
		_, err := strconv.ParseUint(uid, 16, 64)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("0x%s", uid)
	}
}

//UID : cheack the uid of type string and return the uint64 value and err
func UID(uid string) (uint64, error) {
	if len(uid) < 1 {
		return 0, errors.New("invalid Uid")
	}
	bid := []byte(uid)
	if bid[0] == '0' && bid[1] == 'x' {
		iuid, err := strconv.ParseUint(string(bid[2:]), 16, 64)
		if err != nil {
			return 0, errors.New("invalid hexnumber")
		}
		return iuid, nil
	} else {
		iuid, err := strconv.ParseUint(uid, 16, 64)
		if err != nil {
			return 0, errors.New("invalid hexnumber")
		}
		return iuid, nil
	}
}

//UIDNR : cheack the uid of type string and return the uint64 value without any err , 0 on defaults or error
func UIDNR(uid string) uint64 {
	if len(uid) < 1 {
		return 0
	}
	bid := []byte(uid)
	if bid[0] == '0' && bid[1] == 'x' {
		iuid, err := strconv.ParseUint(string(bid[2:]), 16, 64)
		if err != nil {
			return 0
		}
		return iuid
	} else {
		iuid, err := strconv.ParseUint(uid, 16, 64)
		if err != nil {
			return 0
		}
		return iuid
	}
}
