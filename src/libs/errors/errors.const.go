package errors

import (
	"net/http"
	"test-project-hernan/src/libs/i18n"
)

var (
	//ErrInternalServer indicates an internal server error, this is used a default error when it's a general error
	ErrInternalServer = NewMyError(http.StatusInternalServerError, i18n.Message{MessageID: "ERRORS.INTERNAL_SERVER"})

	//ErrUserInvalid indicates an invalid user
	ErrGettingUser = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.GET_USER"})

	//ErrIDNotNumeric indicates a numeric ID wasn't send as numeric
	ErrIDNotNumeric = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.ID_NOT_NUMERIC"})

	//ErrUserIDDuplicated indicates the UserID is duplicated
	ErrUserIDDuplicated = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.USER_ID_DUPLICATED"})

	//ErrProductNotEnabled indicate the Product is not enabled
	ErrProductNotEnabled = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.PRODUCT_NOT_ENABLED"})

	//ErrStockMovementInvalid indicates the Stock Movement is invalid
	ErrStockMovementInvalid = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.STOCK_MOVEMENT_INVALID"})

	//ErrCantDeleteProduct indicates the product can't be deleted
	ErrCantDeleteProduct = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.CANT_DELETE_PRODUCT"})

	//ErrCantDeleteWarehouse indicates the warehouse can't be deleted
	ErrCantDeleteWarehouse = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.CANT_DELETE_WAREHOUSE"})

	//ErrConnectionProvider indicates comunication error with provider
	ErrConnectionProvider = NewMyError(http.StatusServiceUnavailable, i18n.Message{MessageID: "ERRORS.CONNECTION_PROVIDER"})

	//ErrPageSizeTooHigh indicates that pageSize param is higher than the valid number
	ErrPageSizeTooHigh = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.PAGE_SIZE_TOO_LARGE"})

	//ErrPageTooHigh indicates that page param is higher than the valid number
	ErrPageTooHigh = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.PAGE_TOO_LARGE"})

	//ErrURLNotFound indicates a requested URL doesn't exist
	ErrURLNotFound = NewMyError(http.StatusNotFound, i18n.Message{MessageID: "ERRORS.URL_NOT_FOUND"})

	//ErrNotFound indicates an entity not found error
	ErrNotFound = NewMyError(http.StatusNotFound, i18n.Message{MessageID: "ERRORS.NOT_FOUND"})
)

//Private errors
var (
	errUnsupportedFieldValue = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.UNSUPPORTED_FIELD_VALUE"})
	errFieldValidation       = NewMyError(http.StatusBadRequest, i18n.Message{MessageID: "ERRORS.FIELD_VALIDATION"})
)
