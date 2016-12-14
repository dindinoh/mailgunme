[![Build Status](https://api.shippable.com/projects/564a2e221895ca4474239d2d/badge)](https://api.shippable.com/projects/564a2e221895ca4474239d2d/badge)


# mailgunme
binary to fastly be fetched from bintray and used from cli to send emails using https to mailgun

# howto
Save config to ~/.mailgunme

[mailgun]  
domain=domainname.com  
publickey=public-key-from-mailgun  
privatekey=private-key-from-mailgun   

Usage:  
mailgunme -s "testing subject" -f "Name" -m "A nice message." -t recipient@mailaddressexample.com

# ToDo
- attachments
