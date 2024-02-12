package svc

import (
	"testing"

	"receipt/internal/model"
)

func TestCalcPoints(t *testing.T) {
	type args struct {
		test model.Receipt
		// expected value
		want int64
	}
	tests := []args{
		{
			// integer
			test: model.Receipt{
				PurchaseDate: "2024-02-02",
				PurchaseTime: "11:20",
				Total:        "1.00",
			},
			want: 75,
		},
		{
			// multiple of 0.25
			test: model.Receipt{
				PurchaseDate: "2024-02-02",
				PurchaseTime: "11:20",
				Total:        "1.25",
			},
			want: 25,
		},
		{
			// odd date
			test: model.Receipt{
				PurchaseDate: "2024-02-01",
				PurchaseTime: "11:20",
			},
			want: 6,
		},
		{
			// 2-4PM
			test: model.Receipt{
				PurchaseDate: "2024-02-02",
				PurchaseTime: "15:20",
			},
			want: 10,
		},
		{
			// short description is multiple of 3
			test: model.Receipt{
				PurchaseDate: "2024-02-02",
				PurchaseTime: "20:20",
				Items: []model.ReceiptItem{
					{ShortDescription: "testtesttest", Price: "100.2"},
					{ShortDescription: "testtesttest", Price: "105.3"},
				},
			},
			// 0.2 * 100.2 + 0.2 * 105.3 + 5
			want: 21 + 22 + 5,
		},
		{
			// alphanumeric value in retailer name
			test: model.Receipt{
				Retailer:     "a6s=56q",
				PurchaseDate: "2024-02-02",
				PurchaseTime: "20:20",
			},
			want: 6,
		},
		{
			// alphanumeric value in retailer name, and points for every two items
			test: model.Receipt{
				Retailer:     "a6s=56q",
				PurchaseDate: "2024-02-02",
				PurchaseTime: "20:20",
				Items: []model.ReceiptItem{
					{},
					{},
					{},
					{},
					{},
				},
			},
			want: 6 + 5*2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test.Retailer, func(t *testing.T) {
			got, err := CalcPoints(tt.test)
			if err != nil {
				t.Errorf("CalcPoints() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("CalcPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
