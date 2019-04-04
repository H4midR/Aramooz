package response

//type response
type Response struct {
	Code    int         `json:"Code,omitempry"`
	State   int         `json:"State,omitempry"`
	Message string      `json:"Message,omitempry"`
	Data    interface{} `json:"Data,omitempry"`
}

func (r *Response) HandleErr(err error) {
	if err != nil {
		r.State = -1
		r.Code = -1
		r.Message = err.Error()
		r.Data = err
	} else {
		r.State = 1
		r.Code = 1
		r.Message = "OK"

	}
}
