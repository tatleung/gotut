// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewCreateUserDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewCreateUserDefault creates a CreateUserDefault with default headers values
func NewCreateUserDefault(code int) *CreateUserDefault {
	return &CreateUserDefault{
		_statusCode: code,
	}
}

/*CreateUserDefault handles this case with default header values.

successful operation
*/
type CreateUserDefault struct {
	_statusCode int
}

// Code gets the status code for the create user default response
func (o *CreateUserDefault) Code() int {
	return o._statusCode
}

func (o *CreateUserDefault) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUser default ", o._statusCode)
}

func (o *CreateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
