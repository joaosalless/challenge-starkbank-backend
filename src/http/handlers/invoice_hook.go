package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"joaosalless/challenge-starkbank/src/dtos"
)

func (h InvoiceHandler) InvoiceHookProcess(ctx *gin.Context) {
	var input dtos.InvoiceHookProcessInput

	err := json.NewDecoder(ctx.Request.Body).Decode(&input)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = h.invoiceController.InvoiceHookProcess(ctx, input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, nil)
}
