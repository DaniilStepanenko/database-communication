//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/v2/postgres"
)

var NicerButSlowerFilmList = newNicerButSlowerFilmListTable("dvds", "nicer_but_slower_film_list", "")

type nicerButSlowerFilmListTable struct {
	postgres.Table

	//Columns
	Fid         postgres.ColumnInteger
	Title       postgres.ColumnString
	Description postgres.ColumnString
	Category    postgres.ColumnString
	Price       postgres.ColumnFloat
	Length      postgres.ColumnInteger
	Rating      postgres.ColumnString
	Actors      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type NicerButSlowerFilmListTable struct {
	nicerButSlowerFilmListTable

	EXCLUDED nicerButSlowerFilmListTable
}

// AS creates new NicerButSlowerFilmListTable with assigned alias
func (a NicerButSlowerFilmListTable) AS(alias string) *NicerButSlowerFilmListTable {
	return newNicerButSlowerFilmListTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new NicerButSlowerFilmListTable with assigned schema name
func (a NicerButSlowerFilmListTable) FromSchema(schemaName string) *NicerButSlowerFilmListTable {
	return newNicerButSlowerFilmListTable(schemaName, a.TableName(), a.Alias())
}

func newNicerButSlowerFilmListTable(schemaName, tableName, alias string) *NicerButSlowerFilmListTable {
	return &NicerButSlowerFilmListTable{
		nicerButSlowerFilmListTable: newNicerButSlowerFilmListTableImpl(schemaName, tableName, alias),
		EXCLUDED:                    newNicerButSlowerFilmListTableImpl("", "excluded", ""),
	}
}

func newNicerButSlowerFilmListTableImpl(schemaName, tableName, alias string) nicerButSlowerFilmListTable {
	var (
		FidColumn         = postgres.IntegerColumn("fid")
		TitleColumn       = postgres.StringColumn("title")
		DescriptionColumn = postgres.StringColumn("description")
		CategoryColumn    = postgres.StringColumn("category")
		PriceColumn       = postgres.FloatColumn("price")
		LengthColumn      = postgres.IntegerColumn("length")
		RatingColumn      = postgres.StringColumn("rating")
		ActorsColumn      = postgres.StringColumn("actors")
		allColumns        = postgres.ColumnList{FidColumn, TitleColumn, DescriptionColumn, CategoryColumn, PriceColumn, LengthColumn, RatingColumn, ActorsColumn}
		mutableColumns    = postgres.ColumnList{FidColumn, TitleColumn, DescriptionColumn, CategoryColumn, PriceColumn, LengthColumn, RatingColumn, ActorsColumn}
	)

	return nicerButSlowerFilmListTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Fid:         FidColumn,
		Title:       TitleColumn,
		Description: DescriptionColumn,
		Category:    CategoryColumn,
		Price:       PriceColumn,
		Length:      LengthColumn,
		Rating:      RatingColumn,
		Actors:      ActorsColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}