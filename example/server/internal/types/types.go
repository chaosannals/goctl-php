// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.7

package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}
