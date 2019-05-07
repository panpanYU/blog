package model

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	ArticleModel struct {
		*SqlModel
	}

	Article struct {
		Id         int64     `json:"id" db:"id"`
		UId        int64     `json:"uid" db:"uid"`
		Title      string    `json:"title" db:"title"`
		Content    string    `json:"content" db:"content"`
		Status     int8      `json:"status" db:"status"`
		Clap       int       `json:"clap" db:"clap"`
		Terrible   int       `json:"terrible" db:"terrible"`
		Visits     int       `json:"visits" db:"visits"`
		Replies    int       `json:"replies" db:"replies"`
		Shared     int       `json:"shared" db:"shared"`
		CreateTime time.Time `json:"ctime" db:"ctime"`
		UpdateTime time.Time `json:"mtime" db:"mtime"`
	}
)

const articleQueryRow = `id,uid,content,status,clap,terrible,visits,replies,shared,ctime,mtime`

func NewArticleModel(conn sqlx.DB, table string) *ArticleModel {
	return &ArticleModel{SqlModel: NewSqlModel(conn, table)}
}

func (am *ArticleModel) Insert(article *Article) (int64, error) {
	querySql := `insert into ` + am.table + ` (uid,title,content,status) VALUES ($1,$2,$3,$4)`
	result, err := am.conn.Exec(querySql, article.UId, article.Title, article.Content, article.Status)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return id, nil
}

func (am *ArticleModel) FindById(id int64) (*Article, error) {
	querySql := `select ` + articleQueryRow + ` from ` + am.table + ` where id=?`
	var resp Article
	err := am.conn.Get(&resp, querySql, id)
	return &resp, err
}

func (am *ArticleModel) Delete(id int64) error {
	querySql := `delete from ` + am.table + ` where id=?`
	_, err := am.conn.Exec(querySql, id)
	return err
}

func (am *ArticleModel) Update(a *Article) error {
	querySql := `update ` + am.table + ` set uid=?,title=?,content=?,status=?,clap=?terrible=?visits=?,replies=?
						shared=? where id=?`
	_, err := am.conn.Exec(querySql, a.UId, a.Title, a.Content, a.Status, a.Clap, a.Terrible, a.Visits, a.Replies, a.Status, a.Id)
	return err
}
