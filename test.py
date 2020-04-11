import requests

r=requests.get("http://localhost:8000/")
print(r.content)
print(r.status_code)