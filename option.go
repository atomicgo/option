package option

type Option[T any] struct {
	Value   *T
	Default T
}

type StringOption struct {
	Option[string]
}

func String(v string) StringOption {
	return StringOption{Option[string]{Value: &v}}
}

type IntOption struct {
	Option[int]
}

func Int(v int) IntOption {
	return IntOption{Option[int]{Value: &v}}
}
