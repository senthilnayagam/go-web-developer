require 'rubygems'
require 'bundler/setup'
require 'pg'
require 'sinatra'
require 'json'
require 'connection_pool'
require './common/common.rb'

set :port, 4568


db  = PG.connect()
$pgpool = ConnectionPool.new(:size => 50, :timeout => 5) { PG.connect( dbname: 'dbapi' ) }

def query_db_pg(sql)

  result = ''
$pgpool.with do |c|
  result= c.exec(sql)
end
 result
end

get '/pg/:table/:id' do
  
  tables = ["users","cars"]
  id =  db.escape_string(params[:id])
  table = db.escape_string(params[:table])
  
  if tables.include? table
    results = query_db_pg("select * from #{table} where id=#{id};")
    content_type :json
    results.to_a.to_json
  else
   content_type :json
   status 404
   {"error"=>"table does not exist","code"=>404}.to_json
  end    
    
end