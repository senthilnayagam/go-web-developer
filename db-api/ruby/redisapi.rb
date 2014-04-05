require 'rubygems'
require 'bundler/setup'
require "redis"
require 'sinatra'
require 'json'
require 'connection_pool'
require './common/common.rb'







set :port, 4569



$redispool = ConnectionPool.new(:size => 50, :timeout => 5) { Redis.new }

def query_redis(key)

  result = ''
$redispool.with do |c|
  result= c.get(key)
end
 result
end

get '/redis/:key' do
    results = query_redis(params[:key])
    content_type :json
    results.to_json

#  else
#   content_type :json
#   status 404
#   {"error"=>"table does not exist","code"=>404}.to_json
  end    


get '/redisset/:key/:value' do 

result =  Redis.new.set(params[:key], params[:value])

content_type :json
result.to_json

end

    
