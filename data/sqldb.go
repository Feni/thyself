package data

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	"thyself/log"
)

// mysql.New
var sqldb, Sqlconn_err = sql.Open("postgres", "user=goclient password=gothy@0 dbname=thydb")

// user: email, user_id, pass_hash,
// journal_entries: user_id, je_time, je_text
// metric_entries: user_id, me_id, me_time, me_action, me_description
// metric_details: me_id, d_group, d_type, d_amount

// User account queries
var SQL_CREATE_USER, Sql_create_err = sqldb.Prepare("INSERT INTO users (email, user_id, pass_hash) VALUES ($1, $2, $3)")
var SQL_RETRIEVE_PASS, Sql_pass_err = sqldb.Prepare("SELECT pass_hash, user_id FROM users where email=$1 LIMIT 1")

// Metric queries
var SQL_RETRIEVE_ALL_METRICS, Sql_all_metrics_err = sqldb.Prepare("SELECT me_id, EXTRACT(EPOCH FROM(me_time))::int, me_action, me_description FROM metric_entries where user_id=$1 ORDER BY me_time")
var SQL_RETRIEVE_METRIC_DATE_RANGE, Sql_metric_date_range_err = sqldb.Prepare("SELECT me_id, EXTRACT(EPOCH FROM(me_time))::int, me_action, me_description FROM metric_entries where user_id=$1 AND me_time BETWEEN to_timestamp($2) AND to_timestamp($3) ORDER BY me_time")
var SQL_RETRIEVE_DETAILS, Sql_details_err = sqldb.Prepare("SELECT d_group, d_type, d_amount FROM metric_details WHERE me_id = $1")

// Create metrics
var SQL_ADD_METRIC, Sql_add_metric_err = sqldb.Prepare("INSERT INTO metric_entries (user_id, me_id, me_time, me_action, me_description) VALUES ($1, $2, to_timestamp($3), $4, $5)")
var SQL_ADD_DETAIL, Sql_add_detail_err = sqldb.Prepare("INSERT INTO metric_details (me_id, d_group, d_type, d_amount) VALUES ($1, $2, $3, $4)")
var SQL_ADD_DETAIL_NO_AMT, Sql_add_detail_no_amt_err = sqldb.Prepare("INSERT INTO metric_details (me_id, d_group, d_type) VALUES ($1, $2, $3)")

// je user_id, je_time, je_text
// Journal entry
var SQL_RETRIEVE_JE, Sql_get_je_err = sqldb.Prepare("SELECT EXTRACT(EPOCH FROM(je_time))::int, je_text FROM journal_entries where user_id=$1 AND je_time BETWEEN to_timestamp($2) AND to_timestamp($3) LIMIT 1")
var SQL_ADD_JE, Sql_add_je_err = sqldb.Prepare("INSERT INTO journal_entries (user_id, je_time, je_text) VALUES ($1, to_timestamp($2), $3)")
var SQL_UPDATE_JE, Sql_update_je_err = sqldb.Prepare("UPDATE journal_entries SET je_time=to_timestamp($1), je_text=$2 WHERE user_id=$3 AND je_time BETWEEN to_timestamp($4) AND to_timestamp($5)")

func SqlInit() {
	log.Debug(Sqlconn_err, "  Database connection error ")
	log.Debug(Sql_create_err, "  Sql Error create user: ")
	log.Debug(Sql_pass_err, "  Sql Error pass: ")
	log.Debug(Sql_all_metrics_err, "  Sql Error all mets: ")
	log.Debug(Sql_metric_date_range_err, "  Sql Error date mets: ")
	log.Debug(Sql_add_metric_err, "  Sql error add metrics error: ")
	log.Debug(Sql_add_detail_err, "  Sql Error add details: ")
	log.Debug(Sql_add_detail_no_amt_err, "  Sql Error add detail: ")
	log.Debug(Sql_details_err, "  Sql retrieve details: ")
	log.Debug(Sql_get_je_err, "  Sql get journal entry: ")
	log.Debug(Sql_add_je_err, "  Sql add journal entry: ")
	log.Debug(Sql_update_je_err, "  Sql update journal entry: ")
}
