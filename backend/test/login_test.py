import requests

testcase1 = {
    "name": "xiaoming",
    "password": "123"
}

testcase2 = {
    "name" : "xiaoming",
    "password": "44444"
}

testcase3 = {
    "name" : "zhangsan",
    "password": "44444"
}

testcases = [testcase1, testcase2, testcase3]

def test():
    for i, testcase in enumerate(testcases):
        response = requests.post("http://127.0.0.1:8081/login", data=testcase)
        print(f"testcase {i + 1}: {response.json()}")

if __name__ == '__main__':
    test()