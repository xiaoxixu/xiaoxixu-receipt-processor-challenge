package svc

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"receipt/app/api/receipt"
	"receipt/internal/model"

	uuid "github.com/satori/go.uuid"
)

var db = make(map[string]model.Receipt, 0)

func ReceiptsProcess(r *receipt.ReceiptProcessReq) (string, error) {
	id := uuid.NewV4().String()
	length := len(r.Items)
	items := make([]model.ReceiptItem, length)
	for i := 0; i < length; i++ {
		items[i] = model.ReceiptItem{
			ShortDescription: r.Items[i].ShortDescription,
			Price:            r.Items[i].Price,
		}
	}
	db[id] = model.Receipt{
		ID:           id,
		Retailer:     r.Retailer,
		PurchaseDate: r.PurchaseDate,
		PurchaseTime: r.PurchaseTime,
		Items:        items,
		Total:        r.Total,
	}
	return id, nil
}

func CalcPoints(v model.Receipt) (int64, error) {
	var points int64 = 0

	// total
	ss := strings.Split(v.Total, ".")
	if len(ss) == 2 {
		if ss[1] == "00" {
			// 50 + 25 for positive integer
			points += 75
		}
		if ss[1] == "25" || ss[1] == "50" || ss[1] == "75" {
			// 25 for multiple of 0.25
			points += 25
		}
	}

	date, err := time.Parse("2006-01-02 15:04", v.PurchaseDate+" "+v.PurchaseTime)
	if err != nil {
		return 0, err
	}
	// 6 for odd date
	if date.Day()%2 == 1 {
		points += 6
	}

	// 10 for 2-4PM
	h := date.Hour()
	if h >= 14 && h < 16 {
		points += 10
	}

	// short description is multiple of 3
	var dp int64
	for _, i := range v.Items {
		if len(strings.TrimSpace(i.ShortDescription))%3 == 0 {
			f, _ := strconv.ParseFloat(i.Price, 64)
			if f != 0 {
				dp += int64(math.Ceil(f * 0.2))
			}
		}
	}
	points += dp
	// point for every alphanumeric value
	for _, r := range v.Retailer {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			points += 1
		}
	}

	// 5 points for every pair
	points += int64(len(v.Items) / 2) * 5

	return points, nil
}

func ReceiptsPoints(id string) (int64, error) {
	v, ok := db[id]
	if !ok {
		return 0, errors.New("not found")
	}

	return CalcPoints(v)
}
