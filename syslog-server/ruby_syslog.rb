#!/usr/bin/env ruby

require 'socket'
require 'io/wait'
require 'syslog'


# http://stackoverflow.com/questions/18027146/inspect-udp-syslog-packets-in-ruby
# just saving the input to file
# http://tools.ietf.org/html/rfc5424


class Server
    def initialize
      @log = File.open("commonlogger.log", "a")
        @listener = UDPSocket.new
        @listener.bind("127.0.0.1", "514")
        getdata
        
    end

    def getdata
        while true
            @text, @sender = @listener.recvfrom(9000)
            p @listener
            p @text
            p @sender
            
            
@log.puts   @text.to_s #+ @sender.to_s  + @listener.to_s
 
        end
    end
end

x = Server.new