A Fake SMTP server for debugging/testing email issues from web application




test via telnet

 telnet 127.0.0.1 2525

escape character in telnet is ctrl + ]
in telnet console enter q to quit
telnet> q #to quit
telnet> set crlf # to send \r\n as specified in http/smtp protocols

./fss.go:253: non-declaration statement outside function body
./fss.go:254: syntax error: unexpected }




todo

option to log to a log file, especially when doing test runs or a chatty application

a fake sendmail clone