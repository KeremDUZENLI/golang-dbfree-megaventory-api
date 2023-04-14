package model

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func (l ProductList) ConvertStringGet() string {
	var str strings.Builder

	for _, v := range l.MvProducts {
		str.WriteString(fmt.Sprintf("-SKU: %v\n-Description: %v\n-Sales Price: %v\n-Purchase Price: %v\n\n",
			v.ProductID,
			v.ProductDescription,
			v.ProductSellingPrice,
			v.ProductPurchasePrice))
	}

	return str.String()
}

func (s SupplierClientList) ConvertStringPost() string {
	var str strings.Builder

	for _, l := range s.MvSupplierClient.SupplierClientAddresses {
		str.WriteString(fmt.Sprintf("-Name: %v\n-E-mail: %v\n-Shipping Address: %v, %v\n-Phone: %v\n\n",
			s.MvSupplierClient.SupplierClientName,
			s.MvSupplierClient.SupplierClientEmail,
			l.AddressLine1, l.City,
			s.MvSupplierClient.SupplierClientPhone1))
	}

	return str.String()
}

func (s InventoryList) ConvertStringInventory() string {
	var str strings.Builder

	for _, v := range s.MvProductStockList {
		for _, l := range v.MvStock {
			str.WriteString(fmt.Sprintf("-ProductID: %v\n-InventoryLocationID: %v\n-SubLocation: %v\n\n",
				v.ProductID,
				l.InventoryLocationID,
				l.SubLocation))
		}

	}

	return str.String()
}

// GIN ----------------------------------------------------------------
func (l ProductList) ConvertGinGet(id int) gin.H {
	var product gin.H

	for _, v := range l.MvProducts {
		if v.ProductID == id {
			product = gin.H{
				"sku":           v.ProductID,
				"description":   v.ProductDescription,
				"salesPrice":    v.ProductSellingPrice,
				"purchasePrice": v.ProductPurchasePrice,
			}
			break
		}
	}

	return gin.H{"products": product}
}
