


<template>
  <!-- <div id="app">
    
    <div v-if="!login&&!register">
    用户名：<input type="text" v-model="name">
    <br>
    密码：<input type="password" v-model="password">
    <br>
    <button @click="LogIn">登陆</button> <button @click="Regis">注册</button>
    </div>
    <div v-else-if="register">
      用户名：<input type="text" v-model="name">
    <br>
    密码：<input type="password" v-model="password">
    <br>
    确认密码：<input type="password" v-model="acquire">
    <br>
      <button @click="BackToLogIn">返回</button><button @click="Register">注册</button>
    </div>
  
  </div> -->
  <!-- <button @click="Init">清空asd</button> -->

<body>
  <div class="login-container" v-if="!login&&!register">
        <h2>Login</h2>
          <input type="text" name="username" placeholder="Username" v-model="name">
          <input type="password" name="password" placeholder="Password" v-model="password">
          <button @click="LogIn">Login</button>
          <button @click="Regis">Register</button>
    </div>
    <div class="login-container" v-else-if="register">
        <h2>Register</h2>
          <input type="text" name="username" placeholder="Username" v-model="name">
          <input type="password" name="password" placeholder="Password" v-model="password">
          <input type="password" name="acquire" placeholder="Confirm" v-model="acquire">
          <button @click="BackToLogIn">Back</button>
          <button @click="Register">Register</button>
    </div>
    <div class="container" v-else-if="login">
      <input type="text" @keyup.F5="LogOut" style="display: none;">
      <div class="navigation-div" >
        <ul>
          <li><button @click="ToExplore">发现</button></li>
          <li><button @click="ToPost">发布</button></li>
          <li><button @click="ToHome">我</button></li>
          <li><button @click="ToFavor">收藏</button></li>
          <li v-if="!(display&&searchnote)"><button v-if="!(display&&searchnote)" @click="LogOut">退出登陆</button></li>
          <li v-if="post"><button @click="addImageBox" v-if="post">上传图片</button></li>
          <li v-if="post"><button @click="UploadArticle" v-if="post">发布笔记</button></li>
          <li v-if="display&&!(displaynote.id in favornotesById)">
            <button @click="favourites" v-if="display&&!(displaynote.id in favornotesById)">收藏笔记</button>
          </li>
          <li v-if="display&&(displaynote.id in favornotesById)">
            <button @click="rmFavor" v-if="display&&(displaynote.id in favornotesById)">取消收藏</button>
          </li>
          <li v-if="display&&(displaynote.id in mynotesById)">
            <button @click="deleteNote" v-if="display&&(displaynote.id in mynotesById)">删除笔记</button>
          </li>
          <li v-if="display&&searchnote">
            <button @click="BackToSearch" v-if="display&&searchnote">返回搜索</button>
          </li>
        </ul>
        </div>

        <!-- <div class="button-div" v-if="post">
            <button @click="addImageBox">上传图片</button>
            <button @click="UploadArticle">发布</button>
            </div> -->

        <!-- <div class="main-div" v-if="true">
          <div v-for="(note, index) in notes" :key="index" class="note-div">
            <img :src="imageSrc" class="note-picture">
            <div class="explore-note-title">
              <p>晓山瑞希mzk晓山瑞希mzk晓山瑞希mzk晓山瑞希mzk晓山瑞希mzk晓山瑞希mzk晓山瑞希mzk晓山瑞希mzk</p>
            </div>
          </div>
        </div> -->
        <div class="main-div" v-if="post">
          <div class="explore-note-text" id="title">
            <textarea placeholder="请输入标题" v-model="postTitle"></textarea>
          </div>
          <div class="explore-note-text">
            <textarea placeholder="请输入文本" v-model="postText"></textarea>
          </div>
          <div class="explore-note-picture">
            <div
              v-for="(image, index) in images"
              :key="index"
              class="image-box"
              @mouseover="imageBoxHovered[index] = true"
              @mouseleave="imageBoxHovered[index] = false"
            >
              <img v-if="image" :src="image" alt="Uploaded Image">
              <button
                  v-if="imageBoxHovered[index]"
                  class="delete-button"
                  @click.stop="removeImage(index)"
                >X</button>
            </div>
            <input
              type="file"
              ref="fileInput"
              @change="handleFileUpload"
              multiple accept=".jpg, .jpeg, .png"
              style="display: none;"
            >
          </div>
        </div>

        <div class="main-div" v-if="explore">
          <div class="search-box">
            <input type="text" placeholder="搜索笔记" v-model="searchKeyWords" @keydown.enter="searchNote">
            <button @click="searchNote" >Search</button>
          </div>
          <div v-for="(note, index) in notesById" :key="index" class="note-div" @click="OpenNote(note)">
            <img :src="`http://resautu.cn:7879/${note.image_path}/0.png`"  
            class="note-picture" onerror="this.onerror=null; this.src='default_image_path.jpg'">
            <div class="explore-note-title">
              <p>{{note.title}}</p>
            </div>
          </div>
        </div>

        <div class="main-div" v-if="home">
          <div v-for="(note, index) in mynotesById" :key="index" class="note-div" @click="OpenNote(note)">
            <img :src="`http://resautu.cn:7879/${note.image_path}/0.png`"  class="note-picture">
            <div class="explore-note-title">
              <p>{{note.title}}</p>
            </div>
          </div>
        </div> 
        
        <div class="main-div" v-if="favor">
          <div v-for="(note, index) in favornotesById" :key="index" class="note-div" @click="OpenNote(note)">
            <img v-if="note.invalid"  class="note-picture" src="http://resautu.cn:7879/res/img404.png">
            <img v-else :src="`http://resautu.cn:7879/${note.image_path}/0.png`"  class="note-picture">
            <div class="explore-note-title">
              <p>{{note.title}}</p>
            </div>
          </div>
        </div>
        <div class="main-div" v-if="searchnote&&!display">
          <div v-for="(note, index) in searchnotesById" :key="index" class="note-div" @click="OpenNote(note)">
            <img :src="`http://resautu.cn:7879/${note.image_path}/0.png`"  class="note-picture">
            <div class="explore-note-title">
              <p>{{note.title}}</p>
            </div>
          </div>
        </div>

        <div class="main-div" v-if="display">
          
          <div class="note-main-div">
            <h1 style="font-size: 40px; margin: 20px;">用户：{{displaynote.name}}</h1>
            <p style="font-size: 20px; margin: 20px;">发布时间：{{displaynote.modify_time}}</p>
            <p style="font-size: 20px; margin: 20px;">浏览量：{{displaynote.view_num}}</p>
            <h2 style="font-size: 60px; margin: 20px;">{{displaynote.title}}</h2>
            <p
              v-for="(line, index) in lines()"
              :key="index"
              style="font-size: 30px; margin: 20px; white-space: pre-wrap;"
            >{{ "    " + line }} </p>
      
          </div>

          <div class="noteImgDiv" @mouseover="enterNoteImg" @mouseleave="leaveNoteImg">
            <img v-if="displaynote.invalid" class="noteImage" :src="`http://resautu.cn:7879/res/img404.png`">
            <img v-else class="noteImage" :src="`http://resautu.cn:7879/${displaynote.image_path}/${displaynote.image_idx}.png`">
            <button v-if="isButtonVisible" style="left:5%" @click="leftImg">&lt;</button>
            <button v-if="isButtonVisible" style="right: 5%;" @click="rightImg">&gt;</button>
          </div>
          
        </div>

    </div>


  </body>
