package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_"`
	Invoice_id       string             `json:"invoice_id"`
	Order_id         string             `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CARD|re=CASH|eq="`
	Payment_status   *string            `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"Payment_due_date"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}
