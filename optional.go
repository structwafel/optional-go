package optional

type Option[T any] struct {
	value  T
	isSome bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value:  value,
		isSome: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		isSome: false,
	}
}

// FromPtr creates an option where nil is considered None
// func FromPtr[T any](value T) Option[T] {
// 	if value == nil {
// 		return None[*T]()
// 	}
// 	return Option[*T]{
// 		value:  value,
// 		isSome: true,
// 	}
// }

// option methods all taken from Rust 🦀

func (o Option[T]) IsSome() bool {
	return o.isSome
}

func (o Option[T]) IsSomeAnd(f func(T) bool) bool {
	if o.isSome {
		return f(o.value)
	}
	return false
}

func (o Option[T]) IsNone() bool {
	return !o.isSome
}

func (o Option[T]) IsNoneOr(f func(T) bool) bool {
	if o.isSome {
		return f(o.value)
	}
	return true
}

func (o Option[T]) Expect(msg string) T {
	if o.isSome {
		return o.value
	}
	panic(msg)
}

func (o Option[T]) Unwrap() T {
	if o.isSome {
		return o.value
	}
	panic("called `Option.Unwrap()` on a `None` value")
}

func (o Option[T]) UnwrapOr(defaultT T) T {
	if o.isSome {
		return o.value
	}
	return defaultT
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.isSome {
		return o.value
	}
	return f()
}

// Returns the contained value, or Zero if the option is [None].
func (o Option[T]) UnwrapOrZero() T {
	if o.isSome {
		return o.value
	}
	var zero T
	return zero
}

func Map[T any, U any](o Option[T], f func(T) U) Option[U] {
	if o.isSome {
		return Some(f(o.value))
	}
	return None[U]()
}

func (o Option[T]) Inspect(f func(T)) {
	if o.isSome {
		f(o.value)
	}
}

func MapOr[T any, U any](o Option[T], defaultU U, f func(T) U) U {
	if o.isSome {
		return f(o.value)
	}
	return defaultU
}

func MapOrElse[T any, U any](o Option[T], defaultF func() U, f func(T) U) U {
	if o.isSome {
		return f(o.value)
	}
	return defaultF()
}

func (o Option[T]) OkOr(err error) (T, error) {
	if o.isSome {
		return o.value, nil
	}
	var zero T
	return zero, err
}

func (o Option[T]) OkOrElse(err func() error) (T, error) {
	if o.isSome {
		return o.value, nil
	}
	var zero T
	return zero, err()
}

func (o Option[T]) And(other Option[T]) Option[T] {
	if o.isSome {
		if other.IsSome() {
			return other
		}
	}
	return None[T]()
}

func (o Option[T]) AndThen(f func(T) Option[T]) Option[T] {
	if o.isSome {
		return f(o.value)
	}
	return None[T]()
}

func (o Option[T]) Filter(f func(T) bool) Option[T] {
	if o.IsNone() {
		return None[T]()
	}
	if f(o.value) {
		return o
	}
	return None[T]()
}

func (o Option[T]) Or(other T) T {
	if o.isSome {
		return o.value
	}
	return other
}
func (o Option[T]) OrElse(f func() T) T {
	if o.isSome {
		return o.value
	}
	return f()
}

func (o Option[T]) Xor(other Option[T]) Option[T] {
	if o.isSome {
		if other.IsNone() {
			return o
		}
	}
	if other.IsSome() {
		if o.IsNone() {
			return other
		}
	}
	return None[T]()
}

func (o *Option[T]) Insert(value T) T {
	o.value = value
	o.isSome = true
	return value
}

func (o *Option[T]) GetOrInsert(value T) T {
	if o.isSome {
		return o.value
	}
	o.value = value
	o.isSome = true
	return value
}

func (o *Option[T]) GetOrInsertZero() T {
	if o.isSome {
		return o.value
	}
	var zero T
	o.value = zero
	o.isSome = true
	return zero
}

func (o *Option[T]) GetOrInsertWith(f func() T) T {
	if o.isSome {
		return o.value
	}
	o.value = f()
	o.isSome = true
	return o.value
}

// Takes the value out of the option, leaving a [None] in its place.
func (o *Option[T]) Take() Option[T] {
	if o.isSome {
		old := o.value
		var zero T
		o.isSome = false
		o.value = zero
		return Some(old)
	}

	return None[T]()
}

func (o *Option[T]) TakeIf(f func(T) bool) Option[T] {
	if o.isSome {
		if f(o.value) {
			return o.Take()
		}
	}
	return None[T]()
}

// Replaces the actual value in the option by the value given in parameter,
// returning the old value if present, leaving a [Some] in its place.
func (o *Option[T]) Replace(value T) Option[T] {
	if o.isSome {
		old := o.value
		o.value = value
		return Some(old)

	}
	o.value = value
	o.isSome = true
	return None[T]()
}

// Flatten method to convert Option[Option[T]] to Option[T]
func Flatten[T any](o Option[Option[T]]) Option[T] {
	if o.isSome {
		return o.value
	}
	return None[T]()
}
