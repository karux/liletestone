openssl req -x509 -newkey rsa:4096 -keyout ./keystore/server/server-key.pem -out ./keystore/server/server-cert.pem -days 365 -nodes -subj '/CN=localhost'
