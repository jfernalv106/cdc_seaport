package utils

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func DecodeBinaryDecimal(b64 *string, scale int32) float64 {
	raw, err := base64.StdEncoding.DecodeString(*b64)
	if err != nil {
		return 0
	}

	// two's complement → signed big.Int
	n := new(big.Int).SetBytes(raw)
	if len(raw) > 0 && (raw[0]&0x80) != 0 {
		// si es negativo, ajustar desde two's complement
		bits := uint(len(raw) * 8)
		max := new(big.Int).Lsh(big.NewInt(1), bits) // 2^(8*len)
		n.Sub(n, max)                                // n = n - 2^(8*len)
	}

	dec := decimal.NewFromBigInt(n, -scale)
	f, _ := dec.Float64()
	return f
}

func ToStr(v *int64) *string {
	if v == nil {
		return nil
	}
	s := fmt.Sprintf("%d", *v)
	return &s
}

func ToFormattedDateTime(v *int64) *string {
	if v == nil {
		return nil
	}
	t := time.UnixMilli(*v)
	s := t.Format("2006-01-02 15:04:05")
	return &s
}
func ToFormattedDate(days *int64) *string {
	if days == nil {
		return nil
	}
	loc, err := time.LoadLocation("America/Santiago")
	if err != nil {
		return nil
	}
	// 1970-01-01 + days
	t := time.Unix(0, 0).In(loc).Add(time.Duration(*days) * 24 * time.Hour)
	s := t.Format("2006-01-02") // p.ej. "2006-01-02"
	return &s
}
func ToInt(s *string) *int64 {
	if s == nil {
		return nil
	}
	if i, err := strconv.ParseInt(*s, 10, 64); err == nil {
		return &i
	}
	return nil
}

func MapOperation(op string) string {
	switch op {
	case "c":
		return "CREATE"
	case "u":
		return "UPDATE"
	case "d":
		return "DELETE"
	case "r":
		return "CREATE"
	default:
		return op
	}
}
func ToFloat(val interface{}) float64 {
	switch v := val.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return f
		}
	}
	return 0
}
func ParseToFloat(s *string) *float64 {
	if s == nil {
		return nil
	}
	val := strings.ReplaceAll(strings.TrimSpace(*s), ",", ".")
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil
	}
	return &f
}
func ParseFecha(fechaStr string) (time.Time, error) {
	// Layout debe coincidir con el formato exacto de la fecha
	layout := "2006-01-02 15:04:05"
	return time.Parse(layout, fechaStr)
}
func RemoveFirstBy[T any](s []T, pred func(T) bool) []T {
	dst := s[:0]
	for i := range s {
		if !pred(s[i]) {
			dst = append(dst, s[i])
		} else {

			var zero T
			s[i] = zero
		}
	}
	return dst
}
