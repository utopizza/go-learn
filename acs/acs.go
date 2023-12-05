package acs

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ACS struct {
	tickets []AuthnTicket
}

type AuthnTicket struct {
	Name   string
	Ticket string
}

type AuthnTicketOption func(ticket string) *ACS

func (s *ACS) WithNotLoginTicket(ctx *gin.Context, ticket string) *ACS {
	s.tickets = append(s.tickets, AuthnTicket{
		Name:   "not_login_ticket",
		Ticket: ticket,
	})
	return s
}

func (s *ACS) WithVerifyTicket(ctx *gin.Context, ticket string) *ACS {
	s.tickets = append(s.tickets, AuthnTicket{
		Name:   "verify_ticket",
		Ticket: ticket,
	})
	return s
}

func (s *ACS) CreateAuthnResult(ctx *gin.Context) *ACS {
	return s
}

func (s *ACS) Submit(ctx *gin.Context) error {
	fmt.Printf("%+v", s.tickets)
	return nil
}
