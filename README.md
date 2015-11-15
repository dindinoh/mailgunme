# mailgunme
binary to fastly be fetched from bintray and used from cli to send emails using https

# howto
Save config to ~/.mailgunme

[mailgun]
domain=domainname.com
publickey=public-key-from-mailgun
privatekey=private-key-from-mailgun

Usage:
mailgunme -s "testing subject" -f "Name" -m "A nice message." -t recipient@mailaddressexample.com

# ToDo
- accept message as stdin
