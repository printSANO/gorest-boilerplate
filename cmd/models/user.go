package models

import "time"

type Example struct {
	ID        int       `db:"id" dataType:"integer" constraint:"PRIMARY KEY"`
	Name      string    `db:"name" dataType:"varchar(20)"`
	Password  string    `db:"password" dataType:"varchar(20)"`
	CreatedAt time.Time `db:"created_at" dataType:"timestamp"`
}

type Example2 struct {
	ID        int       `db:"id" dataType:"integer" constraint:"PRIMARY KEY"`
	Title     string    `db:"title" dataType:"varchar(40)"`
	Author    string    `db:"author" dataType:"varchar(20)"`
	CreatedAt time.Time `db:"created_at" dataType:"timestamp"`
}
