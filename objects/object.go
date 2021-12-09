package objects

type Object struct {
	ID    int    `json:"id"`
	Owner int    `json:"owner"`
	Kind  string `json:"kind"`
	Data  string `json:"data"`
}

type ObjectReceiver interface {
	SetObject(o *Object)
}

func (o Object) Insert() error

func (o Object) Refresh(d Any) error

func (o Object) Update(d Any) error

func (o Object) Delete() error

func (o Object) GetData(d Any) error
