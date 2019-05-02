secure dns via https

Mainly follows https://fardog.io/blog/2017/12/30/client-side-certificate-authentication-with-nginx/
but for client certificate, since golang doesn't support des3 (tls.LoadX509KeyPair always report tls: failed to parse private key),
need change below procedure to create client certificate

openssl ecparam -out client.key -name prime256v1 -genkey
openssl req -new -key client.key -out client.csr
openssl x509 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out client.crt