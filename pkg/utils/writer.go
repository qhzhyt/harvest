package utils

type TempWriter []byte

func (w *TempWriter) Write(data []byte) (n int, err error) {
    *w = append(*w, data...)
    return len(data), nil
}

func (w *TempWriter) String() string {
    return string(*w)
}

func (w *TempWriter) Bytes() []byte {
    return *w
}
