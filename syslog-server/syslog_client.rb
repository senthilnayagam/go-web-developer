require 'syslog'



=begin
Syslog.log(Syslog::LOG_CRIT, "Out of disk space")
Syslog.log(Syslog::LOG_CRIT, "User %s logged in", ENV['USER'])



Syslog.log(Syslog::LOG_ALERT, "Out of memory")
Syslog.alert("Out of memory")


=end

Syslog.alert("syslog script is running")