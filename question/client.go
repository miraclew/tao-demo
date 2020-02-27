package question

import (
    "bytes"
    "context"
    "encoding/json"
    "net/http"
)

const URL = ""

type Client struct {
    t *http.Client
}


func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
    res := new(CreateResponse)
    err := c.do(ctx, "create", req, res)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
    res := new(DeleteResponse)
    err := c.do(ctx, "delete", req, res)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (c *Client) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
    res := new(UpdateResponse)
    err := c.do(ctx, "update", req, res)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (c *Client) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
    res := new(GetResponse)
    err := c.do(ctx, "get", req, res)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (c *Client) Query(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
    res := new(QueryResponse)
    err := c.do(ctx, "query", req, res)
    if err != nil {
        return nil, err
    }
    return res, nil
}


func (c *Client) do(ctx context.Context, path string, req interface{}, res interface{}) error {
    buf := new(bytes.Buffer)
    err := json.NewEncoder(buf).Encode(req)
    if err != nil {
        return err
    }
    resp, err := c.t.Post(URL+"/"+path, "application/json", buf)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    err = json.NewDecoder(resp.Body).Decode(res)
    if err != nil {
        return err
    }
    return nil
}
