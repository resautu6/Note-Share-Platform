import requests

user = {
    "name": "xiaoming",
    "password": "123"
}

testcase = {
    "article_title" : "hello world!",
    "article_content" : "this is a test article",
    "image_num" : 1,
}

image_path = ["D:\\document\\web_application_dev\\Note-Share-Platform\\backend\\test\\1241.png", "D:\\document\\web_application_dev\\Note-Share-Platform\\backend\\test\\123.png"]

def login():
    response = requests.post("http://127.0.0.1:8081/login", data=user)
    if response.status_code == 200:
        return response.json()["token"]

def getFiles():
    files = [
        ('image_list', ('file1.png', open(image_path[0], 'rb'), 'image/png')),
        ('image_list', ('file2.png', open(image_path[1], 'rb'), 'image/png')),
        # 添加更多文件...
    ]


    # files = {'image_list': []}
    # for file_path in image_path:
    #     files['image_list'].append(open(file_path, 'rb'))
    # files = {'image_list': open(image_path[0], 'rb')}
    return files

def test():
    token = login()
    headers = {
        "Authorization": token
    }
    files = getFiles()
    testcase["image_num"] = len(files)
    response = requests.post("http://127.0.0.1:8081/upload_article", data=testcase, headers=headers, files=files)
    print(response.status_code)
    print(response.json())

if __name__ == "__main__":
    test()