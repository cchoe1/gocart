package gocart

import (
    "encoding/json"
    "time"
)

type Response struct {

    Message string

    Timestamp int64

    Success bool

    Code int16

    Payload []string

}

type CliResponse struct {

    Response Response
}

type WebResponse struct {

    Response Response
}

/**
 * Returns a json string of the Response struct
 */
func (wr WebResponse) encoded() string {
    bytes, err := json.Marshal(wr.Response)
    if err != nil {
        panic(err)
    }
    return string(bytes)
}

/**
 * Returns a json string of the Response struct
 */
func (cr CliResponse) encoded() string {
    bytes, err := json.Marshal(cr.Response)
    if err != nil {
        panic(err)
    }
    return string(bytes)
}

func (wr *WebResponse) setTime() *WebResponse {
    wr.Response.Timestamp = time.Now().Unix()
    return wr
}

func (cr *CliResponse) setTime() *CliResponse {
    cr.Response.Timestamp = time.Now().Unix()
    return cr
}
