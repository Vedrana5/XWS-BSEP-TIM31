<template>
  <div>
<button @click="GoBack()">Go back</button>
      <input
        class="form-control mr-sm-2"
        type="text"
        placeholder="Username"
        v-model="username"
      />
      <button class="btn btn-success" type="submit" @click="SearchUser(username)">
        Search
      </button>

        <div class="containerInfo">
            {{message}}
    <div class="tab-pane container active">
      <div class="row-boats" v-if="this.showUser==true">
        <div class="col-info">
          <button class="btn btn-success" v-if="this.showButton==true && user.TypeOfProfile=='PUBLIC'" @click="Follow(user.Username)">Follow</button>
          <button class="btn btn-success" v-if="this.showButton==true && user.TypeOfProfile=='PRIVATE'" @click="FollowRequest(user.Username)">Follow request</button>
          <button class="btn btn-success" @click="BlockUser(user.Username)">Block</button>
          <h4 style="width: 600px" class="text">
            Username: {{user.Username}}
          </h4>
          <h4 style="width: 600px" class="text">Firstname: {{user.FirstName}}</h4>
          <h4 style="width: 600px" class="text">
            Lastname: {{user.LastName}}
          </h4>
          <h4 style="width: 600px" class="text">
            Type of profile: {{user.TypeOfProfile}}
          </h4>
          <img src=""/>
          <button  class="btn btn-success" @click="FindPosts(user.Username)">See posts</button>
        </div>
      </div>
    </div>
</div>
<div class="containerInfo">
    <div class="tab-pane container active">
      <div v-if="(this.user.TypeOfProfile=='PUBLIC' && this.counter==1) || (this.user.TypeOfProfile=='PRIVATE' && this.show==true)" >
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
      <div v-if="(this.user.TypeOfProfile=='PUBLIC' && this.counter==1 && this.counter2==1) || (this.user.TypeOfProfile=='PRIVATE' && this.show==true && this.counter2==1)">
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
</div>-->





  </div>
</template>

<script>
import axios from 'axios';
import Swal from 'sweetalert2';
export default {
  name: "FindPublicUserByLogUser",
  data() {
    return {
      user: {
        Username:"",
       FirstName:"",
        LastName:"",
        TypeOfProfile:""
      },
      username:"",
      firstname:"",
      showUser:false,
      message:"",
      showButton:false,
      connection:{FirstUsername:"",SecondUsername:"", IsConfirmed:false, IsDeleted:false},
      counter:0,
      usernameForPosts:"",
      first:"",
      second:"",
      show:false,
      id:0,
      comments:"",
      counter2:0,
      posts:"",
      Post:{Id:"",PostText:"", ImagePaths:null,Links:[]},
      userWhoBlock:"",
      userwhoIsBlocked:"",
      block:{
        Id:0,
        FirstUsername: "",
        SecondUsername: ""
      },
    };
  },


  methods: {
    async BlockUser(username){
      this.userWhoBlock= localStorage.getItem("username");
      this.userwhoIsBlocked=username;
      axios.post("http://localhost:9090/blockedUser",{ 
        FirstUsername:this.userWhoBlock,
        SecondUsername:this.userwhoIsBlocked,
       })
      .then (response => { 
        console.log("RESPONSE JE"+response.data)
                      new Swal({
             title:"Uspesno",
             type: "warning",
             text:'Uspesno ste blokirali korisnika!',
           });     
            this.$router.push({ name: "StartPageUser" });
      })      

    },
    async GoBack(){
         this.$router.push({ name: "StartPageUser" });
    },
    async SearchUser(username) {
        this.username=username;
        this.firstname=localStorage.getItem("username")
        this.FindBlock(this.firstname,this.username);
        console.log("qqqq je:"+this.block.FirstUsername);
        console.log("wwww je:"+this.block.SecondUsername);
       
    },
    async FindBlock(firstname,username){
        this.username=username;
        this.firstname=firstname;
    axios.get("http://localhost:9090/blocked/"+this.firstname+"/"+this.username)
      .then (response => { 
        console.log("FIrstame je:"+response.data.block.FirstUsername)
        console.log("SecondUsername je:"+response.data.block.SecondUsername)
          this.block.FirstUsername = response.data.block.FirstUsername;
          this.block.SecondUsername = response.data.block.SecondUsername;
          if(this.block.FirstUsername=="" && this.block.SecondUsername=="") {
              axios.get("http://localhost:9090/blocked/"+this.username+"/"+this.firstname)
                    .then (response => { 
                                this.block.FirstUsername = response.data.block.FirstUsername;
                                this.block.SecondUsername = response.data.block.SecondUsername;
                                          if(this.block.FirstUsername=="" && this.block.SecondUsername=="") {
                                            this.FindThisUser(username);
                                            this.FindConnection(this.firstname,this.username);
                                            }
                    })
          } 
      })    

    },
  async OpenComments(Id) {
    this.id=Id;
    await this.FindComments(this.id);
  },
  async FindComments(id) {
    this.id=id;
axios.get("http://localhost:9090/post/"+this.id+"/comments")
      .then (response => { 
          this.comments = response.data.Comments;
          if (this.comments.length>0) {
            this.counter2==1;
          }
      })
  },
    async Follow(username){
      this.first=localStorage.getItem("username")
      this.second=username;
      axios.post("http://localhost:9090/connection",{ 
        FirstUsername:this.first,
        SecondUsername:this.second,
        IsConfirmed:true
       })
      .then (response => { 
        console.log("RESPONSE JE"+response.data)
                      new Swal({
             title:"Uspesno",
             type: "warning",
             text:'Uspesno ste zapratili korisnika!',
           });     
            this.$router.push({ name: "StartPageUser" });
      })
      this.$router.go(0)
    },
    async FollowRequest(username){
      this.first=localStorage.getItem("username")
      this.second=username;
      axios.post("http://localhost:9090/connection",{ 
        FirstUsername:this.first,
        SecondUsername:this.second,
        IsConfirmed:false
       })
      .then (response => { 
        console.log("RESPONSE JE"+response.data)
                      new Swal({
             title:"Uspesno",
             type: "warning",
             text:'Uspesno ste poslali zahtev za pracenje korisnika!',
           });     
            this.$router.push({ name: "StartPageUser" });
      })
this.$router.go(0)
    },    
    async FindConnection(firstname,username) {
        axios.get("http://localhost:9090/connect/"+firstname+"/"+username)
        .then (response => { 
          
          if(response.data.connection.FirstUsername=="") {
            this.showButton = true;
          }else {
            if (response.data.connection.IsDeleted==true) {
              this.showButton = true;
            }else {
            this.connection = response.data.connection;
            }
            if(response.data.connection.IsConfirmed==true) {
              this.show=true;
              console.log("Aj da vidim jel usao ovde!")
            }
          }
        }) 
    },
    async FindPosts(username) {
     
        this.username= username;
        console.log("username je"+this.username);
          axios.get("http://localhost:9090/post/user/"+this.username)
      .then (response => { 
        
          this.posts = response.data.Posts;
          console.log("duzina postova je"+this.posts.length)
          if(this.posts.length >0) {
            this.counter=1; 
          }
      })

    },
    async FindThisUser(username) {
        this.showUser=true;
        this.username=username; 
        axios.get("http://localhost:9090/user/"+this.username)
        .then (response => { 
          
          if(response.data.user.Username=="") {
            console.log("usla sam u if")
            this.showUser = false;
            this.message="User does not exists!"
          }else {
            console.log("usla sam u else")
            this.user = response.data.user;
          }
        })  
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