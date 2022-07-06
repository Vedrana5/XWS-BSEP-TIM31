<template>
  <div>
<button @click="GoBack()">Go back</button>
      <input
        class="form-control mr-sm-2"
        type="text"
        placeholder="Username"
        v-model="username"
      />
      <button class="btn btn-success" type="submit" @click="Search(username)">
        Search
      </button>

        <div class="containerInfo">
    <div class="tab-pane container active">
      <div class="row-boats" v-for="(user, index) in users" :key="index">
        <div class="col-info">
          <h4 style="width: 600px" class="text">
            Username: {{user.Username}}
          </h4>
          <h4 style="width: 600px" class="text">Firstname: {{user.FirstName}}</h4>
          <h4 style="width: 600px" class="text">
            Lastname: {{user.LastName}}
          </h4>
          <img src=""/>
          <button  class="btn btn-success" @click="FindPosts(user.Username)">See all posts</button>
        </div>
      </div>
    </div>
</div>
        <div class="containerInfo">
    <div class="tab-pane container active">
      <div v-if="this.counter==1" >
                            <div v-for="(Post, index) in posts" :key="index">
                              <h4 style="width: 600px" class="text">
                                Id: {{Post.Id}}
                              </h4>
                              <h4 style="width: 600px" class="text">
                                Date: {{Post.DatePosted}}
                              </h4>
                            <h4 style="width: 600px" class="text">
                                Text: {{Post.PostText}}
                              </h4>
                              <h4 style="width: 600px" class="text">
                                Links:
                              </h4>
                              <div v-for="(link, index) in Post.Links " :key="index">
                              <a :href="link">{{ link}}</a>
                              </div>
                              <h4 style="width: 600px" class="text">Photo:</h4>
                              <img v-bind:src="'data:image/jpeg;base64,'+Post.ImagePaths"/>
                              <button @click="LikePost(Post.Id)">Like({{Post.LikesNumber}})</button>
                              <button @click="DislikePost(Post.Id)">DisLike({{Post.DislikesNumber}})</button>
                              <button @click="OpenComments(Post.Id)">Comment({{Post.CommentsNumber}})</button>
                            </div>

      </div>
      <div v-if="this.commentsCounter==1" >
                          <div v-for="(comment, index) in comments" :key="index">
                              <h4 style="width: 600px" class="text">
                                {{comment.Username}} : {{comment.CommentText}}
                              </h4>
                          </div>
                              <input class="input-field" placeholder="Comment" type="text" v-model="Addedcomment"/>
                              <button @click="AddComment(Addedcomment)">Add comment</button>
     </div>
     </div> 
</div>





  </div>
</template>

<script>
import axios from 'axios';
export default {
  name: "FindPublicUserByLogUser",
  data() {
    return {
      isLiked:true,
      isDisliked:true,
      Addedcomment:"",
      username:"",
      commentsCounter:0,
      comments:"",
      users:"",
      posts:"",
      link:"",
       Post:{Id:"",PostText:"", ImagePaths:null,Links:[]},
       counter:0,
      user: {
        username:"",
       firstName:"",
        lastName:"",
      },
      id:"",
      name:""
    };
  },


  methods: {
    async GoBack() {
      this.$router.push({ name: "StartPageUser" });
    },
    async Search(username) {
      this.username= username;

     await this.FindUser(username);
    },
  async FindUser(username) {
    this.username= username;
          axios.get("http://localhost:9090/publicUser/"+this.username)
      .then (response => { 
          this.users = response.data.users;
      })  
  },
  async LikePost(id) {
    console.log("hejjj");
    this.isLiked=false;
    this.id=id;
    this.name = localStorage.getItem("username");
    axios.post("http://localhost:9090/post/"+this.id+"/like",{
      PostId:this.id,
      Username: this.name,
    })
  },
    async DislikePost(id) {
    this.isDisliked=false;
    this.id=id;
    this.name = localStorage.getItem("username");
    axios.post("http://localhost:9090/post/"+this.id+"/dislike",{
      PostId:this.id,
      Username: this.name,
    })

  },
  async AddComment(comment) {
    this.name = localStorage.getItem("username");
    this.Addedcomment = comment;
    axios.post("http://localhost:9090/post/"+this.id+"/comment",{
      PostId:this.id,
      Username: this.name,
      CommentText: this.Addedcomment
    })
    this.$router.go(0);
  },
  async OpenComments(Id) {
    this.id=Id;
    this.commentsCounter=1;
    await this.FindComments(this.id);
  console.log("caooo!!!")
  },
  async FindComments(id) {
    this.id=id;
axios.get("http://localhost:9090/post/"+this.id+"/comments")
      .then (response => { 
          this.comments = response.data.Comments;
      })
  },
    async FindPosts(username) {
      this.counter=1;
        this.username= username;
        console.log("username je"+this.username);
          axios.get("http://localhost:9090/post/user/"+this.username)
      .then (response => { 
          this.posts = response.data.Posts;
          console.log("Odgovor je"+response.data.Posts)
      })
  },
  async findOneUser(username){
        this.username= username;
          axios.get("http://localhost:9090/publicUser/"+this.username)
      .then (response => { 
          this.user = response.data.users[0];
      })  
      await this.FindPosts(username);
       console.log("Broj postova je"+this.posts.length)
  }
  }
};
</script>


<style scoped>
.row-boats {
  display: flex;
}
.text {
  text-align: left;
  font-size: 15px;
  font-weight: 600;
  margin-top: 1%;
  margin-bottom: 1%;
}
.col-info {
  margin-left: 30px;
  margin-top: 10px;
}
.row {
  width: 660%;
  padding-left: 60px;
  height: 10%;
  margin-top: 100px;
  margin-left: 5000px;
  margin-right: 500px;
  opacity: 1.2;
  border-radius: 10px;
  align-content: center;
  background-color: rgba(255, 255, 255, 0.288);
  box-shadow: 0 5px 4px rgb(0 0 0 / 30%), 0 1px 1px rgb(0 0 0 / 22%);
}
/* kartica u okviru stavke menija  */
.containerInfo {
  width: 97%;
  margin-top: 20px;
  margin-left: 20px;
  margin-right: 10px;
  opacity: 1;
  border-radius: 10px;
  align-content: center;
  position: relative;
  background-color: fff;
  box-shadow: 0 19px 40px rgb(0 0 0 / 30%), 0 15px 12px rgb(0 0 0 / 22%);
}
.col-with-picture {
  margin-top: 1%;
  margin-bottom: 1%;
}
.col-with-picture {
  margin-top: 1%;
  margin-bottom: 1%;
}
</style>