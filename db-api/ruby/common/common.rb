require 'sinatra'
require 'json'
require 'connection_pool'


# sinatra config
set :logging, true
set :environment, :production
set :root, File.dirname(__FILE__)
set :bind, '0.0.0.0'


log = File.new("sinatra.log", "a+")
#$stdout.reopen(log)
$stderr.reopen(log)