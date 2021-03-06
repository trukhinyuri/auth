// Code generated by noice. DO NOT EDIT.
package autherr

import (
	bytes "bytes"
	cherry "github.com/containerum/cherry"
	template "text/template"
)

const ()

func ErrInvalidToken(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "invalid token received", StatusHTTP: 400, ID: cherry.ErrID{SID: "auth", Kind: 0x1}, Details: []string(nil), Fields: cherry.Fields(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrTokenNotOwnedBySender(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Can`t identify sender as token owner", StatusHTTP: 403, ID: cherry.ErrID{SID: "auth", Kind: 0x2}, Details: []string(nil), Fields: cherry.Fields(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrTokenNotFound(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Token was not found in storage", StatusHTTP: 404, ID: cherry.ErrID{SID: "auth", Kind: 0x3}, Details: []string(nil), Fields: cherry.Fields(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrInternal(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Token was not found in storage", StatusHTTP: 500, ID: cherry.ErrID{SID: "auth", Kind: 0x4}, Details: []string(nil), Fields: cherry.Fields(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrValidation(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Validation error", StatusHTTP: 400, ID: cherry.ErrID{SID: "auth", Kind: 0x5}, Details: []string(nil), Fields: cherry.Fields(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}
func renderTemplate(templText string) string {
	buf := &bytes.Buffer{}
	templ, err := template.New("").Parse(templText)
	if err != nil {
		return err.Error()
	}
	err = templ.Execute(buf, map[string]string{})
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
