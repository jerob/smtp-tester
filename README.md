# SMTP Tester

This tool is useful to check the health status of a SMTP server and to know if everything is running righ.
You can sending a test email so you can be sure that it is functioning correctly.

```
docker run \
-e  FROM=emailfrom@gmail.com \
-e  TO=emailto@gmail.com \
-e  HOST=smtp.gmail.com \
-e  PORT=587 \
-e  USERNAME=emailfrom@gmail.com \
-e  PASSWORD=emailfrompassword \
-e  AUTH=1 \
-e  ENCRYPTION=tls   \
-e  SUBJECT="Hello !" \
-e  BODY="That's work !" \
smtp-tester
```
