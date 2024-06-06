import requests

# url = "http://127.0.0.1:8081/"
url = "http://resautu.cn:7879/"

user = {
    "name": "xiaoming",
    "password": "123"
}

def login():
    response = requests.post(url + "login", data=user)
    if response.status_code == 200:
        return response.json()["token"]

def random_test():
    response = requests.get(url + "article/list")
    print(response.status_code)
    print(response.json())

def get_specific_test():
    response = requests.get(url + "article/17")
    print(response.status_code)
    print(response.json())

    response = requests.get(url + "article/17/content")
    print(response.status_code)
    

    content = response.json()
    print(content)

    image_path = url + content["image_path"] + "/0.png"
    response = requests.get(image_path)
    print(response.status_code)

    with open("test.png", "wb") as f:
        f.write(response.content)

def get_user_article_test():
    token = login()
    headers = {
        "Authorization": token
    }

    response = requests.get(url + "user/article", headers=headers)
    print(response.status_code)
    print(response.json())

def search_artitle():
    response = requests.get(url + "search/article/double")
    print(response.status_code)
    print(response.json())

def delete_article_test():
    token = login()
    headers = {
        "Authorization": token
    }

    data = {
        "command": "delete"
    }
    response = requests.post(url + "article/14/modify", data=data, headers=headers)
    print(response.status_code)
    print(response.json())

if __name__ == "__main__" :
    # get_user_article_test()
    get_specific_test()
    # delete_article_test()
    # get_specific_test()
    # random_test()
    # search_artitle()