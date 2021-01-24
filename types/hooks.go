package types

type THookFunc func(response IResponse, db IGormDB, rw IReadWriter) (resultDD IGormDB, err error)

type IBeforeDBFind interface {
	BeforeDBFind(response IResponse, db IGormDB, rw IReadWriter) (resultDD IGormDB, err error)
}

type IAfterDBFind interface {
	AfterDBFind(response IResponse, db IGormDB, rw IReadWriter) (resultDD IGormDB, err error)
}

type IBeforeDBDelete interface {
	BeforeDBDelete(response IResponse, db IGormDB, rw IReadWriter) (resultDD IGormDB, err error)
}

type IAfterDBDelete interface {
	AfterDBDelete(response IResponse, db IGormDB, rw IReadWriter) (resultDD IGormDB, err error)
}

type IBeforeDBCreate interface {
	BeforeDBCreate(response IResponse, db IGormDB, rw IReadWriter) (resultDD IGormDB, err error)
}

type IAfterDBCreate interface {
	AfterDBCreate(response IResponse, db IGormDB, rw IReadWriter) (resultDD IGormDB, err error)
}

type IBeforeDBChange interface {
	BeforeDBChange(response IResponse, db IGormDB, rw IReadWriter) (resultDD IGormDB, err error)
}

type IAfterDBChange interface {
	AfterDBChange(response IResponse, db IGormDB, rw IReadWriter) (resultDD IGormDB, err error)
}
