// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package proto

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson4a0f95aaDecodeGithubComSamyfodilWeb3authProto(in *jlexer.Lexer, out *Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "token":
			out.Token = string(in.String())
		case "error":
			out.Error = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4a0f95aaEncodeGithubComSamyfodilWeb3authProto(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Token != "" {
		const prefix string = ",\"token\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	if in.Error != "" {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4a0f95aaEncodeGithubComSamyfodilWeb3authProto(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4a0f95aaEncodeGithubComSamyfodilWeb3authProto(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4a0f95aaDecodeGithubComSamyfodilWeb3authProto(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4a0f95aaDecodeGithubComSamyfodilWeb3authProto(l, v)
}
func easyjson4a0f95aaDecodeGithubComSamyfodilWeb3authProto1(in *jlexer.Lexer, out *Message) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "address":
			out.Address = string(in.String())
		case "init":
			easyjson4a0f95aaDecode(in, &out.Init)
		case "signature":
			out.Signature = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4a0f95aaEncodeGithubComSamyfodilWeb3authProto1(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Address != "" {
		const prefix string = ",\"address\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Address))
	}
	if true {
		const prefix string = ",\"init\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson4a0f95aaEncode(out, in.Init)
	}
	if in.Signature != "" {
		const prefix string = ",\"signature\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Signature))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4a0f95aaEncodeGithubComSamyfodilWeb3authProto1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4a0f95aaEncodeGithubComSamyfodilWeb3authProto1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4a0f95aaDecodeGithubComSamyfodilWeb3authProto1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4a0f95aaDecodeGithubComSamyfodilWeb3authProto1(l, v)
}
func easyjson4a0f95aaDecode(in *jlexer.Lexer, out *struct {
	Challenge string `json:"challenge"`
	Signature string `json:"signature"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "challenge":
			out.Challenge = string(in.String())
		case "signature":
			out.Signature = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4a0f95aaEncode(out *jwriter.Writer, in struct {
	Challenge string `json:"challenge"`
	Signature string `json:"signature"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Challenge != "" {
		const prefix string = ",\"challenge\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Challenge))
	}
	if in.Signature != "" {
		const prefix string = ",\"signature\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Signature))
	}
	out.RawByte('}')
}