</template>



<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    body {
        font-family: Arial, sans-serif;
        /* background-color: #f4f4f4; */
        margin: 0;
        padding: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        /* height: 100%; */
        /* background-image: url('2_low0.3.svg');
        background-repeat: no-repeat;
        background-size: cover; */
    }
    .login-container {
        /* background-color: #fff; */
        border-radius: 8px;
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        padding: 20px;
        max-width: 600px;
        height: 500px;
        width: 100%;
        /* margin: auto; */
        display: flex;
        flex-direction: column; /* 将子元素排列在列中 */
        justify-content: center; /* 垂直居中 */
    }
    .login-container h2 {
        text-align: center;
        font-size: 50px;
    }
    .login-container input[type="text"],
    .login-container input[type="password"] {
        width: 70%;
        height: 13%;
        padding: 10px;
        margin-left: auto;
        margin-right: auto;
        margin-top: 20px;
        border: 1px solid #ccc;
        border-radius: 4px;
        box-sizing: border-box;
        font-size: 20px;
    }
    .login-container button {
        width: 35%;
        /* padding: 10px; */
        height: 10%;
        background-color: #007bff;
        color: #fff;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        margin-left: auto;
        margin-right: auto;
        margin-top: 30px;
        font-size: 20px;
    }
    .login-container button:hover {
        background-color: #0056b3;
    }
    .main-div{
      width:calc(98vw - 200px);
      height:100%;
      /* background-color: black; */
      float:right;
    }
    .note-main-div{
      width: 30vw;
      height: 100%;
      /* background-color: red; */
      float: left;
    }
    .noteImgDiv{
      width: 50vw;
      height: 40vw;
      position: fixed;
      left: 50%;
      /* background-color: green; */
    }
    .noteImgDiv button{
      font-size: 70px;
      font-weight: 100;
      z-index: indx + 1;
      position: absolute;
      top: 50%;
      width:  100px;
      height: 100px;
      transform: translateY(-50%);
      border-radius: 50px;
      background-color: rgb(255, 255, 255, 0);
      border-width: 2px ;
    }
    .noteImage {
      max-width: 100%; /* 图片的最大宽度为容器的100% */
      max-height: 100%; /* 图片的最大高度为容器的100% */
      position: absolute; /* 绝对定位以便在容器内居中 */
      top: 50%;
      /* left: 50%; */
      transform: translate(-50%, -50%); /*居中图片*/
      object-fit: contain; /* 保持图片的长宽比，同时确保它不会超出容器的边界 */
    }


    .note-div{
      width: 360px;
      height:550px;
      /* background-color: red; */

      margin:1%;
      float: left;
      position: relative;

      overflow: hidden; /* 隐藏超出div的文本 */
      word-wrap: break-word; /* 允许长单词换行到下一行 */
      
    }
    .note-picture {
      width: calc(100% - 2px);
      height: 80%;
      object-fit: cover; /* 保持图片的宽高比，并填充整个容器 */
      border-radius: 30px;
      margin: 0;
      padding: 0;
      border: 1px solid black;
    }
    .explore-note-title{
      width:100%;
      height: 20%;
      font-size: 24px;
      /* background-color: blue; */
      float: top;
      padding: 0;
      margin: 0;
      position: absolute;

    }
    .explore-note-title p{
      margin: 0; /* 去除段落的默认外边距 */
      overflow: hidden; /* 隐藏超出div的文本 */
      text-overflow: ellipsis; /* 文本超出部分显示省略号 */
    }

    .search-box {
    width: 70%;
    height : 50px;
    z-index: 2;
    margin: 0 auto;
    display: flex;
    align-items: center;
    border: 1px solid #ccc;
    padding: 10px;
    border-radius: 5px;
    /* background-color: white; */
  }

  .search-box input[type="text"] {
    flex-grow: 1;
    border: none;
    outline: none;
    padding: 5px;
    font-size: 35px;
    background-color: transparent;
  }

  .search-box button {
    background-color: #007bff;
    color: white;
    border: none;
    padding: 10px 15px;
    border-radius: 5px;
    cursor: pointer;
    font-size: 35px;
  }

  .search-box button:hover {
    background-color: #0056b3;
  }
    .navigation-div{
      width: 200px; /* 设置导航栏的宽度 */
      height: 100%;
      /* background-color: white; */
      float: left; /* 将导航栏浮动到最左边 */
      position:fixed;
    }
    .navigation-div ul{
      width: 100%;
      height:100%;
      list-style-type: none;
      padding: 0;
    }
    .navigation-div ul li{
      padding: 20px;
    }
    .navigation-div ul li button{
        width: 100%;
        padding: 20px;
        background-color: transparent;
        color: black;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 24px; /* 设置字体大小 */
    }
    .navigation-div ul li button:hover{
        background-color: #007bff;
        color: #fff;
    }
    .explore-note-text{
      width: 95%;
      height: 300px;
    }
    #title{
      width: 95%;
      height: 24px;
      margin-bottom: 20px;
    }
    .explore-note-text textarea{
      width: 100%;
      height: 100%;
      display:block;
      font-size: 20px;
    }
    
    .explore-note-picture{
      width: 95%;
      height: clac(100% - 300px);
      /* background-color: green; */
      margin-top: 2%;
    }
    
    
    .button-div{
      position: fixed;
      bottom: 5vw;
      left: 0vw;
      width: 10vw;
      padding: 10px;
      text-align: center;
      z-index: 2;

    }
    .button-div button {
        width: 150px;
        height: 70px;
        padding: 10px;
        background-color: #007bff;
        color: #fff;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 20px;
        margin-top: 20px;
    }
    .image-box {
      position: relative;
      width: 450px;
      height: 450px;
      border: 1px solid #ccc;
      justify-content: center;
      align-items: center;
      float: left;
      margin-left: 2%;
    }
    .image-box img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
    .delete-button {
      position: absolute;
      top: 0;
      right: 0;
      background: none;
      border: none;
      font-size: 35 px;
      color: red;
      cursor: pointer;
    }
    .container {
            width: 100%;
            height: 100%; /* 设置div的高度为整个区域的高度 */
            /* background-color: white; 设置div的背景颜色 */
        }
        .note {
            margin-bottom: 20px;
            padding: 10px;
            border: 1px solid #ccc;
            border-radius: 5px;
            /* background-color: #f9f9f9; */
        }
        .note h2 {
            margin-top: 0;
        }
        form {
            margin-top: 20px;
        }
        form textarea {
            width: 100%;
            height: 100px;
            resize: none;
        }
        form button {
            padding: 10px 20px;
            background-color: #007bff;
            color: #fff;
            border: none;
            border-radius: 5px;
            cursor: pointer;
        }
        form button:hover {
            background-color: #0056b3;
        }
        @media (max-width: 1000px) {
          .navigation-div{
            /* 当视口宽度小于600px时隐藏 */
            display: none;
          }
          .main-div{
            width: 100%;
          }
        }
