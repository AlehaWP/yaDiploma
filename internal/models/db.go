package models

import "context"

type ServerDB interface {
	CheckDBConnectiond(context.Context)
}
