package models

import (
	"time"

	"github.com/sev-2/raiden"
)

type Books struct {
    raiden.ModelBase

    Id        int64      `json:"id,omitempty" column:"name:id;type:bigint;primaryKey;autoIncrement;nullable:false"`
    Title     *string    `json:"title,omitempty" column:"name:title;type:varchar;nullable;default:'255'::character varying"`
    Body      *string    `json:"body,omitempty" column:"name:body;type:text;nullable"`
    CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`

    // Table information
    Metadata string `json:"-" schema:"public"`

    // Access control
    Acl string `json:"-" read:"" write:""`
}