import requests
for _ in range(20):
    print(requests.post("http://localhost:5000/test-endpoint").text)