package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) Uint8() (uint8, error) {
	v, err := strconv.Atoi(s.String())
	return uint8(v), err
}

func (s StrTo) MustUint8() uint8 {
	v, _ := s.Uint8()
	return v
}

func (s StrTo) Uint32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	if v < 1 {
		return 0, nil
	}
	return uint32(v), err
}

func (s StrTo) MustUint32() uint32 {
	v, _ := s.Uint32()
	return v
}
