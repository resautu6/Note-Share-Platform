import requests

user = {
    "name": "xiaoming",
    "password": "123"
}

url = "http://127.0.0.1:8081/"
# url = "http://resautu.cn:7879/"


def login():
    response = requests.post(url + "login", data=user)
    if response.status_code == 200:
        return response.json()["token"]
    
article_id = 16
def add_favourite_test():
    token = login()
    headers = {
        "Authorization": token,
    }
    data = {
        "article_id" : article_id,
        "command" : "add"
    }
    response = requests.post(url + "user/favourites", data=data, headers=headers)
    print(response.status_code)
    print(response.json())

def add_repeat_test():
    token = login()
    headers = {
        "Authorization": token
    }
    data = {
        "article_id" : article_id,
        "command" : "add"
    }
    response = requests.post(url + "user/favourites", data=data, headers=headers)
    print(response.status_code)
    print(response.json())

def get_favourite_list():
    token = login()
    headers = {
        "Authorization": token
    }

    response = requests.get(url + "user/favourites", headers=headers)
    print(response.status_code)
    print(response.json())

def delete_favourite_test():
    token = login()
    headers = {
        "Authorization": token
    }
    data = {
        "article_id" : article_id,
        "command" : "delete"
    }
    response = requests.post(url + "user/favourites", data=data, headers=headers)
    print(response.status_code)
    print(response.json())


if __name__ == "__main__":
    # add_favourite_test()
    # add_repeat_test()
    delete_favourite_test()
    # get_favourite_list()
