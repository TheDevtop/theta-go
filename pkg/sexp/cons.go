package sexp

type Cons struct {
	Car Value
	Cdr List
}

func (list List) Cons() Cons {
	if len(list) < 1 {
		return Cons{
			Car: nil,
			Cdr: nil,
		}
	}
	if len(list) == 1 {
		return Cons{
			Car: list[0],
			Cdr: nil,
		}
	}
	return Cons{
		Car: list[0],
		Cdr: list[1:],
	}
}

func (cons Cons) List() List {
	return append(List{cons.Car}, cons.Cdr...)
}
