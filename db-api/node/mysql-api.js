var express = require('express');
var app = express();




var mysql = require('mysql');
var pool = mysql.createPool({
	host: 'localhost',
	database: 'dbapi',
	user: 'root',
	password: 'root',
	queueLimit: 0
});
var http = require('http');
var url = require('url');


/*
pool.getConnection(function(err, connection) {
  // Use the connection
  connection.query( 'SELECT * FROM car where id=1', function(err, rows) {   //,fields
    // And done with the connection.
	if (err) throw err;
	console.log(rows)
	// console.log(fields)
    connection.release();

    // Don't use the connection here, it has been returned to the pool.
  });
});
*/

/*
res.setHeader({
	'Content-Type': 'application/json'
});


app.get('/', function(req, res) {
	res.send('Hello World');
});
*/

app.get('/mysql2/:table/:id', function(req, res){
	pool.getConnection(function(err, connection) {
		if (err) {
			console.error('CONNECTION error: ', err);
			res.statusCode = 503;
			res.send({
				result: 'error',
				err: err.code
			});
		} else {
			connection.query('SELECT * FROM ' + req.params.table + ' WHERE id = ?', req.params.id, function(err, rows) {
				if (err) {
					console.error(err);
					res.statusCode = 500;
					res.send({
						result: 'error',
						err: err.code
					});
				} else {
					res.send({
						result: 'success',
						err: '',
						id: req.params.id,
						json: rows[0],
						length: 1
					});
					connection.release();
				}
			});
		}
	});
});

app.get('/mysql/:table/:id', function(req, res){
	pool.getConnection(function(err, connection) {
		if (err) {
			console.error('CONNECTION error: ', err);
			res.statusCode = 503;
			res.send({
				result: 'error',
				err: err.code
			});
		} else {
			connection.query('SELECT * FROM ' + req.params.table + ' WHERE id = ?', req.params.id, function(err, rows) {
				if (err) {
					console.error(err);
					res.statusCode = 500;
					res.send({
						result: 'error',
						err: err.code
					});
				} else {
					res.send(
						rows
					);
					connection.release();
				}
			});
		}
	});
});







app.listen(3000);
console.log('app running at port 3000');
console.log('http://127.0.0.1/mysql/:table/:id');

