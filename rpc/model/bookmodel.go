/*
 * @Author: ming
 * @LastEditors: ming
 * @Date: 2020-11-14 20:00:22
 * @LastEditTime: 2020-11-14 20:09:38
 * @FilePath: /zero-demo/rpc/model/bookmodel.go
 * @Description: Description
 */
package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	bookFieldNames          = builderx.FieldNames(&Book{})
	bookRows                = strings.Join(bookFieldNames, ",")
	bookRowsExpectAutoSet   = strings.Join(stringx.Remove(bookFieldNames, "create_time", "update_time"), ",")
	bookRowsWithPlaceHolder = strings.Join(stringx.Remove(bookFieldNames, "book", "create_time", "update_time"), "=?,") + "=?"

	cacheBookBookPrefix = "cache#Book#book#"
)

type (
	BookModel struct {
		sqlc.CachedConn
		table string
	}

	Book struct {
		Book  string `db:"book"`  // book name
		Price int64  `db:"price"` // book price
	}
)

func NewBookModel(conn sqlx.SqlConn, c cache.CacheConf, table string) *BookModel {
	return &BookModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      table,
	}
}

func (m *BookModel) Insert(data Book) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, bookRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Book, data.Price)

	return ret, err
}

func (m *BookModel) FindOne(book string) (*Book, error) {
	bookBookKey := fmt.Sprintf("%s%v", cacheBookBookPrefix, book)
	var resp Book
	err := m.QueryRow(&resp, bookBookKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where book = ? limit 1", bookRows, m.table)
		return conn.QueryRow(v, query, book)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *BookModel) Update(data Book) error {
	bookBookKey := fmt.Sprintf("%s%v", cacheBookBookPrefix, data.Book)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where book = ?", m.table, bookRowsWithPlaceHolder)
		return conn.Exec(query, data.Price, data.Book)
	}, bookBookKey)
	return err
}

func (m *BookModel) Delete(book string) error {

	bookBookKey := fmt.Sprintf("%s%v", cacheBookBookPrefix, book)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where book = ?", m.table)
		return conn.Exec(query, book)
	}, bookBookKey)
	return err
}

func (m *BookModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBookBookPrefix, primary)
}

func (m *BookModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where book = ? limit 1", bookRows, m.table)
	return conn.QueryRow(v, query, primary)
}
