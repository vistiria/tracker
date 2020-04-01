from tornado import ioloop, httpclient
import json
import random

i = 0

def get_token():
    http_client = httpclient.HTTPClient()
    response = http_client.fetch("http://127.0.0.1:6382/auth", method='GET')
    body = json.loads(response.body)
    return str(body["token"])

def handle_request(response):
    # print(response.code)
    global i
    i -= 1
    if i == 0:
        ioloop.IOLoop.instance().stop()

token1 = get_token()
token2 = get_token()
def prepare_data():
    ports = ["6382", "6383", "6384"]
    urls = []
    urls.extend([["http://127.0.0.1:{}/test".format(ports[0]), token1]] * 500)
    urls.extend([["http://127.0.0.1:{}/test".format(ports[1]), token1]] * 300)
    urls.extend([["http://127.0.0.1:{}/test".format(ports[2]), token1]] * 200)
    urls.extend([["http://127.0.0.1:{}/test".format(ports[2]), "invalid"]] * 500)
    urls.extend([["http://127.0.0.1:{}/test".format(ports[1]), "invalid"]] * 500)
    urls.extend([["http://127.0.0.1:{}/test".format(ports[0]), token2]] * 500)
    urls.extend([["http://127.0.0.1:{}/test".format(ports[1]), token2]] * 600)
    urls.extend([["http://127.0.0.1:{}/test2".format(ports[0]), token1]] * 300)
    urls.extend([["http://127.0.0.1:{}/test2".format(ports[1]), token1]] * 400)
    urls.extend([["http://127.0.0.1:{}/test2".format(ports[2]), token1]] * 200)
    random.shuffle(urls)
    return urls

print("sending 4003 requests...")

http_client = httpclient.AsyncHTTPClient()
for url,token in prepare_data():
    i += 1
    http_client.fetch(url, handle_request, method='GET', headers={"Authorization": token})
ioloop.IOLoop.instance().start()


http_client = httpclient.HTTPClient()
def get_counter(url, token):
    response = http_client.fetch(url, method='GET', headers={"Authorization": token})
    body = json.loads(response.body)
    return body["count"]

print(get_counter("http://127.0.0.1:6382/test", token1) == 1001)
print(get_counter("http://127.0.0.1:6382/test2", token1) == 901)
print(get_counter("http://127.0.0.1:6382/test", token2) == 1101)
