// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProjectColumns holds the columns for the "project" table.
	ProjectColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "author", Type: field.TypeString},
		{Name: "author_email", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "star", Type: field.TypeInt},
		{Name: "topic_id", Type: field.TypeInt},
	}
	// ProjectTable holds the schema information for the "project" table.
	ProjectTable = &schema.Table{
		Name:       "project",
		Columns:    ProjectColumns,
		PrimaryKey: []*schema.Column{ProjectColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_topic_projects",
				Columns:    []*schema.Column{ProjectColumns[7]},
				RefColumns: []*schema.Column{TopicColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TopicColumns holds the columns for the "topic" table.
	TopicColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "topic_category_id", Type: field.TypeInt},
	}
	// TopicTable holds the schema information for the "topic" table.
	TopicTable = &schema.Table{
		Name:       "topic",
		Columns:    TopicColumns,
		PrimaryKey: []*schema.Column{TopicColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "topic_topic_category_topics",
				Columns:    []*schema.Column{TopicColumns[3]},
				RefColumns: []*schema.Column{TopicCategoryColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TopicCategoryColumns holds the columns for the "topic_category" table.
	TopicCategoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "icon", Type: field.TypeString},
	}
	// TopicCategoryTable holds the schema information for the "topic_category" table.
	TopicCategoryTable = &schema.Table{
		Name:       "topic_category",
		Columns:    TopicCategoryColumns,
		PrimaryKey: []*schema.Column{TopicCategoryColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProjectTable,
		TopicTable,
		TopicCategoryTable,
	}
)

func init() {
	ProjectTable.ForeignKeys[0].RefTable = TopicTable
	ProjectTable.Annotation = &entsql.Annotation{
		Table: "project",
	}
	TopicTable.ForeignKeys[0].RefTable = TopicCategoryTable
	TopicTable.Annotation = &entsql.Annotation{
		Table: "topic",
	}
	TopicCategoryTable.Annotation = &entsql.Annotation{
		Table: "topic_category",
	}
}
