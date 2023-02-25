// Package ws_verify_cookie_req comment
// This file was generated by tars2go 1.1.4
// Generated from ws_verify_cookie_req.tars
package ws_verify_cookie_req

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// WSVerifyCookieReq struct implement
type WSVerifyCookieReq struct {
	LUid    int64  `json:"lUid"`
	SUA     string `json:"sUA"`
	SCookie string `json:"sCookie"`
}

func (st *WSVerifyCookieReq) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *WSVerifyCookieReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_int64(&st.LUid, 0, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SUA, 1, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.SCookie, 2, false)
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
func (st *WSVerifyCookieReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require WSVerifyCookieReq, but not exist. tag %d", tag)
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
func (st *WSVerifyCookieReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int64(st.LUid, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SUA, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.SCookie, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *WSVerifyCookieReq) WriteBlock(_os *codec.Buffer, tag byte) error {
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
