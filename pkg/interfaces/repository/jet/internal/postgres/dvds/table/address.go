//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Address = newAddressTable("dvds", "address", "")

type addressTable struct {
	postgres.Table

	//Columns
	AddressID  postgres.ColumnInteger
	Address    postgres.ColumnString
	Address2   postgres.ColumnString
	District   postgres.ColumnString
	CityID     postgres.ColumnInteger
	PostalCode postgres.ColumnString
	Phone      postgres.ColumnString
	LastUpdate postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AddressTable struct {
	addressTable

	EXCLUDED addressTable
}

// AS creates new AddressTable with assigned alias
func (a AddressTable) AS(alias string) *AddressTable {
	return newAddressTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AddressTable with assigned schema name
func (a AddressTable) FromSchema(schemaName string) *AddressTable {
	return newAddressTable(schemaName, a.TableName(), a.Alias())
}

func newAddressTable(schemaName, tableName, alias string) *AddressTable {
	return &AddressTable{
		addressTable: newAddressTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newAddressTableImpl("", "excluded", ""),
	}
}

func newAddressTableImpl(schemaName, tableName, alias string) addressTable {
	var (
		AddressIDColumn  = postgres.IntegerColumn("address_id")
		AddressColumn    = postgres.StringColumn("address")
		Address2Column   = postgres.StringColumn("address2")
		DistrictColumn   = postgres.StringColumn("district")
		CityIDColumn     = postgres.IntegerColumn("city_id")
		PostalCodeColumn = postgres.StringColumn("postal_code")
		PhoneColumn      = postgres.StringColumn("phone")
		LastUpdateColumn = postgres.TimestampColumn("last_update")
		allColumns       = postgres.ColumnList{AddressIDColumn, AddressColumn, Address2Column, DistrictColumn, CityIDColumn, PostalCodeColumn, PhoneColumn, LastUpdateColumn}
		mutableColumns   = postgres.ColumnList{AddressColumn, Address2Column, DistrictColumn, CityIDColumn, PostalCodeColumn, PhoneColumn, LastUpdateColumn}
	)

	return addressTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		AddressID:  AddressIDColumn,
		Address:    AddressColumn,
		Address2:   Address2Column,
		District:   DistrictColumn,
		CityID:     CityIDColumn,
		PostalCode: PostalCodeColumn,
		Phone:      PhoneColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}