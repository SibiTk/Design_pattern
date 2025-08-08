package main

import (
    "fmt"
    "strings"
)

type QueryBuilder struct {
    table   string
    columns []string
    where   string
    orderBy string
    limit   int
}

func NewQueryBuilder(table string) *QueryBuilder {
    return &QueryBuilder{
        table: table,
        limit: -1, 
    }
}

func (qb *QueryBuilder) Select(columns ...string) *QueryBuilder {
    qb.columns = columns
    return qb
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
    qb.where = condition
    return qb
}

func (qb *QueryBuilder) OrderBy(column string) *QueryBuilder {
    qb.orderBy = column
    return qb
}

func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
    qb.limit = limit
    return qb
}

func (qb *QueryBuilder) Build() string {
    var sb strings.Builder

  
    sb.WriteString("SELECT ")
    if len(qb.columns) > 0 {
        sb.WriteString(strings.Join(qb.columns, ", "))
    } else {
        sb.WriteString("*")
    }

   
    sb.WriteString(" FROM ")
    sb.WriteString(qb.table)


    if qb.where != "" {
        sb.WriteString(" WHERE ")
        sb.WriteString(qb.where)
    }

    
    if qb.orderBy != "" {
        sb.WriteString(" ORDER BY ")
        sb.WriteString(qb.orderBy)
    }


    if qb.limit >= 0 {
        sb.WriteString(fmt.Sprintf(" LIMIT %d", qb.limit))
    }

    sb.WriteString(";")
    return sb.String()
}


func main() {
	query := NewQueryBuilder("users").
		Select("name", "age").
		Where("age > 30").
		OrderBy("age").
		Limit(10).
		Build()

	fmt.Println(query)
}