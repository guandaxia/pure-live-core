// Package ws_user_info comment
// This file was generated by tars2go 1.1.4
// Generated from ws_user_info.tars
package ws_user_info

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// WSUserInfo struct implement
type WSUserInfo struct {
	LUid       int64  `json:"lUid"`
	BAnonymous bool   `json:"bAnonymous"`
	SGuid      string `json:"sGuid"`
	SToken     string `json:"sToken"`
	LTid       int64  `json:"lTid"`
	LSid       int64  `json:"lSid"`
	LGroupId   int64  `json:"lGroupId"`
	LGroupType int64  `json:"lGroupType"`
}

func (st *WSUserInfo) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *WSUserInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int64(&st.LUid, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.BAnonymous, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SGuid, 2, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SToken, 3, false)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.LTid, 4, false)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.LSid, 5, false)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.LGroupId, 6, false)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.LGroupType, 7, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *WSUserInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require WSUserInfo, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *WSUserInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int64(st.LUid, 0)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.BAnonymous, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SGuid, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SToken, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.LTid, 4)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.LSid, 5)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.LGroupId, 6)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.LGroupType, 7)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *WSUserInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
