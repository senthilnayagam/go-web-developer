require 'rubygems'
require 'bundler/setup'
require 'mysql2'
require 'sinatra'
require 'json'
require 'connection_pool'
require './common/common.rb'


set :port, 4567

 client = Mysql2::Client.new

$clientpool=  ConnectionPool.new(:size => 50, :timeout => 5) { Mysql2::Client.new(:host => "localhost", :username => "root",:password =>'root',:database => 'dbapi') } 



def query_db(sql)
    result = ''
  $clientpool.with do |c|
    result= c.query(sql, :as => :hash)
  end
   result
end  



get '/mysql2/:table/:id' do
  
  tables = ["user","car"]
  id =  client.escape(params[:id])
  table = client.escape(params[:table])
  
  if tables.include? table
    results = query_db("select * from #{table} where id=#{id};")
    content_type :json
    results.to_a.to_json
  else
   content_type :json
   status 404
   {"error"=>"table does not exist","code"=>404}.to_json
  end    
    
end

