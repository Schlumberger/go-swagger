package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// PostEventHandlerFunc turns a function with the right signature into a post event handler
type PostEventHandlerFunc func(PostEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostEventHandlerFunc) Handle(params PostEventParams) middleware.Responder {
	return fn(params)
}

// PostEventHandler interface for that can handle valid post event params
type PostEventHandler interface {
	Handle(PostEventParams) middleware.Responder
}

// NewPostEvent creates a new http.Handler for the post event operation
func NewPostEvent(ctx *middleware.Context, handler PostEventHandler) *PostEvent {
	return &PostEvent{Context: ctx, Handler: handler}
}

/*
Create new event.
*/
type PostEvent struct {
	Context *middleware.Context
	Params  PostEventParams
	Handler PostEventHandler
}

func (o *PostEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
