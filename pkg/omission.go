package middle

import "io"

func (ctx *Context) ReadBody() ([]byte, error) {
	body, err := io.ReadAll(ctx.Body())
	if err != nil {
		return nil, err
	}
	return body, nil
}
