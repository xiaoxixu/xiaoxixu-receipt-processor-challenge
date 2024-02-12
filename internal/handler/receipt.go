package handler

import (
	"receipt/app/api/receipt"
	"receipt/internal/common"
	"receipt/internal/svc"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func ReceiptsProcess(c *gin.Context) {
	var req receipt.ReceiptProcessReq

	if err := c.ShouldBind(&req); err != nil {
		common.JsonError(c, "The receipt is invalid. It cannot be parsed.")
		return
	}
 
    if _, err := time.Parse("2006-01-02 15:04", req.PurchaseDate + " " + req.PurchaseTime); err != nil {
		common.JsonError(c, "The receipt is invalid. Purchase date cannot be parsed.")
		return
	}

	if len(req.Items) == 0 {
		common.JsonError(c, "The receipt is invalid. Receipt does not contain any items.")
		return
	}

	if !receiptsProcessValidate(`^\d+\.\d{2}$`, req.Total) {
		common.JsonError(c, "The receipt is invalid. Receipt total cannot be parsed.")
		return
	}

	for _, v := range req.Items {
		if !receiptsProcessValidate(`^\d+\.\d{2}$`, v.Price) {
			common.JsonError(c, "The receipt is invalid. Item price cannot be parsed.")
			return
		}
	}

	id, err := svc.ReceiptsProcess(&req)
	if err != nil {
		common.JsonError(c, "The receipt cannot be processed.")
		return
	}

	common.JsonOk(c, receipt.ReceiptProcessResp{ID: id})
}

func receiptsProcessValidate(reg, val string) bool {
	re := regexp.MustCompile(reg)
	return re.MatchString(val)
}

func ReceiptsPoints(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		common.JsonNotFoundError(c, "Input id is empty.")
		return
	}

	points, err := svc.ReceiptsPoints(id)
	if err != nil {
		common.JsonNotFoundError(c, "No receipt found for that id.")
		return
	}

	common.JsonOk(c, receipt.ReceiptPointResp{Points: points})
}