</style>

<script>
import axios from 'axios';

import querystring from 'querystring'
export default {
  name: 'HelloWorld',
  data(){
    return{
      login:false,
      register:false,
      name: "",
      uname:"",
      token:"",
      password: "",
      acquire: "",

      explore: false,
      post: false,
      home: false,
      favor: false,
      display: false,
      searchnote: false,
      displaynote: undefined,
      searchKeyWords:"",

      // notesid: [],
      notesById: {},
      mynotesById: {},
      favornotesById: {},
      searchnotesById: {},
      notesId2Images: {},
      //notesnum: 0,
      noteBoxHovered: [], //用于跟踪笔记是否被鼠标悬停
      mynoteBoxHovered: [],
      favornoteBoxHovered:[],
      images: [],
      imageBoxHovered: [], // 用于跟踪每个图片框是否被鼠标悬停
      imagefiles: [],
      imageSrc: require('@/2D662B8B12645D48E376A61EB7AE138C.jpg'),
      postText: "",
      postTitle: "",
      isButtonVisible: false,
    }
  },
  mounted() {
    // 初始化一些数据
    this.notesById = {};
    //this.loadMoreNotes("explore");
    this.notesnum = 0;
    
    // 监听滚动事件
    window.addEventListener('scroll', this.handleScroll);
  },
  methods:{
    lines() {
      // 将内容按换行符分割成数组
      return this.displaynote.content.split('\n');
    },
    enterNoteImg(){
      this.isButtonVisible = true;
    },
    leaveNoteImg(){
      this.isButtonVisible = false;
    },
    leftImg(){
      console.log(this.displaynote.image_idx);
      this.displaynote.image_idx = (this.displaynote.image_idx + this.displaynote.image_num - 1)%this.displaynote.image_num;
      console.log(this.displaynote.image_idx);
    },
    rightImg(){
      this.displaynote.image_idx = (this.displaynote.image_idx + 1)%this.displaynote.image_num;
      console.log(this.displaynote.image_idx);
    },
    searchNote(){
      this.explore = false;
      this.searchnote = true;
      this.loadMoreNotes("search");
    },
    searchNoteByEnter(){
      console.log("keydown enter");
      if(this.isInput === true)
      {
        this.explore = false;
        this.searchnote = true;
        this.loadMoreNotes("search");
      }
    },
    BackToSearch(){
      this.display = false;
    },
    favourites(){
      var self = this;
      axios.post("http://resautu.cn:7879/user/favourites",querystring.stringify({
        article_id : self.displaynote.id,
        command : "add"
      }),{headers :{
          Authorization: self.token,
        }})
        .then(function(response){
        if(response.data !== undefined){
          alert(response.data.message);
          self.loadMoreNotes("favor");
        }
      })
      .catch(function(error){
        if(error.response.data !== undefined)
          alert(error.response.data.message);
        console.log("rmFavor error" + error);
      })
    },
    rmFavor(){
      var self = this;
      axios.post("http://resautu.cn:7879/user/favourites",querystring.stringify({
        article_id : self.displaynote.id,
        command : "delete"
      }),{headers :{
          Authorization: self.token,
        }})
      .then(function(response){
        if(response.data !== undefined){
          alert(response.data.message);
          delete self.favornotesById[self.displaynote.id];
        }
      })
      .catch(function(error){
        if(error.response.data !== undefined)
          alert(error.response.data.message);
        console.log("rmFavor error" + error);
      })
    },
    deleteNote(){
      var self = this;
      axios.post(`http://resautu.cn:7879/article/${self.displaynote.id}/modify`,querystring.stringify({
        command : "delete"
      }),{headers :{
          Authorization: self.token,
        }})
      .then(function(response){
        if(response.data !== undefined){
          alert(response.data.message);
          var id = self.displaynote.id;
          delete self.notesById[id];
          delete self.mynotesById[id];
          delete self.favornotesById[id];
          delete self.searchnotesById[id];
          self.ToHome();
        }
      })
      .catch(function(error){
        if(error.response.data !== undefined)
          alert(error.response.data.message);
        console.log("rmFavor error" + error);
      })
    },

    OpenNote(note){
      
      console.log(note);
      console.log(this.notesId2Images[note.id]);
      this.displaynote = note;
      this.displaynote.image_idx = 0;
      let id = this.displaynote.id;
      var self = this;
      axios.get(`http://resautu.cn:7879/article/${id}/content`)
              .then(function(response){
                self.displaynote.content = response.data.content;
                self.displaynote.view_num = response.data.view_num;
                self.display = true;
              })
              .catch(function(error){
                  self.displaynote.content = "笔记或许已经被删除";
                  self.displaynote.view_num = -1;
                  console.log("get note content error " + error);
                  self.display = true;
              });
              this.explore = false; this.home = false; this.favor = false; this.post = false;
      
    },

    loadMoreNotes(type) {
      // 这里应该是获取初始数据的逻辑
      console.log("type:" + type);
      var self = this;
      var URL1;
      if(type === "explore") URL1 = "http://resautu.cn:7879/article/list";
      if(type === "home") URL1 = "http://resautu.cn:7879/user/article";
      if(type === "favor") URL1 = "http://resautu.cn:7879/user/favourites";
      if(type === "search") URL1 = `http://resautu.cn:7879/search/article/${this.searchKeyWords}`;
      console.log("URL1:" + URL1);
      axios.get(URL1,{headers :{
          Authorization: self.token,
        }})
      .then(function(response){
        console.log("note sum: " + response.data.item_sum + " " + response.data.items.length);
        var notesById1;
        
        
        if(type === "explore") notesById1 = self.notesById;
        else if(type === "home") notesById1 = self.mynotesById;
        else if(type === "favor") notesById1 = self.favornotesById;
        else if(type === "search"){
          notesById1 = self.searchnotesById;
        }
        else {
          console.log("loadMoreNotes type error");
          return;
        }
        for(let i = 0; i < response.data.items.length; i++){

          if(!(response.data.items[i] in notesById1)){
            console.log("i:" + response.data.items[i]);
            let id = response.data.items[i];
            //self.notesid.push(id);
            axios.get(`http://resautu.cn:7879/article/${id}`)
            .then(function(response){
              if(type === "explore"){
                self.notesById[id] = response.data;
                //self.notesById[id].images = [];
                self.notesById[id].BoxHovered = false;
              }
              if(type === "home"){
                self.mynotesById[id] = response.data;
                //self.mynotesById[id].images = [];
                self.mynotesById[id].BoxHovered = false;
              }
              if(type === "favor"){
                self.favornotesById[id] = response.data;
                self.favornotesById[id].BoxHovered = false;
              }
              if(type === "search"){
                self.searchnotesById[id] = response.data;
                self.searchnotesById[id].BoxHovered = false;
              }
              self.notesId2Images[id] = [];
              // axios.get(`http://resautu.cn:7879/article/${id}/content`)
              // .then(function(response){
              //   if(type === "explore")
              //     self.notesById[id].content = response.data.content;
              //   if(type === "home")
              //     self.mynotesById[id].content = response.data.content;
              //   if(type === "favor")
              //     self.favornotesById[id].content = response.data.content;
              //   if(type === "search")
              //     self.searchnotesById[id].content = response.data.content;
              // })
              // .catch(function(error){
              //     console.log("get note content error " + error);
              // });
              var name;
              if(type === "explore")
                name = self.notesById[id].uname;
              if(type === "home")
                name = self.mynotesById[id].uname;
              if(type === "favor")
                name = self.favornotesById[id].uname;
              if(type === "search")
                name = self.searchnotesById[id].uname;
              console.log("name:" + name);
              axios.get(`http://resautu.cn:7879/user/uname/${name}`)
              .then(function(response){
                if(type === "explore")
                  self.notesById[id].name = response.data.uname;
                if(type === "home")
                  self.mynotesById[id].name = response.data.uname;
                if(type === "favor")
                  self.favornotesById[id].name = response.data.uname;
                if(type === "search")
                  self.searchnotesById[id].name = response.data.uname;
              })
              .catch(function(error){
                alert("get user name error");
                console.log(error);
              });

            })
            .catch(function(error){
              if(type !== "favor"){
                alert("读取笔记id出错"); 
                console.log("get note id error " + error.response.data.message);
              }
              else{
                self.favornotesById[id] = {};
                self.favornotesById[id].id = id;
                self.favornotesById[id].title = "找不到笔记";
                self.favornotesById[id].invalid = true;
                self.favornotesById[id].content = "";
                self.favornotesById[id].view_num = "";
                self.favornotesById[id].name = "";
                self.favornotesById[id].modify_time  = "";
              }
            });
          }
        }
        //console.log(response.data);
        console.log(self.notesById);
      })
      .catch(function(error){
        alert("读取笔记出错 " + error);
        //console.log(error);
      })
      ;
    },
    handleScroll() {
    const documentHeight = document.documentElement.scrollHeight;
    const currentScroll = window.scrollY || window.pageYOffset;
    const windowHeight = window.innerHeight;
    

    if (documentHeight - currentScroll <= windowHeight + 1) {
      var type = "";
        if(this.explore) type = "explore";
        else if(this.home) type = "home";
        else if(this.favor) type = "favor";
        else if(this.searchnote) type = "search";
        else return;
        this.loadMoreNotes(type);
    }
      // if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
        
      // }
    },
    addImageBox() {
      if(this.images.length < 3){
        this.waitingForUpload = true;
        // 自动触发文件上传输入框
        this.$refs.fileInput.click();
      }
      else {
        alert("图片数量达到上限");
      }
    },
    UploadArticle(){
      var self = this;
      if(this.postTitle == ""){
        alert("请输入标题！");
        return;
      }
      else if(this.imagefiles.length == 0){
        alert("请上传至少一张图片！");
        return;
      }
      const form = new FormData();
      form.append('article_title', this.postTitle);
      form.append('article_content', this.postText);
      form.append('image_num', this.images.length);
      //console.log("formdata: " + this.imagefiles);
      for(let i = 0; i < this.imagefiles.length; i++)
        form.append('image_list', this.imagefiles[i]);
      axios.post("http://resautu.cn:7879/upload_article", 
      form, {
        headers :{
          Authorization: self.token,
          'Content-Type': 'multipart/form-data'
        }
      }
    )
      .then(function(response){
        alert(response.data.message);
      })
      .catch(function(error){
        if(error.response !== undefined)
          alert("上传笔记失败！" + error.response.data.message);
        console.log("上传笔记错误 " + error);
      })
    },
    handleFileUpload(event) {
      const file = event.target.files[0];
      if (file) {
        this.images.push(null);
        this.imageBoxHovered.push(false);
        const reader = new FileReader();
        reader.onload = (e) => {
          // 将图片URL设置到当前空白框中
          this.images[this.images.length - 1] = e.target.result;
          this.imagefiles.push(file);
        };
        reader.readAsDataURL(file);
      }
    },
    removeImage(index) {
      // 从images数组中移除对应的图片
      this.images.splice(index, 1);
      // 从imageBoxHovered数组中移除对应的状态
      this.imageBoxHovered.splice(index, 1);
      this.imagefiles.splice(index, 1);
    },
    Init:function(){
      this.login = false;
      this.register = false;
      this.name = "";
      this.password = "";
    },
    BackToLogIn:function(){  
      this.login = false;
      this.register = false;
      this.name = "";
      this.password = "";
    },

    LogIn:function(){
      var self = this;
      axios.post("http://resautu.cn:7879/login", 
      //{name:this.name, password: this.password}
      querystring.stringify({
        
        name:this.name,
        password:this.password
      })
    )
      .then(function(response){
        alert(response.data.message);
        self.login = true;
        self.explore = true;
        var id = response.data.uname;
        axios.get(`http://resautu.cn:7879/user/uname/${id}`,{headers :{
          Authorization: self.token,
        }})
        .then(function(response){
          self.uname = response.data.uname;
        })
        .catch(function(error){
          if(error.response !== undefined){
            alert(error.response.data.message);
          }
        });
        self.token = response.data.token;
        self.loadMoreNotes("explore");
        self.loadMoreNotes("home");
        self.loadMoreNotes("favor");
      })
      .catch(function(error){
        if(error.response !== undefined)
          alert(error.response.data.message);
        console.log("LogIn error " + error);
      })
    },
    LogOut(){
      this.login = false; this.explore = false; 
      this.home = false; this.post = false; this.favor = false;
      this.display = false; this.searchnote = false;
      this.notesById = {} ; this.mynotesById = {}; this.favornotesById = {}; this.searchnotesById = {};
      this.displaynote = {};
      this.name = ""; this.uname = "";this.token = ""; this.password = ""; this.acquire = "";
      this.postText = ""; this.postTitle = "";
      alert("退出登陆成功！");
    },
    Regis:function(){
      this.register = true;
      this.name = "";
      this.password = "";
      this.acquire = "";
    },
    Register:function(){
      var self = this;
      if(this.password != this.acquire){
        alert("两次密码不一致！");
      }
      axios.post("http://resautu.cn:7879/register", 
      //{name:this.name, password: this.password}
      querystring.stringify({
        name:this.name,
        password:this.password,
      })
    )
      .then(function(response){
        self.BackToLogIn();
        alert(response.data.message);
      })
       .catch(function(error){
        alert(error.response.data.message);
       })
    },
    ToExplore:function(){
      this.explore = true;
      this.home = false;
      this.post = false;
      this.favor = false;
      this.display = false;
      this.searchnote = false;
      this.loadMoreNotes("explore");
      this.loadMoreNotes("home");
      this.loadMoreNotes("favor");
      this.searchnotesById = {};
    },
    ToPost:function(){
      this.explore = false;
      this.home = false;
      this.post = true;
      this.favor = false;
      this.display = false;
      this.searchnote = false;
    },
    ToHome:function(){
      this.explore = false;
      this.home = true;
      this.post = false;
      this.favor = false;
      this.display = false;
      this.searchnote = false;
      this.loadMoreNotes("home");
      this.loadMoreNotes("explore");
      this.loadMoreNotes("favor");
    },
    ToFavor:function(){
      this.explore = false;
      this.post = false;
      this.home = false;
      this.favor = true;
      this.searchnote = false;
      this.display = false;
      this.loadMoreNotes("favor");
      this.loadMoreNotes("home");
      this.loadMoreNotes("explore");
    }
  },
  beforeUnmount() {
    // 在组件销毁前移除监听
    window.removeEventListener('scroll', this.handleScroll);
  },
  props: {
    //msg: String
  }
}
</script>